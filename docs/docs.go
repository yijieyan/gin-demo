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
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/user/getUserInfo": {
            "get": {
                "description": "用户获取个人信息接口",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "用户管理"
                ],
                "summary": "用户获取个人信息",
                "operationId": "/user/getUserInfo",
                "parameters": [
                    {
                        "type": "string",
                        "description": "header",
                        "name": "authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "success",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/middleware.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/dto.GetUserInfoOutput"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/user/login": {
            "post": {
                "description": "用户登录接口",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户管理"
                ],
                "summary": "用户登录",
                "operationId": "/user/login",
                "parameters": [
                    {
                        "description": "body",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.LoginInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "success",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/middleware.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/dto.LoginOutput"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/user/register": {
            "post": {
                "description": "用户注册接口",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户管理"
                ],
                "summary": "用户注册",
                "operationId": "/user/register",
                "parameters": [
                    {
                        "description": "body",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.RegisterInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "success",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/middleware.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "dto.GetUserInfoOutput": {
            "type": "object",
            "properties": {
                "account": {
                    "description": "账号",
                    "type": "string"
                },
                "age": {
                    "description": "年龄",
                    "type": "integer"
                },
                "avatar_url": {
                    "description": "头像",
                    "type": "string"
                },
                "birth": {
                    "description": "出生日期",
                    "type": "string"
                },
                "id": {
                    "description": "用户id",
                    "type": "string"
                },
                "name": {
                    "description": "用户名",
                    "type": "string"
                },
                "sex": {
                    "description": "性别 1：男 2：女",
                    "type": "string"
                }
            }
        },
        "dto.LoginInput": {
            "type": "object",
            "required": [
                "account",
                "password"
            ],
            "properties": {
                "account": {
                    "description": "账号",
                    "type": "string",
                    "maxLength": 22,
                    "minLength": 3
                },
                "password": {
                    "description": "密码",
                    "type": "string"
                }
            }
        },
        "dto.LoginOutput": {
            "type": "object",
            "properties": {
                "token": {
                    "description": "用户token",
                    "type": "string"
                }
            }
        },
        "dto.RegisterInput": {
            "type": "object",
            "required": [
                "account",
                "birth",
                "name",
                "sex"
            ],
            "properties": {
                "account": {
                    "description": "账号",
                    "type": "string"
                },
                "age": {
                    "description": "年龄",
                    "type": "integer"
                },
                "birth": {
                    "description": "出生日期",
                    "type": "string"
                },
                "name": {
                    "description": "用户名",
                    "type": "string"
                },
                "password": {
                    "description": "密码",
                    "type": "string",
                    "maxLength": 22,
                    "minLength": 6
                },
                "sex": {
                    "description": "性别 1：男 2：女",
                    "type": "integer"
                }
            }
        },
        "middleware.Response": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {},
                "message": {
                    "type": "string"
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
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
