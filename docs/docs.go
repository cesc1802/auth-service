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
            "name": "Cesc Nguyen",
            "email": "thuocnv1802@gmail.com"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/v1/roles": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Create Role",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Roles"
                ],
                "summary": "Create Role",
                "parameters": [
                    {
                        "description": "Create Role",
                        "name": "role",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.CreateRoleRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/common.AppError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/common.AppError"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "common.AppError": {
            "type": "object",
            "properties": {
                "error_key": {
                    "type": "string"
                },
                "log": {
                    "type": "string"
                },
                "message": {
                    "type": "string"
                },
                "status_code": {
                    "type": "integer"
                },
                "ve": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/common.ValidationErrorField"
                    }
                }
            }
        },
        "common.ValidationErrorField": {
            "type": "object",
            "properties": {
                "errorMessage": {
                    "type": "string"
                },
                "field": {
                    "type": "string"
                },
                "tag": {
                    "type": "string"
                }
            }
        },
        "dto.CreateRoleRequest": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
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

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Auth Service API",
	Description:      "This is Auth Service API.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}