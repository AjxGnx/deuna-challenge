// Package docs GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
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
        "license": {
            "name": "Alirio Gutierrez"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/health": {
            "get": {
                "description": "health service",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Health"
                ],
                "summary": "Check if service is active",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handler.Health"
                        }
                    }
                }
            }
        },
        "/transactions/stripe/intent": {
            "post": {
                "description": "Create a payment intent with stripe",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Transactions"
                ],
                "summary": "Create payment intent",
                "parameters": [
                    {
                        "description": "Request Body",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.Transaction"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/entity.Transaction"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/dto.MessageError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/dto.MessageError"
                        }
                    }
                }
            }
        },
        "/transactions/stripe/refund": {
            "post": {
                "description": "Create a refund with stripe (to run this method the transaction needs to be succeeded in stripe)",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Transactions"
                ],
                "summary": "Create refund",
                "parameters": [
                    {
                        "description": "Request Body",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.Refund"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/entity.Transaction"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/dto.MessageError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/dto.MessageError"
                        }
                    }
                }
            }
        },
        "/transactions/{id}": {
            "get": {
                "description": "Get transaction by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Transactions"
                ],
                "summary": "Get by ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "value of transaction to find",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/entity.Transaction"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/dto.MessageError"
                        }
                    }
                }
            },
            "put": {
                "description": "Update transaction by ID (to run this acction you need to know the state of a transaction \"payment intent\")",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Transactions"
                ],
                "summary": "Update By ID",
                "parameters": [
                    {
                        "description": "Request Body",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.UpdateTransaction"
                        }
                    },
                    {
                        "type": "integer",
                        "description": "value of transaction to update",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/entity.Transaction"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/dto.MessageError"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "dto.MessageError": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                }
            }
        },
        "dto.Refund": {
            "type": "object",
            "properties": {
                "transaction_id": {
                    "type": "integer"
                }
            }
        },
        "dto.Transaction": {
            "type": "object",
            "properties": {
                "amount": {
                    "type": "number"
                },
                "customer_id": {
                    "type": "integer"
                },
                "merchant_id": {
                    "type": "integer"
                }
            }
        },
        "dto.UpdateTransaction": {
            "type": "object",
            "properties": {
                "error_code": {
                    "type": "string"
                },
                "error_message": {
                    "type": "string"
                },
                "error_type": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "payment_method": {
                    "type": "string"
                },
                "status": {
                    "type": "string"
                }
            }
        },
        "entity.Transaction": {
            "type": "object",
            "properties": {
                "amount": {
                    "type": "number"
                },
                "client_secret": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "merchant_id": {
                    "type": "integer"
                },
                "payment_method": {
                    "type": "string"
                },
                "payment_source": {
                    "type": "string"
                },
                "status": {
                    "type": "string"
                },
                "type": {
                    "type": "string"
                }
            }
        },
        "handler.Health": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                },
                "status": {
                    "type": "integer"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0.0",
	Host:             "",
	BasePath:         "/api/exercise",
	Schemes:          []string{"http"},
	Title:            "Deuna Challenge",
	Description:      "Deuna Challenge Manager",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}