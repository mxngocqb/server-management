package server

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	// labstak/echo is a web framework for Go
	"github.com/labstack/echo/v4"
	"github.com/mxngocqb/VCS-SERVER/back-end/internal/handler"
	"github.com/mxngocqb/VCS-SERVER/back-end/internal/model"
	"github.com/mxngocqb/VCS-SERVER/back-end/internal/repository"
	"github.com/mxngocqb/VCS-SERVER/back-end/pkg/service/cache"
	"github.com/mxngocqb/VCS-SERVER/back-end/pkg/service/elastic"
	"github.com/mxngocqb/VCS-SERVER/back-end/pkg/service/kafka"
	pb "github.com/mxngocqb/VCS-SERVER/back-end/pkg/service/report/proto"
	util "github.com/mxngocqb/VCS-SERVER/back-end/pkg/util"
	"gorm.io/gorm"

	// gRPC framework for Go
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type IServerService interface {
	View(c echo.Context, perPage int, offset int, status, field, order string) ([]model.Server, int, error)
	Create(c echo.Context, server *model.Server) (*model.Server, error)
	CreateMany(c echo.Context, servers []model.Server) ([]model.Server, []string, []string, error)
	Update(c echo.Context, id string, server *model.Server) (*model.Server, error)
	Delete(c echo.Context, id string) error
	GetServersFiltered(c echo.Context, perPage int, offset int, status, field, order string) error
	GetServerUptime(c echo.Context, serverID string, date string) (time.Duration, error)
	GetServerReport(c echo.Context, mail, start, end string) error
	GetServerStatus(c echo.Context) (online int64, offline int64, err error)
	
}

type Service struct {
	repository repository.ServerRepository
	rbac       handler.RbacService
	elastic    elastic.ElasticService
	cache      cache.ServerCache
	producer   *kafka.ProducerService
}

func NewServerService(repository repository.ServerRepository, rbac handler.RbacService, elastic elastic.ElasticService, sc cache.ServerCache,
	producer *kafka.ProducerService) *Service {
	return &Service{
		repository: repository,
		rbac:       rbac,
		elastic:    elastic,
		cache:      sc,
		producer:   producer,
	}
}

// GetServerStatus retrieves the number of online and offline servers.
func (s *Service) GetServerStatus(c echo.Context) (int64, int64, error) {
	// Role ID for viewing server status is 1 (Admin)
	online, offline, err := s.repository.GetServerStatus()
	if err != nil {
		return 0, 0, echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Failed to retrieve server status: %v", err))
	}
	return online, offline, nil
}


// View retrieves servers from the database with optional pagination and status filtering.
func (s *Service) View(c echo.Context, perPage int, offset int, status, field, order string) ([]model.Server, int, error) {
	// ctx := c.Request().Context()
	key := s.cache.ConstructCacheKey(perPage, offset, status, field, order)
	key_total := key + "_total"
	// Try to get data from Redis first
	values := s.cache.GetMultiRequest(key)
	total := s.cache.GetTotalServer(key_total)

	if values != nil && total != -1 {
		// Data found in cache
		return values, total, nil
	}

	// Data not found in cache, fetch from database
	servers, numberOfServer, err := s.repository.GetServersFiltered(perPage, offset, status, field, order)
	if err != nil {
		return nil, -1, err
	}

	s.cache.SetMultiRequest(key, servers) // Adjust expiration as needed
	s.cache.SetTotalServer(key_total, numberOfServer)
	return servers, numberOfServer, nil
}

