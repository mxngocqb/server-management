// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/login": {
            "post": {
                "description": "Authenticates user and returns a JWT token.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Authentication"
                ],
                "summary": "User login",
                "parameters": [
                    {
                        "description": "Login credentials",
                        "name": "LoginRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/internal_handler_auth_transport.LoginRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Login successful",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "400": {
                        "description": "Invalid input",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/servers": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Retrieves a list of servers based on the provided filters and pagination.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Server"
                ],
                "summary": "View servers",
                "parameters": [
                    {
                        "type": "integer",
                        "default": 50,
                        "description": "Number of servers returned",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "default": 0,
                        "description": "Offset in server list",
                        "name": "offset",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Filter by status",
                        "name": "status",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "The field to sort by",
                        "name": "field",
                        "in": "query"
                    },
                    {
                        "enum": [
                            "asc",
                            "desc"
                        ],
                        "type": "string",
                        "description": "Arrangement order",
                        "name": "order",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/github_com_mxngocqb_VCS-SERVER_back-end_internal_model.Server"
                            }
                        }
                    },
                    "400": {
                        "description": "Invalid parameters for limit or offset",
                        "schema": {
                            "$ref": "#/definitions/echo.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Failed to fetch servers due to server error",
                        "schema": {
                            "$ref": "#/definitions/echo.HTTPError"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Adds a new server to the database",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Server"
                ],
                "summary": "Create server",
                "parameters": [
                    {
                        "description": "Server data to create",
                        "name": "server",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/internal_handler_server_transport.CreateRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/github_com_mxngocqb_VCS-SERVER_back-end_internal_model.Server"
                        }
                    },
                    "400": {
                        "description": "Bad request - Invalid server data",
                        "schema": {
                            "$ref": "#/definitions/echo.HTTPError"
                        }
                    },
                    "403": {
                        "description": "Forbidden - User does not have permission to create server",
                        "schema": {
                            "$ref": "#/definitions/echo.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal server error - Failed to create serve",
                        "schema": {
                            "$ref": "#/definitions/echo.HTTPError"
                        }
                    }
                }
            }
        },
        "/servers/export": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Exports filtered server data to an Excel file.",
                "produces": [
                    "application/octet-stream"
                ],
                "tags": [
                    "Server"
                ],
                "summary": "Export servers",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Filter by creation date start",
                        "name": "startCreated",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Filter by creation date end",
                        "name": "endCreated",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Filter by update date start",
                        "name": "startUpdated",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Filter by update date end",
                        "name": "endUpdated",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Field to sort by",
                        "name": "field",
                        "in": "query"
                    },
                    {
                        "enum": [
                            "asc",
                            "desc"
                        ],
                        "type": "string",
                        "description": "Arrangement order",
                        "name": "order",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Excel file",
                        "schema": {
                            "type": "file"
                        }
                    },
                    "400": {
                        "description": "Bad request - Invalid filter parameters",
                        "schema": {
                            "$ref": "#/definitions/echo.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal server error - Failed to generate or send file",
                        "schema": {
                            "$ref": "#/definitions/echo.HTTPError"
                        }
                    }
                }
            }
        },
        "/servers/import": {
            "post": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Creates multiple servers from an uploaded Excel file.",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Server"
                ],
                "summary": "Bulk create servers",
                "parameters": [
                    {
                        "type": "file",
                        "description": "Excel file with list server data",
                        "name": "listserver",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "400": {
                        "description": "Bad request - Invalid or corrupt file",
                        "schema": {
                            "$ref": "#/definitions/echo.HTTPError"
                        }
                    },
                    "403": {
                        "description": "Forbidden - User does not have permission to delete server",
                        "schema": {
                            "$ref": "#/definitions/echo.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal server error - Failed to parse or save servers",
                        "schema": {
                            "$ref": "#/definitions/echo.HTTPError"
                        }
                    }
                }
            }
        },
        "/servers/report": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Retrieves a report of server statuses for a given date range and sends it to the specified email address.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Server"
                ],
                "summary": "Generate server status report",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Start Date",
                        "name": "start",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "End Date",
                        "name": "end",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "name": "mail",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Report sent successfully",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Invalid date format or email",
                        "schema": {
                            "$ref": "#/definitions/echo.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Error occurred while sending the report",
                        "schema": {
                            "$ref": "#/definitions/echo.HTTPError"
                        }
                    }
                }
            }
        },
        "/servers/{id}": {
            "put": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Updates server details.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Server"
                ],
                "summary": "Update a server",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Server ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Server update data",
                        "name": "server",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/internal_handler_server_transport.UpdateRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/github_com_mxngocqb_VCS-SERVER_back-end_internal_model.Server"
                        }
                    },
                    "400": {
                        "description": "Bad request - Invalid update data",
                        "schema": {
                            "$ref": "#/definitions/echo.HTTPError"
                        }
                    },
                    "403": {
                        "description": "Forbidden - User does not have permission to delete server",
                        "schema": {
                            "$ref": "#/definitions/echo.HTTPError"
                        }
                    },
                    "404": {
                        "description": "Not found - Server not found",
                        "schema": {
                            "$ref": "#/definitions/echo.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal server error - Failed to update server",
                        "schema": {
                            "$ref": "#/definitions/echo.HTTPError"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Removes a server based on ID.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Server"
                ],
                "summary": "Delete a server",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Server ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content"
                    },
                    "403": {
                        "description": "Forbidden - User does not have permission to delete server",
                        "schema": {
                            "$ref": "#/definitions/echo.HTTPError"
                        }
                    },
                    "404": {
                        "description": "Not found - Server not found",
                        "schema": {
                            "$ref": "#/definitions/echo.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal server error - Failed to delete server",
                        "schema": {
                            "$ref": "#/definitions/echo.HTTPError"
                        }
                    }
                }
            }
        },
        "/servers/{id}/uptime": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Returns the uptime of a server based on a specific date provided in the query.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Server"
                ],
                "summary": "Retrieve server uptime",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Server ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Date",
                        "name": "date",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Hours of uptime",
                        "schema": {
                            "type": "number"
                        }
                    },
                    "400": {
                        "description": "Invalid date format or server ID",
                        "schema": {
                            "$ref": "#/definitions/echo.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal server error occurred while retrieving uptime",
                        "schema": {
                            "$ref": "#/definitions/echo.HTTPError"
                        }
                    }
                }
            }
        },
        "/users": {
            "post": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Create a new user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Create user",
                "parameters": [
                    {
                        "description": "Create User",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/internal_handler_user_transport.CreateRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/github_com_mxngocqb_VCS-SERVER_back-end_internal_model.User"
                        }
                    },
                    "400": {
                        "description": "Bad Request - Invalid user data",
                        "schema": {
                            "$ref": "#/definitions/echo.HTTPError"
                        }
                    },
                    "403": {
                        "description": "Forbidden - Insufficient permissions",
                        "schema": {
                            "$ref": "#/definitions/echo.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error - Unable to create user",
                        "schema": {
                            "$ref": "#/definitions/echo.HTTPError"
                        }
                    }
                }
            }
        },
        "/users/{id}": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Get details of a user by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "View user",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "User ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/github_com_mxngocqb_VCS-SERVER_back-end_internal_model.User"
                        }
                    },
                    "400": {
                        "description": "Bad Request - Invalid user ID format",
                        "schema": {
                            "$ref": "#/definitions/echo.HTTPError"
                        }
                    },
                    "404": {
                        "description": "Not Found - User not found",
                        "schema": {
                            "$ref": "#/definitions/echo.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error - Unable to retrieve user",
                        "schema": {
                            "$ref": "#/definitions/echo.HTTPError"
                        }
                    }
                }
            },
            "put": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Update a user by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Update user",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "User ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Update User",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/internal_handler_user_transport.UpdateRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/github_com_mxngocqb_VCS-SERVER_back-end_internal_model.User"
                        }
                    },
                    "400": {
                        "description": "Bad Request - Invalid user data or ID",
                        "schema": {
                            "$ref": "#/definitions/echo.HTTPError"
                        }
                    },
                    "403": {
                        "description": "Forbidden - Insufficient permissions",
                        "schema": {
                            "$ref": "#/definitions/echo.HTTPError"
                        }
                    },
                    "404": {
                        "description": "Not Found - User not found",
                        "schema": {
                            "$ref": "#/definitions/echo.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error - Unable to update user",
                        "schema": {
                            "$ref": "#/definitions/echo.HTTPError"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Delete a user by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Delete user",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "User ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content"
                    },
                    "403": {
                        "description": "Forbidden - Insufficient permissions",
                        "schema": {
                            "$ref": "#/definitions/echo.HTTPError"
                        }
                    },
                    "404": {
                        "description": "Not Found - User not found",
                        "schema": {
                            "$ref": "#/definitions/echo.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error - Unable to delete user",
                        "schema": {
                            "$ref": "#/definitions/echo.HTTPError"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "echo.HTTPError": {
            "type": "object",
            "properties": {
                "message": {}
            }
        },
        "github_com_mxngocqb_VCS-SERVER_back-end_internal_model.Role": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "github_com_mxngocqb_VCS-SERVER_back-end_internal_model.Server": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "type": "string"
                },
                "deletedAt": {
                    "$ref": "#/definitions/gorm.DeletedAt"
                },
                "id": {
                    "type": "integer"
                },
                "ip": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "status": {
                    "type": "boolean"
                },
                "updatedAt": {
                    "type": "string"
                }
            }
        },
        "github_com_mxngocqb_VCS-SERVER_back-end_internal_model.User": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "type": "string"
                },
                "deletedAt": {
                    "$ref": "#/definitions/gorm.DeletedAt"
                },
                "id": {
                    "type": "integer"
                },
                "password": {
                    "description": "Password should be hashed and never returned in API calls",
                    "type": "string"
                },
                "role_ids": {
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                },
                "roles": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/github_com_mxngocqb_VCS-SERVER_back-end_internal_model.Role"
                    }
                },
                "updatedAt": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "gorm.DeletedAt": {
            "type": "object",
            "properties": {
                "time": {
                    "type": "string"
                },
                "valid": {
                    "description": "Valid is true if Time is not NULL",
                    "type": "boolean"
                }
            }
        },
        "internal_handler_auth_transport.LoginRequest": {
            "type": "object",
            "required": [
                "password",
                "username"
            ],
            "properties": {
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "internal_handler_server_transport.CreateRequest": {
            "type": "object",
            "required": [
                "ip",
                "name",
                "status"
            ],
            "properties": {
                "ip": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "status": {
                    "type": "boolean"
                }
            }
        },
        "internal_handler_server_transport.UpdateRequest": {
            "type": "object",
            "required": [
                "ip",
                "name"
            ],
            "properties": {
                "ip": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "status": {
                    "type": "boolean"
                }
            }
        },
        "internal_handler_user_transport.CreateRequest": {
            "type": "object",
            "required": [
                "password",
                "role_ids",
                "username"
            ],
            "properties": {
                "password": {
                    "type": "string"
                },
                "role_ids": {
                    "type": "array",
                    "minItems": 1,
                    "items": {
                        "type": "integer"
                    }
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "internal_handler_user_transport.UpdateRequest": {
            "type": "object",
            "required": [
                "password",
                "role_ids",
                "username"
            ],
            "properties": {
                "password": {
                    "type": "string"
                },
                "role_ids": {
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                },
                "username": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "Bearer": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8090",
	BasePath:         "/api",
	Schemes:          []string{"http", "https"},
	Title:            "Viettel Cyber Security - Server Management System",
	Description:      "This is the API documentation for the Viettel Cyber Security - Server Management System.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
