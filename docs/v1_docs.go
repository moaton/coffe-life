// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplatev1 = `{
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
        "/admin/categories": {
            "get": {
                "description": "Get all categories",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "categories"
                ],
                "summary": "Get categories",
                "responses": {
                    "200": {
                        "description": "categories",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/dto.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "payload": {
                                            "type": "array",
                                            "items": {
                                                "$ref": "#/definitions/dto.Category"
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/dto.Response"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/dto.Response"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/dto.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/dto.Response"
                        }
                    }
                }
            },
            "post": {
                "description": "Create category",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "categories"
                ],
                "summary": "Create category",
                "parameters": [
                    {
                        "description": "Category",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.CategoryRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "category",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/dto.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "payload": {
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
                            "$ref": "#/definitions/dto.Response"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/dto.Response"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/dto.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/dto.Response"
                        }
                    }
                }
            }
        },
        "/admin/categories/{id}": {
            "put": {
                "description": "Update category",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "categories"
                ],
                "summary": "Update category",
                "parameters": [
                    {
                        "type": "string",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Category",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.CategoryRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "category",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/dto.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "payload": {
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
                            "$ref": "#/definitions/dto.Response"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/dto.Response"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/dto.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/dto.Response"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete category",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "categories"
                ],
                "summary": "Delete category",
                "parameters": [
                    {
                        "type": "string",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "category",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/dto.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "payload": {
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
                            "$ref": "#/definitions/dto.Response"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/dto.Response"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/dto.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/dto.Response"
                        }
                    }
                }
            }
        },
        "/admin/foods": {
            "get": {
                "description": "Get all foods",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "foods"
                ],
                "summary": "Get foods",
                "responses": {
                    "200": {
                        "description": "foods",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/dto.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "payload": {
                                            "type": "array",
                                            "items": {
                                                "$ref": "#/definitions/dto.Food"
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/dto.Response"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/dto.Response"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/dto.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/dto.Response"
                        }
                    }
                }
            },
            "post": {
                "description": "Create food",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "foods"
                ],
                "summary": "Create food",
                "parameters": [
                    {
                        "description": "Food",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.FoodRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "food",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/dto.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "payload": {
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
                            "$ref": "#/definitions/dto.Response"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/dto.Response"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/dto.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/dto.Response"
                        }
                    }
                }
            }
        },
        "/admin/foods/{id}": {
            "put": {
                "description": "Update food",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "foods"
                ],
                "summary": "Update food",
                "parameters": [
                    {
                        "type": "string",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Food",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.FoodRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "food",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/dto.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "payload": {
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
                            "$ref": "#/definitions/dto.Response"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/dto.Response"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/dto.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/dto.Response"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete food",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "foods"
                ],
                "summary": "Delete food",
                "parameters": [
                    {
                        "type": "string",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "food",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/dto.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "payload": {
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
                            "$ref": "#/definitions/dto.Response"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/dto.Response"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/dto.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/dto.Response"
                        }
                    }
                }
            }
        },
        "/admin/login": {
            "post": {
                "description": "Login",
                "tags": [
                    "auth"
                ],
                "summary": "Login",
                "parameters": [
                    {
                        "description": "Login",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.LoginRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "login",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/dto.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "payload": {
                                            "$ref": "#/definitions/dto.LoginResponse"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/dto.Response"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/dto.Response"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/dto.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/dto.Response"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "dto.Category": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "dto.CategoryRequest": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "dto.Composition": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                },
                "unit": {
                    "type": "string"
                },
                "weight": {
                    "type": "integer"
                }
            }
        },
        "dto.Food": {
            "type": "object",
            "properties": {
                "category": {
                    "type": "string"
                },
                "compositions": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/dto.Composition"
                    }
                },
                "description": {
                    "type": "string"
                },
                "error_reason": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "is_new": {
                    "type": "boolean"
                },
                "is_spicy": {
                    "type": "boolean"
                },
                "name": {
                    "type": "string"
                },
                "price": {
                    "type": "string"
                },
                "type": {
                    "type": "string"
                }
            }
        },
        "dto.FoodRequest": {
            "type": "object",
            "properties": {
                "category": {
                    "type": "string"
                },
                "composition": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/dto.Composition"
                    }
                },
                "description": {
                    "type": "string"
                },
                "error_reason": {
                    "type": "string"
                },
                "is_new": {
                    "type": "boolean"
                },
                "is_spicy": {
                    "type": "boolean"
                },
                "name": {
                    "type": "string"
                },
                "price": {
                    "type": "string"
                },
                "type": {
                    "type": "string"
                }
            }
        },
        "dto.LoginRequest": {
            "type": "object",
            "properties": {
                "password": {
                    "type": "string",
                    "format": "base64"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "dto.LoginResponse": {
            "type": "object",
            "properties": {
                "token": {
                    "type": "string"
                }
            }
        },
        "dto.Response": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "errorCode": {
                    "type": "string"
                },
                "meta": {
                    "type": "object"
                },
                "payload": {
                    "type": "object"
                }
            }
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfov1 holds exported Swagger Info so clients can modify it
var SwaggerInfov1 = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "Coffe-Life API",
	Description:      "API for processing clients data",
	InfoInstanceName: "v1",
	SwaggerTemplate:  docTemplatev1,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfov1.InstanceName(), SwaggerInfov1)
}
