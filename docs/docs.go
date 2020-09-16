// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag at
// 2020-09-16 23:31:19.109258 +0800 CST m=+0.082074901

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
        "license": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/auth/AddAuth": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "summary": "新增个人用户",
                "parameters": [
                    {
                        "type": "string",
                        "description": "用户名",
                        "name": "username",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "密码",
                        "name": "password",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "姓名",
                        "name": "name",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "组：crm/calculate",
                        "name": "group",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "电话",
                        "name": "telephone",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":200,\"data\":{},\"msg\":\"ok\"}",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "{\"code\":500,\"data\":{},\"msg\":\"ok\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/exchange/addMyExchange": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "summary": "新增本人换班申请表信息",
                "parameters": [
                    {
                        "type": "string",
                        "description": "申请对象",
                        "name": "respondent",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "申请时间",
                        "name": "request_time",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "被申请时间",
                        "name": "requested_time",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "换班类型，1，晚班，2,周末白班，3，crm工作日特殊班，4，周末全天班",
                        "name": "exchange_type",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":200,\"data\":{},\"msg\":\"ok\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/exchange/deleteAll": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "summary": "清空所有的换班请求表",
                "responses": {
                    "200": {
                        "description": "{\"code\":200,\"data\":{},\"msg\":\"ok\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/exchange/deleteMyExchange{id}": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "summary": "删除本人的未审批换班请求表",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Id",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":200,\"data\":{},\"msg\":\"ok\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/exchange/examineExchange": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "summary": "回复换班申请表",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Id",
                        "name": "id",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "回复，状态 0为默认、1为拒绝、2为同意",
                        "name": "response",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":200,\"data\":{},\"msg\":\"ok\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/exchange/getAll": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "summary": "查看所有的换班请求表",
                "responses": {
                    "200": {
                        "description": "{\"code\":200,\"data\":{},\"msg\":\"ok\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/exchange/getMyExamine": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "summary": "查看本人回复的换班申请表信息(未审批/已审批)",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "状态 0：未审批 1：已审批",
                        "name": "state",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":200,\"data\":{},\"msg\":\"ok\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/exchange/myExchange": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "summary": "查看本人的换班请求表(未审批/已审批)",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "状态 0：未审批 1：已审批",
                        "name": "state",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":200,\"data\":{},\"msg\":\"ok\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/rests/addMyRest": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "summary": "新增本人调休申请表信息",
                "parameters": [
                    {
                        "type": "string",
                        "description": "申请调休时间",
                        "name": "request_time",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "申请调休类型，0：上午，1：下午，2：全天",
                        "name": "type",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "申请调休类型，0：调休，1：年休",
                        "name": "vacation_type",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":200,\"data\":{},\"msg\":\"ok\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/rests/deleteAll": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "summary": "删除所有调休申请表信息",
                "responses": {
                    "200": {
                        "description": "{\"code\":200,\"data\":{},\"msg\":\"ok\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/rests/deleteMyRest{id}": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "summary": "删除本人未审批调休申请表信息",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Id",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":200,\"data\":{},\"msg\":\"ok\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/rests/examineRest": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "summary": "审批调休申请表",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Id",
                        "name": "id",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "回复，状态 0为默认、1为拒绝、2为同意",
                        "name": "response",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":200,\"data\":{},\"msg\":\"ok\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/rests/getAll": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "summary": "获取所有调休申请表信息",
                "responses": {
                    "200": {
                        "description": "{\"code\":200,\"data\":{},\"msg\":\"ok\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/rests/getMe": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "summary": "获取本人调休申请表信息(未审批/已审批)",
                "parameters": [
                    {
                        "type": "integer",
                        "description": " 0 未审批、1 已审批",
                        "name": "state",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":200,\"data\":{rest},\"msg\":\"ok\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/rests/getNeedExamine": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "summary": "查看需要本人审核的未审核调休申请表信息",
                "responses": {
                    "200": {
                        "description": "{\"code\":200,\"data\":{},\"msg\":\"ok\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/rota/import": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "summary": "导入值班表",
                "parameters": [
                    {
                        "type": "string",
                        "description": "excel表格文件",
                        "name": "file",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":200,\"data\":{},\"msg\":\"ok\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/rotas/addDay": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "summary": "添加日值班表",
                "parameters": [
                    {
                        "type": "string",
                        "description": "日期",
                        "name": "datetime",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "星期",
                        "name": "week",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "计费晚班人员",
                        "name": "billing_late",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "计费周末白班人员",
                        "name": "billing_weekend_day",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "crm晚班人员",
                        "name": "crm_late",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "crm周末白班人员",
                        "name": "crm_weekend_day",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "crm工作日特殊班值班人员",
                        "name": "crm_duty_special",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":200,\"data\":{},\"msg\":\"ok\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/rotas/deleteDay{datetime}": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "summary": "删除日值班表",
                "parameters": [
                    {
                        "type": "string",
                        "description": "日期，例如：2020-09-01",
                        "name": "datetime",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":200,\"data\":{},\"msg\":\"ok\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/rotas/deleteMonth{month}": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "summary": "删除月值班表",
                "parameters": [
                    {
                        "type": "string",
                        "description": "月份，例如：2020-09",
                        "name": "month",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":200,\"data\":{},\"msg\":\"ok\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/rotas/getMonth{month}": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "summary": "获取当月值班表",
                "parameters": [
                    {
                        "type": "string",
                        "description": "月份，例如：2020-09",
                        "name": "month",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":200,\"data\":{},\"msg\":\"ok\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/vacation/deleteAll": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "summary": "清空所有调休信息",
                "responses": {
                    "200": {
                        "description": "{\"code\":200,\"data\":{},\"msg\":\"ok\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/vacation/deleteByName": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "summary": "删除某人调休信息",
                "parameters": [
                    {
                        "type": "string",
                        "description": "姓名",
                        "name": "name",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":200,\"data\":{},\"msg\":\"ok\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/vacation/editByName": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "summary": "修改某人调休信息",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "姓名",
                        "name": "name",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "number",
                        "description": "剩余调休",
                        "name": "remain_vacation",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "number",
                        "description": "剩余年休",
                        "name": "remain_annual_vacation",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":200,\"data\":{},\"msg\":\"ok\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/vacation/getAll": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "summary": "获取所有员工的调休信息",
                "responses": {
                    "200": {
                        "description": "{\"code\":200,\"data\":{},\"msg\":\"ok\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/auth": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "summary": "通过用户名密码获取token",
                "parameters": [
                    {
                        "type": "string",
                        "description": "用户名",
                        "name": "username",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "密码",
                        "name": "password",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":200,\"data\":{},\"msg\":\"ok\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/register": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "summary": "注册个人用户",
                "parameters": [
                    {
                        "type": "string",
                        "description": "用户名",
                        "name": "username",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "密码",
                        "name": "password",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "名字",
                        "name": "name",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "组：crm/calculate",
                        "name": "group",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "电话",
                        "name": "telephone",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":200,\"data\":{},\"msg\":\"ok\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
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
	Version:     "",
	Host:        "",
	BasePath:    "",
	Schemes:     []string{},
	Title:       "",
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
