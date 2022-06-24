// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
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
        "/v1/users": {
            "post": {
                "description": "Used to create a new user",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "returns details about a new user was created",
                "parameters": [
                    {
                        "description": "struct to create a new user",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/entity.User"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/entity.UserDTO"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/entity.FailureResponse"
                        }
                    },
                    "409": {
                        "description": "Conflict",
                        "schema": {
                            "$ref": "#/definitions/entity.FailureResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/entity.FailureResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "entity.FailureResponse": {
            "type": "object",
            "properties": {
                "details": {
                    "type": "string",
                    "example": "Error message for developers"
                },
                "message": {
                    "type": "string",
                    "example": "Error message for users"
                },
                "status": {
                    "type": "string",
                    "example": "failure"
                }
            }
        },
        "entity.User": {
            "type": "object",
            "properties": {
                "password": {
                    "type": "string",
                    "example": "Pantufla1234"
                },
                "username": {
                    "type": "string",
                    "example": "Pantufla89"
                }
            }
        },
        "entity.UserDTO": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer",
                    "example": 1
                },
                "username": {
                    "type": "string",
                    "example": "Pantufla89"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Google Books API - TESTCASE",
	Description:      "Google Books testcase API",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
