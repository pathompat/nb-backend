// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://tickbook.net/",
        "contact": {
            "name": "API Support",
            "url": "http://tickbook.net/support",
            "email": "support@tickbook.net"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/login": {
            "post": {
                "security": [
                    {
                        "JWTToken": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "login"
                ],
                "operationId": "Login",
                "parameters": [
                    {
                        "description": "Login request",
                        "name": "loginRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.Login"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.ResponseWithToken"
                        }
                    },
                    "400": {
                        "description": "Bad request"
                    },
                    "401": {
                        "description": "Unauthorized"
                    }
                }
            }
        },
        "/quotation": {
            "get": {
                "security": [
                    {
                        "JWTToken": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "quotations"
                ],
                "operationId": "GetAllQuotation",
                "parameters": [
                    {
                        "type": "boolean",
                        "name": "includeProduction",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/helper.ApiSuccessResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "array",
                                            "items": {
                                                "$ref": "#/definitions/dto.QuotationResponse"
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad request"
                    },
                    "401": {
                        "description": "Unauthorized"
                    }
                }
            }
        },
        "/school": {
            "get": {
                "security": [
                    {
                        "JWTToken": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "schools"
                ],
                "operationId": "GetSchoolByUserId",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User's UUID",
                        "name": "userId",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/helper.ApiSuccessResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "array",
                                            "items": {
                                                "$ref": "#/definitions/dto.SchoolResponse"
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad request"
                    },
                    "401": {
                        "description": "Unauthorized"
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "JWTToken": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "schools"
                ],
                "operationId": "CreateSchool",
                "parameters": [
                    {
                        "description": "Create school request",
                        "name": "createSchoolDTO",
                        "in": "body",
                        "schema": {
                            "$ref": "#/definitions/dto.CreateSchool"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/helper.ApiSuccessResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/dto.SchoolResponse"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad request"
                    },
                    "401": {
                        "description": "Unauthorized"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/user": {
            "get": {
                "security": [
                    {
                        "JWTToken": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "operationId": "GetAllUsers",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/helper.ApiSuccessResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "array",
                                            "items": {
                                                "$ref": "#/definitions/dto.UserResponse"
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad request"
                    },
                    "401": {
                        "description": "Unauthorized"
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "JWTToken": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "operationId": "CreateUser",
                "parameters": [
                    {
                        "description": "Create user request",
                        "name": "createUserDTO",
                        "in": "body",
                        "schema": {
                            "$ref": "#/definitions/dto.CreateUser"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/helper.ApiSuccessResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/dto.UserResponse"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad request"
                    },
                    "401": {
                        "description": "Unauthorized"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/user/info": {
            "get": {
                "security": [
                    {
                        "JWTToken": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "operationId": "GetInfoUser",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/helper.ApiSuccessResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/dto.UserResponse"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad request"
                    },
                    "401": {
                        "description": "Unauthorized"
                    }
                }
            }
        },
        "/user/{userId}": {
            "put": {
                "security": [
                    {
                        "JWTToken": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "operationId": "UpdateUser",
                "parameters": [
                    {
                        "description": "Update user request",
                        "name": "updateUserDTO",
                        "in": "body",
                        "schema": {
                            "$ref": "#/definitions/dto.UpdateUser"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/helper.ApiSuccessResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/dto.UserResponse"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad request"
                    },
                    "401": {
                        "description": "Unauthorized"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "JWTToken": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "operationId": "DeleteUser",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User's UUID",
                        "name": "userId",
                        "in": "path"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Bad request"
                    },
                    "401": {
                        "description": "Unauthorized"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        }
    },
    "definitions": {
        "dto.CreatePriceRef": {
            "type": "object",
            "properties": {
                "color": {
                    "description": "Color",
                    "type": "string",
                    "example": "1"
                },
                "gram": {
                    "description": "Gram",
                    "type": "integer",
                    "example": 12
                },
                "hasReference": {
                    "description": "HasReference",
                    "type": "boolean",
                    "example": false
                },
                "page": {
                    "description": "Page",
                    "type": "integer",
                    "example": 30
                },
                "pattern": {
                    "description": "Pattern",
                    "type": "string",
                    "example": "TABLE"
                },
                "plate": {
                    "description": "Plate",
                    "type": "string",
                    "example": "LARGE"
                },
                "priceRef": {
                    "description": "Price",
                    "type": "number",
                    "example": 5.5
                },
                "productTitle": {
                    "description": "ProductTitle",
                    "type": "string",
                    "example": "cut8"
                },
                "tierId": {
                    "description": "TierID",
                    "type": "integer",
                    "example": 1
                }
            }
        },
        "dto.CreateSchool": {
            "type": "object",
            "required": [
                "address",
                "name",
                "telephone",
                "userId"
            ],
            "properties": {
                "address": {
                    "description": "Address",
                    "type": "string",
                    "example": "81 test address"
                },
                "name": {
                    "description": "Name",
                    "type": "string",
                    "example": "school 2"
                },
                "telephone": {
                    "description": "Telephone",
                    "type": "string",
                    "example": "0815231112"
                },
                "userId": {
                    "description": "UserID",
                    "type": "string",
                    "example": "ebf889fd-4f3c-4c15-b44b-1d37cd2ee5e4"
                }
            }
        },
        "dto.CreateUser": {
            "type": "object",
            "required": [
                "password",
                "storeName",
                "tierId",
                "username"
            ],
            "properties": {
                "password": {
                    "description": "Secure password",
                    "type": "string",
                    "example": "Password@1234"
                },
                "storeName": {
                    "description": "User's shop name",
                    "type": "string",
                    "example": "Example Shop"
                },
                "tierId": {
                    "description": "User tier (1,2,3)",
                    "type": "integer",
                    "example": 1
                },
                "username": {
                    "description": "Username",
                    "type": "string",
                    "example": "testuser123"
                }
            }
        },
        "dto.Login": {
            "type": "object",
            "required": [
                "password",
                "username"
            ],
            "properties": {
                "password": {
                    "description": "Password",
                    "type": "string",
                    "minLength": 8,
                    "example": "Password@123"
                },
                "username": {
                    "description": "Username",
                    "type": "string",
                    "example": "testuser123"
                }
            }
        },
        "dto.PriceRefResponse": {
            "type": "object",
            "properties": {
                "color": {
                    "description": "Color",
                    "type": "string",
                    "example": "1"
                },
                "gram": {
                    "description": "Gram",
                    "type": "integer",
                    "example": 12
                },
                "hasReference": {
                    "description": "HasReference",
                    "type": "boolean",
                    "example": false
                },
                "page": {
                    "description": "Page",
                    "type": "integer",
                    "example": 30
                },
                "pattern": {
                    "description": "Pattern",
                    "type": "string",
                    "example": "TABLE"
                },
                "plate": {
                    "description": "Plate",
                    "type": "string",
                    "example": "LARGE"
                },
                "priceRef": {
                    "description": "Price",
                    "type": "number",
                    "example": 5.5
                },
                "productTitle": {
                    "description": "ProductTitle",
                    "type": "string",
                    "example": "cut8"
                }
            }
        },
        "dto.ResponseWithToken": {
            "type": "object",
            "properties": {
                "expiredIn": {
                    "description": "Token expired in (second)",
                    "type": "integer",
                    "example": 3600
                },
                "token": {
                    "description": "JWT Token",
                    "type": "string",
                    "example": "token123"
                }
            }
        },
        "dto.SchoolResponse": {
            "type": "object",
            "properties": {
                "address": {
                    "description": "address",
                    "type": "string",
                    "example": "22/11 test address"
                },
                "createdAt": {
                    "description": "Created user date",
                    "type": "string",
                    "example": "2024-12-02T00:26:21.087061Z"
                },
                "id": {
                    "description": "id",
                    "type": "integer",
                    "example": 1
                },
                "name": {
                    "description": "name",
                    "type": "string",
                    "example": "school 1"
                },
                "telephone": {
                    "description": "User tier (1,2,3)",
                    "type": "string",
                    "example": "0815231112"
                },
                "updatedAt": {
                    "description": "Latest update user date",
                    "type": "string",
                    "example": "2024-12-02T00:26:21.087061Z"
                }
            }
        },
        "dto.UpdateUser": {
            "type": "object",
            "required": [
                "password",
                "storeName",
                "tierId",
                "username"
            ],
            "properties": {
                "password": {
                    "description": "Secure password",
                    "type": "string",
                    "example": "Password@1234"
                },
                "storeName": {
                    "description": "User's shop name",
                    "type": "string",
                    "example": "Example Shop"
                },
                "tierId": {
                    "description": "User tier (1,2,3)",
                    "type": "integer",
                    "example": 1
                },
                "username": {
                    "description": "Username",
                    "type": "string",
                    "example": "testuser123"
                }
            }
        },
        "dto.UserResponse": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "description": "Created user date",
                    "type": "string",
                    "example": "2024-12-02T00:26:21.087061Z"
                },
                "id": {
                    "description": "UUID generate from database",
                    "type": "string",
                    "example": "be40de0f-ba3d-44d8-9c80-023ac23e0b9a"
                },
                "role": {
                    "description": "User role (ADMIN, CUSTOMER)",
                    "type": "string",
                    "example": "CUSTOMER"
                },
                "storeName": {
                    "description": "User's shop name",
                    "type": "string",
                    "example": "Test Store"
                },
                "tierId": {
                    "description": "User tier (1,2,3)",
                    "type": "integer",
                    "example": 1
                },
                "updatedAt": {
                    "description": "Latest update user date",
                    "type": "string",
                    "example": "2024-12-02T00:26:21.087061Z"
                },
                "username": {
                    "description": "Username",
                    "type": "string",
                    "example": "testuser1"
                }
            }
        },
        "helper.ApiSuccessResponse": {
            "type": "object",
            "properties": {
                "code": {
                    "description": "HTTP status",
                    "type": "integer",
                    "example": 200
                },
                "data": {
                    "description": "Returning data"
                },
                "message": {
                    "description": "Return message",
                    "type": "string",
                    "example": "Success"
                }
            }
        }
    },
    "securityDefinitions": {
        "JwtToken": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    },
    "externalDocs": {
        "description": "OpenAPI",
        "url": "https://swagger.io/resources/open-api/"
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{"https", "http"},
	Title:            "Tickbook API",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