// GetServersFiltered retrieves servers with optional date range filtering.
func (s *Service) GetServersFiltered(c echo.Context, perPage int, offset int, status, field, order string) error {
	key := s.cache.ConstructCacheKey(perPage, offset, status, field, order)
	key_total := key + "_total"
	// Try to get data from Redis first
	values := s.cache.GetMultiRequest(key)
	total := s.cache.GetTotalServer(key_total)

	if values != nil && total != -1 {
		f, err2 := util.CreateExcelFile(values)
		if err2 != nil {
			return echo.NewHTTPError(http.StatusConflict, fmt.Sprintf("Error creating Excel file: %v", err2))
		}

		// Save Excel file to disk
		filePath := "export.xlsx"
		if err := f.SaveAs(filePath); err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "Failed to save Excel file")
		}

		// Serve the file
		return c.Attachment(filePath, "export.xlsx")
	}	
	
	// Data not found in cache, fetch from database
	servers, _, err := s.repository.GetServersFiltered(perPage, offset, status, field, order)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, fmt.Sprintf("Failed to retrieve servers: %v", err))
	}

	f, err2 := util.CreateExcelFile(servers)

	if err2 != nil {
		return echo.NewHTTPError(http.StatusConflict, fmt.Sprintf("Error creating Excel file: %v", err2))
	}

	// Save Excel file to disk
	filePath := "export.xlsx"
	if err := f.SaveAs(filePath); err != nil {
		return echo.NewHTTPError(http.StatusConflict, "Failed to save Excel file")
	}

	// Serve the file
	return c.Attachment(filePath, "export.xlsx")
}

// Create creates a new server.
func (s *Service) Create(c echo.Context, server *model.Server) (*model.Server, error) {
	s.cache.InvalidateCache()
	fmt.Println("Create call")
	// Role ID for creating a new server is 1 (Admin)
	requiredRoleID := uint(1)

	// Enforce role check
	if err := s.rbac.EnforceRole(c, requiredRoleID); err != nil {
		return &model.Server{}, err // This will handle forbidden access
	}

	// Create new server in the database
	err := s.repository.Create(server)
	if err != nil {
		return &model.Server{}, echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Failed to create server: %v", err))
	}

	fmt.Println("server", server)

	err = s.elastic.IndexServer(*server)
	if err != nil {
		return &model.Server{}, echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Failed to index server: %v", err))
	}

	// After successfully creating the server, log the status change
	err = s.elastic.LogStatusChange(*server, server.Status)
	if err != nil {
		// Handle logging error, you may choose to return an error or just log it
		return &model.Server{}, echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Error logging status change: %v", err))
	}

	s.producer.SendServer(server.ID, *server)

	// Cache the server
	s.cache.Set(strconv.Itoa(int(server.ID)), server)

	return server, nil
}

// CreateMany creates multiple servers and returns detailed results.
func (s *Service) CreateMany(c echo.Context, servers []model.Server) ([]model.Server, []string, []string, error) {
	s.cache.InvalidateCache()
	requiredRoleID := uint(1)
	if err := s.rbac.EnforceRole(c, requiredRoleID); err != nil {
		return nil, nil, nil, err
	}

	var createdServers []model.Server
	var successLines, failedLines []string

	for _, server := range servers {
		err := s.repository.Create(&server)
		if err != nil {
			failedLines = append(failedLines, server.Name) // +2 to account for zero index and header row
			continue
		}
		createdServers = append(createdServers, server)
		successLines = append(successLines, server.Name)

		err = s.elastic.IndexServer(server)
		if err != nil {
			log.Printf("Error indexing server %d: %v", server.ID, err)
		}
		// After successfully creating the server, log the status change
		err = s.elastic.LogStatusChange(server, server.Status)
		s.producer.SendServer(server.ID, server)
		if err != nil {
			// Handle logging error
			log.Printf("Error logging status change for server ID in Elasticsearch", server.ID)
		}
	}

	return createdServers, successLines, failedLines, nil
}

