// Code generated by swaggo/swag. DO NOT EDIT.

package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "https://momentumindiana.org/statement-of-faith/",
        "contact": {
            "name": "Mwai Banda",
            "email": "bandamwai@gmail.com"
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
        "/api/v1/meals": {
            "get": {
                "description": "Used to get all meals",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Meals"
                ],
                "summary": "Get all Meals",
                "parameters": [
                    {
                        "type": "string",
                        "example": "Bearer XXX-xxx-XXX-xxx-XX",
                        "description": "Provide a bearer token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.MealResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/httputil.HTTPError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/httputil.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/httputil.HTTPError"
                        }
                    }
                }
            },
            "post": {
                "description": "Used to post a meal",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Meals"
                ],
                "summary": "Post a Meal",
                "parameters": [
                    {
                        "type": "string",
                        "example": "Bearer XXX-xxx-XXX-xxx-XX",
                        "description": "Provide a bearer token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "Post a meal",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.MealRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.MealRequest"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/httputil.HTTPError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/httputil.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/httputil.HTTPError"
                        }
                    }
                }
            }
        },
        "/api/v1/meals/meal": {
            "post": {
                "description": "Used to post a meal participate",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Meals"
                ],
                "summary": "Post a meal participate",
                "parameters": [
                    {
                        "type": "string",
                        "example": "Bearer XXX-xxx-XXX-xxx-XX",
                        "description": "Provide a bearer token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "Post participant information",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.VolunteeredRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.VolunteeredRequest"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/httputil.HTTPError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/httputil.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/httputil.HTTPError"
                        }
                    }
                }
            }
        },
        "/api/v1/meals/participant": {
            "post": {
                "description": "Used to post a meal participate",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Meals"
                ],
                "summary": "Post a meal participate",
                "parameters": [
                    {
                        "type": "string",
                        "example": "Bearer XXX-xxx-XXX-xxx-XX",
                        "description": "Provide a bearer token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "Post participant information",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.ParticipantRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.ParticipantRequest"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/httputil.HTTPError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/httputil.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/httputil.HTTPError"
                        }
                    }
                }
            }
        },
        "/api/v1/payments": {
            "post": {
                "description": "Used to initiate stripe transactions",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Payments"
                ],
                "summary": "Post a payment",
                "parameters": [
                    {
                        "type": "string",
                        "example": "Bearer XXX-xxx-XXX-xxx-XX",
                        "description": "Provide a bearer token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "Post a transaction",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.PaymentRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.PaymentResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/httputil.HTTPError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/httputil.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/httputil.HTTPError"
                        }
                    }
                }
            }
        },
        "/api/v1/sermons": {
            "get": {
                "description": "get a list of sermons",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Sermons"
                ],
                "summary": "Show a list of sermons",
                "parameters": [
                    {
                        "type": "string",
                        "example": "Bearer XXX-xxx-XXX-xxx-XX",
                        "description": "Provide a bearer token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "number",
                        "description": "Provide a page number",
                        "name": "page",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.SermonResponse"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/httputil.HTTPError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/httputil.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/httputil.HTTPError"
                        }
                    }
                }
            }
        },
        "/api/v1/users": {
            "put": {
                "description": "Used to update a user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Update a user",
                "parameters": [
                    {
                        "type": "string",
                        "example": "Bearer XXX-xxx-XXX-xxx-XX",
                        "description": "Provide a bearer token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "Post a user",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.UserRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.UserRequest"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/httputil.HTTPError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/httputil.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/httputil.HTTPError"
                        }
                    }
                }
            },
            "post": {
                "description": "Used to post a user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Post a user",
                "parameters": [
                    {
                        "type": "string",
                        "example": "Bearer XXX-xxx-XXX-xxx-XX",
                        "description": "Provide a bearer token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "Post a user",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.UserRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.UserRequest"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/httputil.HTTPError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/httputil.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/httputil.HTTPError"
                        }
                    }
                }
            }
        },
        "/api/v1/users/{userId}": {
            "get": {
                "description": "get a user information by providing an Id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Get user information by Id",
                "parameters": [
                    {
                        "type": "string",
                        "example": "Bearer XXX-xxx-XXX-xxx-XX",
                        "description": "Provide a bearer token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "provide user Id",
                        "name": "userId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.UserResponse"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/httputil.HTTPError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/httputil.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/httputil.HTTPError"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete a user's information by providing an Id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Delete user information by Id",
                "parameters": [
                    {
                        "type": "string",
                        "example": "Bearer XXX-xxx-XXX-xxx-XX",
                        "description": "Provide a bearer token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "provide user Id",
                        "name": "userId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.UserResponse"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/httputil.HTTPError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/httputil.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/httputil.HTTPError"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "httputil.HTTPError": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer",
                    "example": 400
                },
                "message": {
                    "type": "string",
                    "example": "status bad request"
                }
            }
        },
        "model.MealRequest": {
            "type": "object",
            "properties": {
                "allergies": {
                    "type": "string"
                },
                "city": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "favourites": {
                    "type": "string"
                },
                "instructions": {
                    "type": "string"
                },
                "least_favourites": {
                    "type": "string"
                },
                "num_of_adults": {
                    "type": "integer"
                },
                "number_of_kids": {
                    "type": "integer"
                },
                "phone": {
                    "type": "string"
                },
                "preferred_time": {
                    "type": "string"
                },
                "reason": {
                    "type": "string"
                },
                "recipient": {
                    "type": "string"
                },
                "street": {
                    "type": "string"
                },
                "user_id": {
                    "type": "string"
                }
            }
        },
        "model.MealResponse": {
            "type": "object",
            "properties": {
                "allergies": {
                    "type": "string"
                },
                "city": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "favourites": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "instructions": {
                    "type": "string"
                },
                "least_favourites": {
                    "type": "string"
                },
                "meals": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.VolunteeredMeal"
                    }
                },
                "num_of_adults": {
                    "type": "integer"
                },
                "number_of_kids": {
                    "type": "integer"
                },
                "participants": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.Participant"
                    }
                },
                "phone": {
                    "type": "string"
                },
                "preferred_time": {
                    "type": "string"
                },
                "reason": {
                    "type": "string"
                },
                "recipient": {
                    "type": "string"
                },
                "street": {
                    "type": "string"
                },
                "user": {
                    "$ref": "#/definitions/model.UserResponse"
                }
            }
        },
        "model.MealUpdates": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "model.Participant": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                },
                "user": {
                    "$ref": "#/definitions/model.UserResponse"
                }
            }
        },
        "model.ParticipantRequest": {
            "type": "object",
            "properties": {
                "meal_id": {
                    "type": "string"
                },
                "user_id": {
                    "type": "string"
                }
            }
        },
        "model.PaymentRequest": {
            "type": "object",
            "properties": {
                "amount": {
                    "type": "integer"
                },
                "email": {
                    "type": "string"
                },
                "fullname": {
                    "type": "string"
                },
                "phone": {
                    "type": "string"
                }
            }
        },
        "model.PaymentResponse": {
            "type": "object",
            "properties": {
                "customer": {
                    "type": "string"
                },
                "ephemeralKey": {
                    "type": "string"
                },
                "paymentIntent": {
                    "type": "string"
                },
                "publishableKey": {
                    "type": "string"
                }
            }
        },
        "model.Sermon": {
            "type": "object",
            "properties": {
                "categories": {},
                "date": {
                    "type": "object",
                    "properties": {
                        "carbon": {
                            "type": "string"
                        },
                        "date": {
                            "type": "string"
                        }
                    }
                },
                "detail_url": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "interactions": {
                    "type": "object",
                    "properties": {
                        "likes": {
                            "type": "integer"
                        },
                        "shares": {
                            "type": "integer"
                        }
                    }
                },
                "interactive_note": {},
                "keywords": {},
                "media": {
                    "type": "object",
                    "properties": {
                        "audio": {},
                        "embed": {},
                        "image": {},
                        "notes": {},
                        "video": {
                            "type": "string"
                        },
                        "video_thumbnail": {
                            "type": "string"
                        }
                    }
                },
                "passages": {
                    "type": "array",
                    "items": {}
                },
                "preacher": {
                    "type": "string"
                },
                "series": {
                    "type": "object",
                    "properties": {
                        "description": {},
                        "image": {},
                        "slug": {
                            "type": "string"
                        },
                        "title": {
                            "type": "string"
                        }
                    }
                },
                "slug": {
                    "type": "string"
                },
                "summary": {},
                "text": {},
                "title": {
                    "type": "string"
                }
            }
        },
        "model.SermonResponse": {
            "type": "object",
            "properties": {
                "current_page": {
                    "type": "integer"
                },
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.Sermon"
                    }
                },
                "first_page_url": {
                    "type": "string"
                },
                "from": {
                    "type": "integer"
                },
                "last_page": {
                    "type": "integer"
                },
                "last_page_url": {
                    "type": "string"
                },
                "links": {
                    "type": "array",
                    "items": {
                        "type": "object",
                        "properties": {
                            "active": {
                                "type": "boolean"
                            },
                            "label": {
                                "type": "string"
                            },
                            "url": {}
                        }
                    }
                },
                "next_page_url": {
                    "type": "string"
                },
                "path": {
                    "type": "string"
                },
                "per_page": {
                    "type": "integer"
                },
                "prev_page_url": {},
                "to": {
                    "type": "integer"
                },
                "total": {
                    "type": "string"
                }
            }
        },
        "model.UserRequest": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "fullname": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "phone": {
                    "type": "string"
                }
            }
        },
        "model.UserResponse": {
            "type": "object",
            "properties": {
                "createdOn": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "fullname": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "meals": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.MealResponse"
                    }
                },
                "phone": {
                    "type": "string"
                }
            }
        },
        "model.VolunteeredMeal": {
            "type": "object",
            "properties": {
                "created_on": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "notes": {
                    "type": "string"
                },
                "updates": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.MealUpdates"
                    }
                }
            }
        },
        "model.VolunteeredRequest": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "meal_id": {
                    "type": "string"
                },
                "notes": {
                    "type": "string"
                },
                "user_id": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "v1",
	Host:             "localhost:8085",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Momentum API",
	Description:      "This is a backend service to process payments and provide some CRUD functionality",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}