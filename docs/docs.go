// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag

package docs

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{.Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "license": {
            "name": "GPLv3",
            "url": "https://github.com/gibigo/cornix-tv-channel/blob/master/LICENSE"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/channels": {
            "get": {
                "security": [
                    {
                        "BasicAuth": []
                    }
                ],
                "description": "Get all channels of the current user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "channels"
                ],
                "summary": "Get all channels",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/types.Channel"
                            }
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/utils.HTTPError"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "BasicAuth": []
                    }
                ],
                "description": "Create a new channel",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "channels"
                ],
                "summary": "Create a channel",
                "parameters": [
                    {
                        "description": "Channel to create",
                        "name": "channel",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/types.AddChannel"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/types.Channel"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "409": {
                        "description": "Conflict",
                        "schema": {
                            "$ref": "#/definitions/utils.HTTPError"
                        }
                    }
                }
            }
        },
        "/channels/{channel_id}": {
            "get": {
                "security": [
                    {
                        "BasicAuth": []
                    }
                ],
                "description": "Get a spectific channel",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "channels"
                ],
                "summary": "Get a spectific channel",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Channel ID",
                        "name": "channel_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/types.Channel"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/utils.HTTPError"
                        }
                    }
                }
            },
            "put": {
                "security": [
                    {
                        "BasicAuth": []
                    }
                ],
                "description": "Change the telegram id of a channel and keep all related strategies",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "channels"
                ],
                "summary": "Update a channel",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Channel ID",
                        "name": "channel_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Channel to create",
                        "name": "channel",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/types.UpdateChannel"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/types.Channel"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/utils.HTTPError"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "BasicAuth": []
                    }
                ],
                "description": "Delete a channel and all related strategies",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "channels"
                ],
                "summary": "Delete a channel",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Channel ID",
                        "name": "channel_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/utils.HTTPError"
                        }
                    }
                }
            }
        },
        "/channels/{channel_id}/strategies": {
            "get": {
                "security": [
                    {
                        "BasicAuth": []
                    }
                ],
                "description": "Get all strategies for a particular channel",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "strategies"
                ],
                "summary": "Get all strategies",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Channel ID",
                        "name": "channel_id",
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
                                "$ref": "#/definitions/types.Strategy"
                            }
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/utils.HTTPError"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "BasicAuth": []
                    }
                ],
                "description": "Create a new strategy for the current user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "strategies"
                ],
                "summary": "Create a new strategy",
                "parameters": [
                    {
                        "description": "Strategy to create",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/types.AddStrategy"
                        }
                    },
                    {
                        "type": "integer",
                        "description": "Channel ID",
                        "name": "channel_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/types.Strategy"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/utils.HTTPError"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/channels/{channel_id}/strategies/{strategy_symbol}": {
            "get": {
                "security": [
                    {
                        "BasicAuth": []
                    }
                ],
                "description": "Get a strategy by the channel id and the symbol",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "strategies"
                ],
                "summary": "Get a strategy",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Channel ID",
                        "name": "channel_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Strategy Symbol, use 'all' for the default strategy",
                        "name": "strategy_symbol",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/types.Strategy"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/utils.HTTPError"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "BasicAuth": []
                    }
                ],
                "description": "Delete a strategy for a particular symbol",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "strategies"
                ],
                "summary": "Delete a strategy",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Channel ID",
                        "name": "channel_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Strategy Symbol, use 'all' for the default strategy",
                        "name": "strategy_symbol",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/utils.HTTPError"
                        }
                    }
                }
            }
        },
        "/users": {
            "get": {
                "security": [
                    {
                        "BasicAuth": []
                    }
                ],
                "description": "Get the current user, can be used to verify the user exists",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Get the current user",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/types.GetUser"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "put": {
                "security": [
                    {
                        "BasicAuth": []
                    }
                ],
                "description": "Change the current users setting. The request body must contain either a new name or a new password. If both, the username and the password get changed.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Change the current users setting",
                "parameters": [
                    {
                        "description": "Userupdate",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/types.AddUser"
                        }
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/utils.HTTPError"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "post": {
                "description": "Create a new user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Create a user",
                "parameters": [
                    {
                        "description": "User to create",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/types.AddUser"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/types.GetUser"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "409": {
                        "description": "Conflict",
                        "schema": {
                            "$ref": "#/definitions/utils.HTTPError"
                        }
                    },
                    "501": {
                        "description": "if user registration is disabled on the server",
                        "schema": {
                            "$ref": "#/definitions/utils.HTTPError"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "BasicAuth": []
                    }
                ],
                "description": "Delete the current user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Delete the current user",
                "responses": {
                    "204": {
                        "description": "No Content",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/utils.HTTPError"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "types.AddChannel": {
            "type": "object",
            "required": [
                "telegramId"
            ],
            "properties": {
                "telegramId": {
                    "type": "integer"
                }
            }
        },
        "types.AddStrategy": {
            "type": "object",
            "required": [
                "symbol"
            ],
            "properties": {
                "allowCounter": {
                    "type": "boolean"
                },
                "leverage": {
                    "type": "integer"
                },
                "symbol": {
                    "type": "string"
                },
                "targetStrategy": {
                    "$ref": "#/definitions/types.TargetStrategy"
                },
                "zoneStrategy": {
                    "$ref": "#/definitions/types.ZoneStrategy"
                }
            }
        },
        "types.AddUser": {
            "type": "object",
            "required": [
                "name",
                "password"
            ],
            "properties": {
                "name": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "types.Channel": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "telegramId": {
                    "type": "integer"
                }
            }
        },
        "types.Entry": {
            "type": "object",
            "properties": {
                "diff": {
                    "type": "number"
                }
            }
        },
        "types.GetUser": {
            "type": "object",
            "properties": {
                "identifier": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "types.SL": {
            "type": "object",
            "properties": {
                "diff": {
                    "type": "number"
                }
            }
        },
        "types.Strategy": {
            "type": "object",
            "properties": {
                "allowCounter": {
                    "type": "boolean"
                },
                "leverage": {
                    "type": "integer"
                },
                "symbol": {
                    "type": "string"
                },
                "targetStrategy": {
                    "$ref": "#/definitions/types.TargetStrategy"
                },
                "zoneStrategy": {
                    "$ref": "#/definitions/types.ZoneStrategy"
                }
            }
        },
        "types.TP": {
            "type": "object",
            "properties": {
                "diff": {
                    "type": "number"
                }
            }
        },
        "types.TargetStrategy": {
            "type": "object",
            "properties": {
                "entries": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/types.Entry"
                    }
                },
                "isBreakout": {
                    "type": "boolean"
                },
                "sl": {
                    "$ref": "#/definitions/types.SL"
                },
                "tps": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/types.TP"
                    }
                }
            }
        },
        "types.UpdateChannel": {
            "type": "object",
            "required": [
                "telegramId"
            ],
            "properties": {
                "telegramId": {
                    "type": "integer"
                }
            }
        },
        "types.ZoneStrategy": {
            "type": "object",
            "properties": {
                "entryStart": {
                    "type": "number"
                },
                "entryStop": {
                    "type": "number"
                },
                "isBreakout": {
                    "type": "boolean"
                },
                "sl": {
                    "$ref": "#/definitions/types.SL"
                },
                "tps": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/types.TP"
                    }
                }
            }
        },
        "utils.HTTPError": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer",
                    "example": 400
                },
                "message": {
                    "type": "string",
                    "example": "error"
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

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "1.0",
	Host:        "https://yourforwarder.io",
	BasePath:    "/api/v1",
	Schemes:     []string{},
	Title:       "Cornix-TV-Channel API",
	Description: "",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}
