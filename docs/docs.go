// Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "SoulChild",
            "url": "http://soulchild.cn"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/ping": {
            "get": {
                "description": "do ping",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "core"
                ],
                "summary": "health check",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/core.resp"
                        }
                    }
                }
            }
        },
        "/system/user/info": {
            "get": {
                "description": "用户信息",
                "tags": [
                    "User"
                ],
                "summary": "用户信息",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authorization token",
                        "name": "X-Token",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "成功返回用户信息",
                        "schema": {
                            "$ref": "#/definitions/httputil.ResponseBody"
                        }
                    }
                }
            }
        },
        "/system/user/login": {
            "post": {
                "description": "用户登录",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "用户登录",
                "parameters": [
                    {
                        "description": "手机号,密码",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.SystemLogin"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "成功返回token",
                        "schema": {
                            "$ref": "#/definitions/httputil.ResponseBody"
                        }
                    }
                }
            }
        },
        "/system/user/register": {
            "post": {
                "description": "用户注册",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "用户注册",
                "parameters": [
                    {
                        "description": "用户信息",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.SystemRegister"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "成功返回",
                        "schema": {
                            "$ref": "#/definitions/httputil.ResponseBody"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "core.resp": {
            "type": "object",
            "properties": {
                "status": {
                    "type": "string",
                    "example": "pong"
                }
            }
        },
        "dto.SystemLogin": {
            "type": "object",
            "required": [
                "mobile",
                "password"
            ],
            "properties": {
                "mobile": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "dto.SystemRegister": {
            "type": "object",
            "required": [
                "mobile",
                "password",
                "username"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "mobile": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "httputil.ResponseBody": {
            "type": "object",
            "properties": {
                "data": {},
                "msg": {
                    "type": "string"
                },
                "status": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "type": "apiKey",
            "name": "x-token",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "0.0.1",
	Host:             "localhost:8080",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Swagger Example API",
	Description:      "This is a sample server celler server.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
