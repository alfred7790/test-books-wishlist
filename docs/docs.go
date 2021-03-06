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
        "contact": {
            "name": "API Support",
            "email": "alfred.7790@gmail.com"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/v1/books/search": {
            "get": {
                "security": [
                    {
                        "APIToken": []
                    }
                ],
                "description": "List of books from google",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Google Books"
                ],
                "summary": "returns a list of books from google books",
                "parameters": [
                    {
                        "type": "string",
                        "description": "terms (Author, Title, Publisher)",
                        "name": "terms",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Google Books API key",
                        "name": "key",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/entity.BooksDTO"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/entity.FailureResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/entity.FailureResponse"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/entity.FailureResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/entity.FailureResponse"
                        }
                    }
                }
            }
        },
        "/v1/login": {
            "post": {
                "description": "Used login and get a new token",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "returns a new token",
                "parameters": [
                    {
                        "description": "username and password",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/entity.User"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/entity.UserTokenDTO"
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
        },
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
        },
        "/v1/wishlists": {
            "get": {
                "security": [
                    {
                        "APIToken": []
                    }
                ],
                "description": "List of books from google saved in datastore",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Wish List"
                ],
                "summary": "returns wishlist of books",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/entity.WishListDTO"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/entity.FailureResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/entity.FailureResponse"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
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
            },
            "post": {
                "security": [
                    {
                        "APIToken": []
                    }
                ],
                "description": "Used to create a wish list",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Wish List"
                ],
                "summary": "create or update a wish list for books",
                "parameters": [
                    {
                        "description": "struct to create a new wishlist",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/entity.WishList"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/entity.SuccessResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/entity.FailureResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/entity.FailureResponse"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
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
        },
        "/v1/wishlists/{id}": {
            "get": {
                "security": [
                    {
                        "APIToken": []
                    }
                ],
                "description": "List of books from google saved in datastore",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Wish List"
                ],
                "summary": "returns wishlist of books",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "WishListID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/entity.WishListDTO"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/entity.FailureResponse"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
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
            },
            "delete": {
                "security": [
                    {
                        "APIToken": []
                    }
                ],
                "description": "all books from a wishlist will be removed",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Wish List"
                ],
                "summary": "removing a complete wishlist",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "WishListID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/entity.SuccessResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/entity.FailureResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/entity.FailureResponse"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
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
        },
        "/v1/wishlists/{id}/books": {
            "post": {
                "security": [
                    {
                        "APIToken": []
                    }
                ],
                "description": "adding a book into a wishlist and creating a new book if it doesn't exist",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Books of Wish List"
                ],
                "summary": "adding a book into a wishlist",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "WishListID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "struct to add a book to wishlist",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/entity.Book"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/entity.SuccessResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/entity.FailureResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/entity.FailureResponse"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
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
        },
        "/v1/wishlists/{id}/books/{bookid}": {
            "delete": {
                "security": [
                    {
                        "APIToken": []
                    }
                ],
                "description": "removing a book from wishlist if it exists",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Books of Wish List"
                ],
                "summary": "removing a book from wishlist",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "WishListID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "BookID",
                        "name": "bookid",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/entity.SuccessResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/entity.FailureResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/entity.FailureResponse"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
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
        "entity.Book": {
            "type": "object",
            "properties": {
                "authors": {
                    "type": "string",
                    "example": "Miranda De Moura"
                },
                "id": {
                    "type": "string",
                    "example": "KNYxEAAAQBAJ"
                },
                "publisher": {
                    "type": "string",
                    "example": "Clube de Autores"
                },
                "title": {
                    "type": "string",
                    "example": "Viagens Na Madrugada"
                }
            }
        },
        "entity.BooksDTO": {
            "type": "object",
            "properties": {
                "books": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/entity.Book"
                    }
                },
                "items": {
                    "type": "integer",
                    "example": 20
                },
                "results": {
                    "type": "integer",
                    "example": 999
                }
            }
        },
        "entity.FailureResponse": {
            "type": "object",
            "properties": {
                "details": {
                    "type": "string",
                    "example": "Message error for developers"
                },
                "message": {
                    "type": "string",
                    "example": "Message error for users"
                },
                "status": {
                    "type": "string",
                    "example": "failure"
                }
            }
        },
        "entity.SuccessResponse": {
            "type": "object",
            "properties": {
                "details": {
                    "type": "string",
                    "example": "operation completed"
                },
                "message": {
                    "type": "string",
                    "example": "ok"
                },
                "success": {
                    "type": "boolean",
                    "example": true
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
        },
        "entity.UserTokenDTO": {
            "type": "object",
            "properties": {
                "accessBearerToken": {
                    "type": "string",
                    "example": "Bearer eyJhbGciOiJ..."
                },
                "accessToken": {
                    "type": "string",
                    "example": "eyJhbGciOiJ..."
                },
                "userID": {
                    "type": "integer",
                    "example": 1
                },
                "username": {
                    "type": "string",
                    "example": "Pantufla89"
                }
            }
        },
        "entity.WishList": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string",
                    "example": "MyFirstWishList"
                }
            }
        },
        "entity.WishListDTO": {
            "type": "object",
            "properties": {
                "books": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/entity.Book"
                    }
                },
                "createdAt": {
                    "type": "string",
                    "example": "2022-06-24T19:38:46.814728Z"
                },
                "id": {
                    "type": "integer",
                    "example": 1
                },
                "name": {
                    "type": "string",
                    "example": "MyFirstWishList"
                },
                "userId": {
                    "type": "integer",
                    "example": 1
                }
            }
        }
    },
    "securityDefinitions": {
        "APIToken": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
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
