// Package docs GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import (
	"bytes"
	"encoding/json"
	"strings"
	"text/template"

	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "consumes": [
        "application/json",
        "application/x-www-form-urlencoded"
    ],
    "produces": [
        "application/json"
    ],
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
        "/v1/greet": {
            "get": {
                "security": [
                    {
                        "Authorization": []
                    }
                ],
                "description": "示例接口",
                "consumes": [
                    "application/x-www-form-urlencoded"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "示例"
                ],
                "summary": "示例接口",
                "parameters": [
                    {
                        "type": "string",
                        "format": "string",
                        "default": "Tom",
                        "description": "名称",
                        "name": "name",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "成功响应",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/example.Success"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/greet.HelloResponse"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "客户端请求错误（code 类型应为 int，string 仅为了表达多个错误码）",
                        "schema": {
                            "$ref": "#/definitions/example.ClientError"
                        }
                    },
                    "401": {
                        "description": "登陆失效",
                        "schema": {
                            "$ref": "#/definitions/example.Unauthorized"
                        }
                    },
                    "403": {
                        "description": "没有权限",
                        "schema": {
                            "$ref": "#/definitions/example.PermissionDenied"
                        }
                    },
                    "404": {
                        "description": "资源不存在",
                        "schema": {
                            "$ref": "#/definitions/example.ResourceNotFound"
                        }
                    },
                    "429": {
                        "description": "请求过于频繁",
                        "schema": {
                            "$ref": "#/definitions/example.TooManyRequest"
                        }
                    },
                    "500": {
                        "description": "服务器出错",
                        "schema": {
                            "$ref": "#/definitions/example.ServerError"
                        }
                    }
                }
            }
        },
        "/v1/trace": {
            "get": {
                "security": [
                    {
                        "Authorization": []
                    }
                ],
                "description": "示例接口",
                "consumes": [
                    "application/x-www-form-urlencoded"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "示例"
                ],
                "summary": "示例接口",
                "responses": {
                    "200": {
                        "description": "成功响应",
                        "schema": {
                            "$ref": "#/definitions/example.Success"
                        }
                    },
                    "400": {
                        "description": "客户端请求错误（code 类型应为 int，string 仅为了表达多个错误码）",
                        "schema": {
                            "$ref": "#/definitions/example.ClientError"
                        }
                    },
                    "401": {
                        "description": "登陆失效",
                        "schema": {
                            "$ref": "#/definitions/example.Unauthorized"
                        }
                    },
                    "403": {
                        "description": "没有权限",
                        "schema": {
                            "$ref": "#/definitions/example.PermissionDenied"
                        }
                    },
                    "404": {
                        "description": "资源不存在",
                        "schema": {
                            "$ref": "#/definitions/example.ResourceNotFound"
                        }
                    },
                    "429": {
                        "description": "请求过于频繁",
                        "schema": {
                            "$ref": "#/definitions/example.TooManyRequest"
                        }
                    },
                    "500": {
                        "description": "服务器出错",
                        "schema": {
                            "$ref": "#/definitions/example.ServerError"
                        }
                    }
                }
            }
        },
        "/v1/user": {
            "post": {
                "security": [
                    {
                        "Authorization": []
                    }
                ],
                "description": "创建用户",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户"
                ],
                "summary": "创建用户",
                "parameters": [
                    {
                        "format": "string",
                        "description": "用户信息",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/user.CreateRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "成功响应",
                        "schema": {
                            "$ref": "#/definitions/example.Success"
                        }
                    },
                    "400": {
                        "description": "客户端请求错误（code 类型应为 int，string 仅为了表达多个错误码）",
                        "schema": {
                            "$ref": "#/definitions/example.ClientError"
                        }
                    },
                    "401": {
                        "description": "登陆失效",
                        "schema": {
                            "$ref": "#/definitions/example.Unauthorized"
                        }
                    },
                    "403": {
                        "description": "没有权限",
                        "schema": {
                            "$ref": "#/definitions/example.PermissionDenied"
                        }
                    },
                    "404": {
                        "description": "资源不存在",
                        "schema": {
                            "$ref": "#/definitions/example.ResourceNotFound"
                        }
                    },
                    "429": {
                        "description": "请求过于频繁",
                        "schema": {
                            "$ref": "#/definitions/example.TooManyRequest"
                        }
                    },
                    "500": {
                        "description": "服务器出错",
                        "schema": {
                            "$ref": "#/definitions/example.ServerError"
                        }
                    }
                }
            }
        },
        "/v1/user/{id}": {
            "get": {
                "security": [
                    {
                        "Authorization": []
                    }
                ],
                "description": "用户详情",
                "consumes": [
                    "text/plain"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户"
                ],
                "summary": "用户详情",
                "parameters": [
                    {
                        "minimum": 1,
                        "type": "integer",
                        "format": "uint",
                        "description": "用户 id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "成功响应",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/example.Success"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/user.DetailResponse"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "客户端请求错误（code 类型应为 int，string 仅为了表达多个错误码）",
                        "schema": {
                            "$ref": "#/definitions/example.ClientError"
                        }
                    },
                    "401": {
                        "description": "登陆失效",
                        "schema": {
                            "$ref": "#/definitions/example.Unauthorized"
                        }
                    },
                    "403": {
                        "description": "没有权限",
                        "schema": {
                            "$ref": "#/definitions/example.PermissionDenied"
                        }
                    },
                    "404": {
                        "description": "资源不存在",
                        "schema": {
                            "$ref": "#/definitions/example.ResourceNotFound"
                        }
                    },
                    "429": {
                        "description": "请求过于频繁",
                        "schema": {
                            "$ref": "#/definitions/example.TooManyRequest"
                        }
                    },
                    "500": {
                        "description": "服务器出错",
                        "schema": {
                            "$ref": "#/definitions/example.ServerError"
                        }
                    }
                }
            },
            "put": {
                "security": [
                    {
                        "Authorization": []
                    }
                ],
                "description": "更新用户",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户"
                ],
                "summary": "更新用户",
                "parameters": [
                    {
                        "minimum": 1,
                        "type": "integer",
                        "format": "uint",
                        "description": "用户 id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "format": "string",
                        "description": "用户信息",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/user.CreateRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "成功响应",
                        "schema": {
                            "$ref": "#/definitions/example.Success"
                        }
                    },
                    "400": {
                        "description": "客户端请求错误（code 类型应为 int，string 仅为了表达多个错误码）",
                        "schema": {
                            "$ref": "#/definitions/example.ClientError"
                        }
                    },
                    "401": {
                        "description": "登陆失效",
                        "schema": {
                            "$ref": "#/definitions/example.Unauthorized"
                        }
                    },
                    "403": {
                        "description": "没有权限",
                        "schema": {
                            "$ref": "#/definitions/example.PermissionDenied"
                        }
                    },
                    "404": {
                        "description": "资源不存在",
                        "schema": {
                            "$ref": "#/definitions/example.ResourceNotFound"
                        }
                    },
                    "429": {
                        "description": "请求过于频繁",
                        "schema": {
                            "$ref": "#/definitions/example.TooManyRequest"
                        }
                    },
                    "500": {
                        "description": "服务器出错",
                        "schema": {
                            "$ref": "#/definitions/example.ServerError"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "Authorization": []
                    }
                ],
                "description": "删除用户",
                "consumes": [
                    "text/plain"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户"
                ],
                "summary": "删除用户",
                "parameters": [
                    {
                        "minimum": 1,
                        "type": "integer",
                        "format": "uint",
                        "description": "用户 id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "成功响应",
                        "schema": {
                            "$ref": "#/definitions/example.Success"
                        }
                    },
                    "400": {
                        "description": "客户端请求错误（code 类型应为 int，string 仅为了表达多个错误码）",
                        "schema": {
                            "$ref": "#/definitions/example.ClientError"
                        }
                    },
                    "401": {
                        "description": "登陆失效",
                        "schema": {
                            "$ref": "#/definitions/example.Unauthorized"
                        }
                    },
                    "403": {
                        "description": "没有权限",
                        "schema": {
                            "$ref": "#/definitions/example.PermissionDenied"
                        }
                    },
                    "404": {
                        "description": "资源不存在",
                        "schema": {
                            "$ref": "#/definitions/example.ResourceNotFound"
                        }
                    },
                    "429": {
                        "description": "请求过于频繁",
                        "schema": {
                            "$ref": "#/definitions/example.TooManyRequest"
                        }
                    },
                    "500": {
                        "description": "服务器出错",
                        "schema": {
                            "$ref": "#/definitions/example.ServerError"
                        }
                    }
                }
            }
        },
        "/v1/users": {
            "get": {
                "security": [
                    {
                        "Authorization": []
                    }
                ],
                "description": "用户列表",
                "consumes": [
                    "application/x-www-form-urlencoded"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户"
                ],
                "summary": "用户列表",
                "parameters": [
                    {
                        "type": "string",
                        "format": "string",
                        "description": "查询字符串",
                        "name": "keyword",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "成功响应",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/example.Success"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "array",
                                            "items": {
                                                "$ref": "#/definitions/user.ListItem"
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "客户端请求错误（code 类型应为 int，string 仅为了表达多个错误码）",
                        "schema": {
                            "$ref": "#/definitions/example.ClientError"
                        }
                    },
                    "401": {
                        "description": "登陆失效",
                        "schema": {
                            "$ref": "#/definitions/example.Unauthorized"
                        }
                    },
                    "403": {
                        "description": "没有权限",
                        "schema": {
                            "$ref": "#/definitions/example.PermissionDenied"
                        }
                    },
                    "404": {
                        "description": "资源不存在",
                        "schema": {
                            "$ref": "#/definitions/example.ResourceNotFound"
                        }
                    },
                    "429": {
                        "description": "请求过于频繁",
                        "schema": {
                            "$ref": "#/definitions/example.TooManyRequest"
                        }
                    },
                    "500": {
                        "description": "服务器出错",
                        "schema": {
                            "$ref": "#/definitions/example.ServerError"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "example.ClientError": {
            "type": "object",
            "properties": {
                "code": {
                    "description": "code 类型应为 int，string 仅为了表达多个错误码",
                    "type": "string",
                    "example": "10002|10003"
                },
                "msg": {
                    "type": "string",
                    "example": "客户端请求错误|参数校验错误"
                }
            }
        },
        "example.PermissionDenied": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer",
                    "example": 10005
                },
                "msg": {
                    "type": "string",
                    "example": "暂无权限"
                }
            }
        },
        "example.ResourceNotFound": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer",
                    "example": 10006
                },
                "msg": {
                    "type": "string",
                    "example": "资源不存在"
                }
            }
        },
        "example.ServerError": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer",
                    "example": 10001
                },
                "msg": {
                    "type": "string",
                    "example": "服务器出错"
                }
            }
        },
        "example.Success": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer",
                    "example": 10000
                },
                "data": {},
                "msg": {
                    "type": "string",
                    "example": "OK"
                }
            }
        },
        "example.TooManyRequest": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer",
                    "example": 10007
                },
                "msg": {
                    "type": "string",
                    "example": "请求过于频繁"
                }
            }
        },
        "example.Unauthorized": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer",
                    "example": 10004
                },
                "msg": {
                    "type": "string",
                    "example": "未经授权"
                }
            }
        },
        "greet.HelloResponse": {
            "type": "object",
            "properties": {
                "msg": {
                    "type": "string"
                }
            }
        },
        "user.CreateRequest": {
            "type": "object",
            "properties": {
                "age": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "phone": {
                    "type": "string"
                }
            }
        },
        "user.DetailResponse": {
            "type": "object",
            "properties": {
                "age": {
                    "type": "integer"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "phone": {
                    "type": "string"
                }
            }
        },
        "user.ListItem": {
            "type": "object",
            "properties": {
                "age": {
                    "type": "integer"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "phone": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "Authorization": {
            "type": "apiKey",
            "name": "Token",
            "in": "header"
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
	Version:     "0.0.0",
	Host:        "localhost",
	BasePath:    "/api",
	Schemes:     []string{"http", "https"},
	Title:       "API 接口文档",
	Description: "API 接口文档",
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
		"escape": func(v interface{}) string {
			// escape tabs
			str := strings.Replace(v.(string), "\t", "\\t", -1)
			// replace " with \", and if that results in \\", replace that with \\\"
			str = strings.Replace(str, "\"", "\\\"", -1)
			return strings.Replace(str, "\\\\\"", "\\\\\\\"", -1)
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
	swag.Register("swagger", &s{})
}
