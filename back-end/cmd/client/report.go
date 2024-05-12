package main

import (
	"context"
	"log"

	pb "github.com/mxngocqb/VCS-SERVER/back-end/pkg/service/report/proto"
)

func doSendReport(c pb.ReportServiceClient, email []string, start, end string){
	log.Printf("Sending report to %v from %v to %v", email, start, end)


	res, err := c.Report(context.Background(), &pb.SendReportRequest{
		Mail: email,
		Start: start,
		End: end,
	})

	if err != nil {
		log.Printf("Error sending report: %v", err)
	} else {
		log.Printf("Report sent successfully: %v", res)
	}

}