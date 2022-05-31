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
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "jose aranciba",
            "email": "josearanciba09@gmail.com"
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
        "/user/wallet/:walletid": {
            "get": {
                "description": "Finds in DB a wallet with the provided walletid Parameter",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Gets a user´s wallet with the wallet´s Id",
                "parameters": [
                    {
                        "maxLength": 50,
                        "minLength": 1,
                        "type": "string",
                        "description": "string valid",
                        "name": "walletid",
                        "in": "path"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/wallet.Wallet"
                            }
                        }
                    }
                }
            }
        },
        "/user/wallets": {
            "get": {
                "description": "Finds every wallet in DB for the logged user using session parameters",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Gets the user´s wallets with their corresponding data",
                "parameters": [
                    {
                        "maxLength": 50,
                        "minLength": 1,
                        "type": "string",
                        "description": "string valid",
                        "name": "walletid",
                        "in": "path"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "type": "array",
                                "items": {
                                    "$ref": "#/definitions/wallet.Wallet"
                                }
                            }
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "sql.NullTime": {
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
        "wallet.BallanceResponse": {
            "type": "object",
            "required": [
                "currency",
                "name"
            ],
            "properties": {
                "currency": {
                    "type": "string"
                },
                "name": {
                    "type": "number"
                }
            }
        },
        "wallet.Wallet": {
            "type": "object",
            "required": [
                "ballance",
                "created",
                "currency",
                "deleted",
                "name",
                "updated",
                "walletId"
            ],
            "properties": {
                "ballance": {
                    "type": "number"
                },
                "created": {
                    "$ref": "#/definitions/sql.NullTime"
                },
                "currency": {
                    "type": "string"
                },
                "deleted": {
                    "$ref": "#/definitions/sql.NullTime"
                },
                "name": {
                    "type": "string"
                },
                "updated": {
                    "$ref": "#/definitions/sql.NullTime"
                },
                "userId": {
                    "type": "string"
                },
                "walletId": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "BasicAuth": {
            "type": "basic"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:3010",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Wallet API",
	Description:      "This is a simple wallet API. you can handle users, wallets and transactions between them",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}