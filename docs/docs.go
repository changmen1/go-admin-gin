// Package docs Code generated by swaggo/swag. DO NOT EDIT
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
        "/api/failed": {
            "get": {
                "description": "失败测试接口",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "测试相关接口"
                ],
                "summary": "失败测试",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/result.Result"
                        }
                    }
                }
            }
        },
        "/api/success": {
            "get": {
                "description": "成功测试接口",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "测试相关接口"
                ],
                "summary": "成功测试",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/result.Result"
                        }
                    }
                }
            }
        },
        "/api/sysMenu/add": {
            "post": {
                "description": "新增菜单",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "菜单相关接口"
                ],
                "parameters": [
                    {
                        "description": "data",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.AddSysMenuDto"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/result.Result"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "model.AddSysMenuDto": {
            "type": "object",
            "properties": {
                "icon": {
                    "type": "string"
                },
                "menuName": {
                    "type": "string"
                },
                "menuStatus": {
                    "type": "integer"
                },
                "menuType": {
                    "type": "integer"
                },
                "parentId": {
                    "type": "integer"
                },
                "sort": {
                    "type": "integer"
                },
                "url": {
                    "type": "string"
                },
                "value": {
                    "type": "string"
                }
            }
        },
        "result.Result": {
            "type": "object",
            "properties": {
                "code": {
                    "description": "状态码",
                    "type": "integer"
                },
                "data": {
                    "description": "返回的数据"
                },
                "message": {
                    "description": "提示信息",
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
	Title:            "中央监护系统",
	Description:      "中央监护系统API接口文档",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
