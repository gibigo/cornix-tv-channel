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
            "name": "MIT",
            "url": "https://github.com/gibigo/cornix-tv-channel/blob/master/LICENSE"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/strategies": {
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
        "dal.Entry": {
            "type": "object",
            "required": [
                "diff"
            ],
            "properties": {
                "diff": {
                    "type": "number"
                },
                "targetStrategyID": {
                    "type": "integer"
                },
                "zoneStrategyID": {
                    "type": "integer"
                }
            }
        },
        "dal.SL": {
            "type": "object",
            "required": [
                "diff"
            ],
            "properties": {
                "diff": {
                    "type": "number"
                },
                "targetStrategyID": {
                    "type": "integer"
                },
                "zoneStrategyID": {
                    "type": "integer"
                }
            }
        },
        "dal.TP": {
            "type": "object",
            "required": [
                "diff"
            ],
            "properties": {
                "diff": {
                    "type": "number"
                },
                "targetStrategyID": {
                    "type": "integer"
                },
                "zoneStrategyID": {
                    "type": "integer"
                }
            }
        },
        "types.AddStrategy": {
            "type": "object",
            "properties": {
                "allowCounter": {
                    "type": "boolean"
                },
                "entires": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/dal.Entry"
                    }
                },
                "sl": {
                    "$ref": "#/definitions/dal.SL"
                },
                "tps": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/dal.TP"
                    }
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
        "types.Entry": {
            "type": "object",
            "properties": {
                "diff": {
                    "type": "number"
                },
                "strategyID": {
                    "description": "maybe remove this",
                    "type": "integer"
                }
            }
        },
        "types.SL": {
            "type": "object",
            "properties": {
                "diff": {
                    "type": "number"
                },
                "strategyID": {
                    "type": "integer"
                }
            }
        },
        "types.Strategy": {
            "type": "object",
            "properties": {
                "allowCounter": {
                    "type": "boolean"
                },
                "entires": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/types.Entry"
                    }
                },
                "id": {
                    "type": "integer"
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
        "types.TP": {
            "type": "object",
            "properties": {
                "diff": {
                    "type": "number"
                },
                "strategyID": {
                    "type": "integer"
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
