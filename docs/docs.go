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
        "/login": {
            "post": {
                "description": "Logs the user by running a sql query against the DB checking wether if a user with the provided username and password exists. If it exists, returns it and stores User Id in session",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Logs a user with its username and password",
                "parameters": [
                    {
                        "maxLength": 50,
                        "minLength": 1,
                        "description": "string valid",
                        "name": "name",
                        "in": "body",
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "maxLength": 50,
                        "minLength": 1,
                        "description": "string valid",
                        "name": "password",
                        "in": "body",
                        "schema": {
                            "type": "string"
                        }
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
        "/user": {
            "get": {
                "description": "Finds in DB a user with the userId Parameter provided by cookies",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Gets the logged user",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/user.User"
                            }
                        }
                    }
                }
            }
        },
        "/user/Ballance/:Currency": {
            "get": {
                "description": "Sums every user wallets ballances for the given Currency",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Gets a user´s current total ballance",
                "parameters": [
                    {
                        "maxLength": 50,
                        "minLength": 1,
                        "type": "string",
                        "description": "string valid",
                        "name": "currency",
                        "in": "path"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/wallet.BallanceResponse"
                            }
                        }
                    }
                }
            }
        },
        "/user/new": {
            "post": {
                "description": "Creates a user with given name, mail and (not implemented)password",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Creates a user",
                "parameters": [
                    {
                        "maxLength": 50,
                        "minLength": 1,
                        "description": "string valid",
                        "name": "name",
                        "in": "body",
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "maxLength": 50,
                        "minLength": 1,
                        "description": "string valid",
                        "name": "email",
                        "in": "body",
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "maxLength": 50,
                        "minLength": 1,
                        "description": "string valid",
                        "name": "password",
                        "in": "body",
                        "schema": {
                            "type": "string"
                        }
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
        "/user/transactions/:transactionid": {
            "get": {
                "description": "Finds in DB a transaction with the provided transactionid Parameter",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Gets a transaction by its id",
                "parameters": [
                    {
                        "maxLength": 50,
                        "minLength": 1,
                        "type": "string",
                        "description": "string valid",
                        "name": "transactionid",
                        "in": "path"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/transaction.Transaction"
                            }
                        }
                    }
                }
            }
        },
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
        "/user/wallet/:walletid/transactions": {
            "get": {
                "description": "Finds in DB a transaction with the provided walletid Parameter",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Gets all transactions of a certain wallet",
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
                                "$ref": "#/definitions/transaction.Transaction"
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
        },
        "/wallet/send": {
            "post": {
                "description": "before checking if the origin wallet has funds, it sends the declared amount to the destination wallet and creates a trnsaction",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Sends money from a wallet to another",
                "parameters": [
                    {
                        "maxLength": 50,
                        "minLength": 1,
                        "description": "string valid",
                        "name": "originId",
                        "in": "body",
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "maxLength": 50,
                        "minLength": 1,
                        "description": "string valid",
                        "name": "destinationId",
                        "in": "body",
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "description": "string valid",
                        "name": "amount",
                        "in": "body",
                        "schema": {
                            "type": "number"
                        }
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
        "transaction.Transaction": {
            "type": "object",
            "required": [
                "amount",
                "destinationId",
                "originId",
                "status",
                "transactionId"
            ],
            "properties": {
                "amount": {
                    "type": "number",
                    "maximum": 1000000,
                    "minimum": 1
                },
                "created": {
                    "$ref": "#/definitions/sql.NullTime"
                },
                "destinationId": {
                    "type": "string",
                    "maxLength": 50
                },
                "originId": {
                    "type": "string",
                    "maxLength": 50
                },
                "status": {
                    "type": "string",
                    "maxLength": 50
                },
                "transactionId": {
                    "type": "string",
                    "maxLength": 50
                }
            }
        },
        "user.User": {
            "type": "object",
            "required": [
                "created",
                "name",
                "password",
                "userId"
            ],
            "properties": {
                "created": {
                    "$ref": "#/definitions/sql.NullTime"
                },
                "deleted": {
                    "$ref": "#/definitions/sql.NullTime"
                },
                "email": {
                    "type": "string"
                },
                "name": {
                    "type": "string",
                    "maxLength": 50
                },
                "password": {
                    "type": "string",
                    "maxLength": 50
                },
                "userId": {
                    "type": "string",
                    "maxLength": 50
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
                "userId",
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
                    "type": "string",
                    "maxLength": 3
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
                    "type": "string",
                    "maxLength": 50
                },
                "walletId": {
                    "type": "string",
                    "maxLength": 50
                }
            }
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
