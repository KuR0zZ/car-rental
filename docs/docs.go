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
        "/cars": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Retrieves a list of cars",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Cars"
                ],
                "summary": "Get all cars",
                "parameters": [
                    {
                        "type": "string",
                        "description": "With the bearer started",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Car"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/dtos.ErrorInternalServerError"
                        }
                    }
                }
            }
        },
        "/cars/:id": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Retrieves a car by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Cars"
                ],
                "summary": "Get car by id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "With the bearer started",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Car"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/dtos.ErrorNotFound"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/dtos.ErrorInternalServerError"
                        }
                    }
                }
            }
        },
        "/rentals/rent": {
            "post": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Allows a user to rent a car given a valid request.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Rentals"
                ],
                "summary": "Rent a car",
                "parameters": [
                    {
                        "type": "string",
                        "description": "With the bearer started",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "Rent Request",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dtos.RentRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/dtos.RentResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/dtos.ErrorBadRequest"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/dtos.ErrorNotFound"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/dtos.ErrorUnprocessableEntity"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/dtos.ErrorInternalServerError"
                        }
                    }
                }
            }
        },
        "/rentals/report": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Retrieves a list of rentals associated with the user.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Rentals"
                ],
                "summary": "Get rental report",
                "parameters": [
                    {
                        "type": "string",
                        "description": "With the bearer started",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/dtos.RentalReportResponse"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/dtos.ErrorInternalServerError"
                        }
                    }
                }
            }
        },
        "/users/login": {
            "post": {
                "description": "Authenticates a user and returns a JWT token.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "User login",
                "parameters": [
                    {
                        "description": "Login Request",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dtos.LoginRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dtos.LoginResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/dtos.ErrorBadRequest"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/dtos.ErrorUnauthorized"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/dtos.ErrorNotFound"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/dtos.ErrorInternalServerError"
                        }
                    }
                }
            }
        },
        "/users/register": {
            "post": {
                "description": "Creates a new user account with the provided details.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Register a new user",
                "parameters": [
                    {
                        "description": "Register Request",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dtos.RegisterRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/dtos.RegisterResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/models.User"
                                        },
                                        "message": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/dtos.ErrorBadRequest"
                        }
                    },
                    "409": {
                        "description": "Conflict",
                        "schema": {
                            "$ref": "#/definitions/dtos.ErrorConflict"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/dtos.ErrorInternalServerError"
                        }
                    }
                }
            }
        },
        "/users/topup": {
            "post": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Allows a user to add balance to their account.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Top up balance",
                "parameters": [
                    {
                        "type": "string",
                        "description": "With the bearer started",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "Top Up Request",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dtos.TopUpRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dtos.TopUpResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/dtos.ErrorBadRequest"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/dtos.ErrorInternalServerError"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "dtos.ErrorBadRequest": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string",
                    "example": "bad request data"
                }
            }
        },
        "dtos.ErrorConflict": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string",
                    "example": "email is already registered"
                }
            }
        },
        "dtos.ErrorInternalServerError": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string",
                    "example": "internal server error"
                }
            }
        },
        "dtos.ErrorNotFound": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string",
                    "example": "data not found"
                }
            }
        },
        "dtos.ErrorUnauthorized": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string",
                    "example": "invalid credential"
                }
            }
        },
        "dtos.ErrorUnprocessableEntity": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string",
                    "example": "insufficient balance"
                }
            }
        },
        "dtos.LoginRequest": {
            "type": "object",
            "required": [
                "email",
                "password"
            ],
            "properties": {
                "email": {
                    "type": "string",
                    "example": "John.Doe@example.com"
                },
                "password": {
                    "type": "string",
                    "example": "password123"
                }
            }
        },
        "dtos.LoginResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string",
                    "example": "Successfully logged in"
                },
                "token": {
                    "type": "string",
                    "example": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"
                }
            }
        },
        "dtos.RegisterRequest": {
            "type": "object",
            "required": [
                "email",
                "name",
                "password"
            ],
            "properties": {
                "email": {
                    "type": "string",
                    "example": "John.Doe@example.com"
                },
                "name": {
                    "type": "string",
                    "example": "John Doe"
                },
                "password": {
                    "type": "string",
                    "example": "password123"
                }
            }
        },
        "dtos.RegisterResponse": {
            "type": "object",
            "properties": {
                "data": {},
                "message": {
                    "type": "string",
                    "example": "Successfully Register New User"
                }
            }
        },
        "dtos.RentRequest": {
            "type": "object",
            "required": [
                "car_id",
                "duration"
            ],
            "properties": {
                "car_id": {
                    "type": "integer",
                    "example": 1
                },
                "duration": {
                    "type": "integer",
                    "example": 2
                }
            }
        },
        "dtos.RentResponse": {
            "type": "object",
            "properties": {
                "car_category": {
                    "type": "string",
                    "example": "SUV"
                },
                "car_name": {
                    "type": "string",
                    "example": "Mercedes AMG G63"
                },
                "deposit_amount": {
                    "type": "number",
                    "example": 10000
                },
                "end_date": {
                    "type": "string",
                    "example": "2024-02-03"
                },
                "invoce_url": {
                    "type": "string",
                    "example": "example_url"
                },
                "rental_id": {
                    "type": "integer",
                    "example": 1
                },
                "start_date": {
                    "type": "string",
                    "example": "2024-02-01"
                },
                "user_id": {
                    "type": "integer",
                    "example": 1
                }
            }
        },
        "dtos.RentalReportResponse": {
            "type": "object",
            "properties": {
                "car_category": {
                    "type": "string",
                    "example": "SUV"
                },
                "car_name": {
                    "type": "string",
                    "example": "Mercedes AMG G63"
                },
                "duration": {
                    "type": "integer",
                    "example": 2
                },
                "end_date": {
                    "type": "string",
                    "example": "2024-02-03"
                },
                "rental_id": {
                    "type": "integer",
                    "example": 1
                },
                "start_date": {
                    "type": "string",
                    "example": "2024-02-01"
                },
                "status": {
                    "type": "string",
                    "example": "Active"
                },
                "total_costs": {
                    "type": "number",
                    "example": 10000
                },
                "user_id": {
                    "type": "integer",
                    "example": 1
                }
            }
        },
        "dtos.TopUpRequest": {
            "type": "object",
            "required": [
                "deposit_amount"
            ],
            "properties": {
                "deposit_amount": {
                    "type": "number",
                    "example": 100000
                }
            }
        },
        "dtos.TopUpResponse": {
            "type": "object",
            "properties": {
                "deposit_amount": {
                    "type": "number",
                    "example": 100000
                },
                "message": {
                    "type": "string",
                    "example": "Successfully Top Up Balance"
                },
                "user_id": {
                    "type": "integer",
                    "example": 1
                }
            }
        },
        "models.Car": {
            "type": "object",
            "properties": {
                "car_id": {
                    "type": "integer",
                    "example": 1
                },
                "category": {
                    "type": "string",
                    "example": "SUV"
                },
                "name": {
                    "type": "string",
                    "example": "Mercedes AMG G63"
                },
                "rental_costs": {
                    "type": "number",
                    "example": 10000000
                },
                "stock_availability": {
                    "type": "integer",
                    "example": 3
                }
            }
        },
        "models.User": {
            "type": "object",
            "properties": {
                "depositAmount": {
                    "type": "number",
                    "example": 2000000
                },
                "email": {
                    "type": "string",
                    "example": "John.Doe@example.com"
                },
                "id": {
                    "type": "integer",
                    "example": 1
                },
                "name": {
                    "type": "string",
                    "example": "John Doe"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