// Update updates a server.
func (s *Service) Update(c echo.Context, id string, server *model.Server) (*model.Server, error) {
	s.cache.InvalidateCache()
	// Role ID for updating a server is 1 (Admin)
	requiredRoleID := uint(1)

	// Enforce role check
	if err := s.rbac.EnforceRole(c, requiredRoleID); err != nil {
		return &model.Server{}, err // This will handle forbidden access
	}

	// Update server in Elasticsearch
	existingServer, err := s.repository.GetServerByID(id)
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Failed to retrieve server: %v", err))
	}

	existingServerStatus := existingServer.Status

	// Update server in the database
	err = s.repository.Update(id, server)
	if err != nil {
		return &model.Server{}, echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Failed to update server: %v", err))
	}
	// Retrieve updated server
	updatedServer, err := s.repository.GetServerByID(id)
	if err != nil {
		return &model.Server{}, echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Failed to retrieve updated server: %v", err))
	}

	if existingServerStatus != updatedServer.Status {
		err = s.elastic.LogStatusChange(*updatedServer, server.Status)
		if err != nil {
			return nil, echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Error logging status change: %v", err))
		}
	}

	// Cache the server
	s.producer.SendServer(server.ID, *server)
	s.cache.Set(strconv.Itoa(int(server.ID)), server)

	return updatedServer, nil
}

// Delete deletes a server.
func (s *Service) Delete(c echo.Context, id string) error {
	s.cache.InvalidateCache()
	// Role ID for deleting a server is 1 (Admin)
	requiredRoleID := uint(1)

	// Enforce role check
	if err := s.rbac.EnforceRole(c, requiredRoleID); err != nil {
		return err // This will handle forbidden access
	}

	server, err := s.repository.GetServerByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return echo.NewHTTPError(http.StatusNotFound, fmt.Sprintf("Server with ID %s not found", id))
		}
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Failed to retrieve server: %v", err))
	}

	// Log status change before deleting the server
	err = s.elastic.LogStatusChange(*server, false)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Error logging status change: %v", err))
	}

	// Delete server from the database
	err = s.repository.Delete(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Failed to delete server: %v", err))
	}

	// DELETE FROM ELASTICSEARCH MIGHT CAUSE ERROR IF THE SERVERS ARE NOT CREATED USING THE ENDPOINT (THEY ARE NOT CREATED IN ELASTICSEARCH IF USING SQL COMMAND ONLY)
	// Delete server from Elasticsearch
	err = s.elastic.DeleteServerFromIndex(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Failed to delete server from Elasticsearch: %v", err))
	}

	err = s.elastic.DeleteServerLogs(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Failed to delete server logs from Elasticsearch: %v", err))
	}
	// Cache the server
	s.cache.Delete(strconv.Itoa(int(server.ID)))
	s.producer.DropServer(server.ID)
	return nil
}

// GetServerUptime calculates the uptime for a server for the entire specified day.
func (s *Service) GetServerUptime(c echo.Context, serverID string, date string) (time.Duration, error) {
	layout := "2006-01-02"
	day, err := time.Parse(layout, date)
	if err != nil {
		return 0, echo.NewHTTPError(http.StatusBadRequest, "Invalid date format: "+err.Error())
	}

	uptime, err := s.elastic.CalculateServerUptime(serverID, day)
	if err != nil {
		return 0, echo.NewHTTPError(http.StatusBadRequest, "Error calculating uptime: "+err.Error())
	}

	return uptime, nil
}

// GetServerReport sends a report of server statuses within a specified date range to the client.
func (s *Service) GetServerReport(c echo.Context, mail, start, end string) error {
	layout := "2006-01-02"
	location, err := time.LoadLocation("Asia/Bangkok") // Load the GMT+7 timezone

	if _, err := time.ParseInLocation(layout, start, location); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid start date format")
	}

	if _, err := time.ParseInLocation(layout, end, location); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid end date format")
	}

	mailArr := []string{mail}

	// Create a gRPC client
	var addr string = "report:50052" // Address of the gRPC server
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Printf("Failed to dial server: %v", err)
	}

	defer conn.Close()

	client := pb.NewReportServiceClient(conn)

	err = doSendReport(client, mailArr, start, end)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Error sending report: "+err.Error())
	}

	return c.String(http.StatusOK, "Report sent successfully")
}
