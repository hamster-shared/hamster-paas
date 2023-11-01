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
        "/api/v2/zan/account/access_token": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "用户在zan平台授权hamster访问后，zan平台通过跳转方式返回给前端authCode, 此时需要调用后端接口与zan平台交换成可以访问的access_token\n前端需要从url中解析这个authCode，并调用此接口交换\n用户第一次交换成功后，后续不需要再此交换，此token由后台保存",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "zan"
                ],
                "summary": "交换zan的access_token",
                "parameters": [
                    {
                        "description": "请求交换token参数",
                        "name": "exchangeAccessTokenVo",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/vo.ExchangeAccessTokenVo"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handler.Result"
                        }
                    }
                }
            }
        },
        "/api/v2/zan/account/auth_url": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "获取授权链接接⼝\n返回结果中的data 是前端需要跳转到zan平台进行认证的url\nzan平台认证完成后，会跳转回前端，如",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "zan"
                ],
                "summary": "获取授权链接接⼝",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handler.Result"
                        }
                    }
                }
            }
        },
        "/api/v2/zan/account/authed": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "判断用户是否已经进行了zan授权，如果未授权，需要进行授权流程\n授权流程\n1. 调用 获取授权链接接⼝ /api/v2/zan/auth_url\n2. 前端获取到zan平台跳转的url ，进行跳转\n3. 用户在zan平台进行登陆和授权行为\n4. zan平台将浏览器url跳回到hamster平台，地址为：\n5. 前端通过url解析出authCode参数，调用 交换zan的access_token 接口 /api/v2/zan/exchange/access_token\n6. 调用其他zan平台相关需要授权的接口，如创建ApiKEY",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "zan"
                ],
                "summary": "用户是否已经进行了zan授权",
                "responses": {
                    "200": {
                        "description": "desc",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/handler.Result"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "boolean"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/api/v2/zan/node-service/api-keys": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "为⽤户创建API KEY",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "zan"
                ],
                "summary": "创建API KEY接⼝",
                "parameters": [
                    {
                        "description": "创建api key",
                        "name": "apiKeyCreateReq",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/zan.ApiKeyCreateReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/handler.Result"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/zan.ApiKeyBase"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/api/v2/zan/node-service/api-keys/detail": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "调⽤⽅可通过该接⼝查询特定API KEY的详情",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "zan"
                ],
                "summary": "API KEY详情接口",
                "parameters": [
                    {
                        "type": "string",
                        "description": "apiKeyId",
                        "name": "apiKeyId",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/handler.Result"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/zan.ApiKeyDetailInfo"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/api/v2/zan/node-service/api-keys/list": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "调⽤⽅可通过该接⼝分⻚查询⽤户的API KEY列表",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "zan"
                ],
                "summary": "API KEY分⻚查询接⼝",
                "parameters": [
                    {
                        "type": "integer",
                        "default": 1,
                        "description": "页码",
                        "name": "page",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "default": 10,
                        "description": "每⻚数量",
                        "name": "size",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/handler.Result"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/vo.PageResp-zan_ApiKeyDigestInfo"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/api/v2/zan/node-service/api-keys/stats/credit-cost": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "调⽤⽅可通过该接⼝查询当前API KEY在过去24⼩时credit消耗量统计数据",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "zan"
                ],
                "summary": "API KEY credit cost统计查询接⼝",
                "parameters": [
                    {
                        "type": "string",
                        "description": "apiKeyId",
                        "name": "apiKeyId",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/handler.Result"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "array",
                                            "items": {
                                                "$ref": "#/definitions/zan.StatCreditCostItem"
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/api/v2/zan/node-service/api-keys/stats/requests": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "调⽤⽅可通过该接⼝查询当前API KEY过去⼀段时间不同⽣态下的不同⽅法的接⼝调⽤统计数据",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "zan"
                ],
                "summary": "API Key request统计查询接⼝",
                "parameters": [
                    {
                        "type": "string",
                        "description": "apiKeyId",
                        "name": "apiKeyId",
                        "in": "query",
                        "required": true
                    },
                    {
                        "enum": [
                            "STAT_15_MIN",
                            "STAT_1_HOUR",
                            "STAT_24_HOUR",
                            "STAT_7_DAY",
                            "STAT_1_MONTH"
                        ],
                        "type": "string",
                        "description": "时间间隔",
                        "name": "timeInterval",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "⽣态, 此值为链⽣态摘要信息查询接⼝返回的生态编码",
                        "name": "ecosystem",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/handler.Result"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "array",
                                            "items": {
                                                "$ref": "#/definitions/zan.StatMethodCountItem"
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/api/v2/zan/node-service/api-keys/stats/requests-activity": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "调⽤⽅可通过该接⼝查询当前API KEY过去⼀段时间不同⽣态下不同⽅法的接⼝调⽤成功次数统计数据",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "zan"
                ],
                "summary": "API KEY requests activity统计查询接⼝",
                "parameters": [
                    {
                        "type": "string",
                        "description": "apiKeyId",
                        "name": "apiKeyId",
                        "in": "query",
                        "required": true
                    },
                    {
                        "enum": [
                            "STAT_15_MIN",
                            "STAT_1_HOUR",
                            "STAT_24_HOUR",
                            "STAT_7_DAY",
                            "STAT_1_MONTH"
                        ],
                        "type": "string",
                        "description": "时间间隔",
                        "name": "timeInterval",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "⽣态, 此值为链⽣态摘要信息查询接⼝返回的生态编码",
                        "name": "ecosystem",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/handler.Result"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "array",
                                            "items": {
                                                "$ref": "#/definitions/zan.StatMethodRequestActivityDetail"
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/api/v2/zan/node-service/api-keys/stats/requests-activity/failed": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "调用方可通过该接口查询当前API KEY过去一段时间不同生态下某个方法的接口调用失败次数统计数据",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "zan"
                ],
                "summary": "API KEY requests activity failed 统计查询接口",
                "parameters": [
                    {
                        "type": "string",
                        "description": "apiKeyId",
                        "name": "apiKeyId",
                        "in": "query",
                        "required": true
                    },
                    {
                        "enum": [
                            "STAT_15_MIN",
                            "STAT_1_HOUR",
                            "STAT_24_HOUR",
                            "STAT_7_DAY",
                            "STAT_1_MONTH"
                        ],
                        "type": "string",
                        "description": "时间间隔",
                        "name": "timeInterval",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "⽣态, 此值为链⽣态摘要信息查询接⼝返回的生态编码",
                        "name": "ecosystem",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "方法",
                        "name": "method",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/handler.Result"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "array",
                                            "items": {
                                                "$ref": "#/definitions/zan.StatMethodRequestActivityFailedDetailGwInfo"
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/api/v2/zan/node-service/api-keys/stats/requests-origin": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "调⽤⽅可通过该接⼝查询当前API KEY过去⼀段时间不同请求来源的统计数据",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "zan"
                ],
                "summary": "API KEY request Origin统计查询接⼝",
                "parameters": [
                    {
                        "type": "string",
                        "description": "apiKeyId",
                        "name": "apiKeyId",
                        "in": "query",
                        "required": true
                    },
                    {
                        "enum": [
                            "STAT_15_MIN",
                            "STAT_1_HOUR",
                            "STAT_24_HOUR",
                            "STAT_7_DAY",
                            "STAT_1_MONTH"
                        ],
                        "type": "string",
                        "description": "时间间隔",
                        "name": "timeInterval",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/handler.Result"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "array",
                                            "items": {
                                                "$ref": "#/definitions/zan.StatCreditCostOrigin"
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/api/v2/zan/node-service/ecosystems/digest": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "调⽤⽅可通过该接⼝查询特定API KEY的详情。",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "zan"
                ],
                "summary": "链⽣态摘要信息查询接⼝",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/handler.Result"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "array",
                                            "items": {
                                                "$ref": "#/definitions/zan.EcosystemDigestInfo"
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/api/v2/zan/node-service/plan": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "调⽤⽅可通过该接⼝查询当前⽤户的节点服务套餐情况",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "zan"
                ],
                "summary": "套餐信息查询接⼝",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/handler.Result"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "array",
                                            "items": {
                                                "$ref": "#/definitions/zan.EcosystemDigestInfo"
                                            }
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
        "handler.Result": {
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
        },
        "vo.ExchangeAccessTokenVo": {
            "type": "object",
            "properties": {
                "authCode": {
                    "description": "authCode 交换accessToken的授权码",
                    "type": "string"
                }
            }
        },
        "vo.PageResp-zan_ApiKeyDigestInfo": {
            "type": "object",
            "properties": {
                "data": {
                    "description": "条⽬列表",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/zan.ApiKeyDigestInfo"
                    }
                },
                "page": {
                    "description": "当前页",
                    "type": "integer"
                },
                "pageCount": {
                    "description": "总页数",
                    "type": "integer"
                },
                "pageSize": {
                    "description": "每页条目数",
                    "type": "integer"
                },
                "total": {
                    "description": "总条⽬数",
                    "type": "integer"
                }
            }
        },
        "zan.ApiKeyBase": {
            "type": "object",
            "properties": {
                "apiKeyId": {
                    "description": "API KEY ID",
                    "type": "string"
                }
            }
        },
        "zan.ApiKeyCreateReq": {
            "type": "object",
            "properties": {
                "name": {
                    "description": "API Key 名字",
                    "type": "string"
                }
            }
        },
        "zan.ApiKeyDetailInfo": {
            "type": "object",
            "properties": {
                "apiKeyId": {
                    "description": "API KEY ID",
                    "type": "string"
                },
                "createdTime": {
                    "description": "创建时间",
                    "type": "string"
                },
                "ecosystemDetailInfos": {
                    "description": "API KEY的⽣态访问信息",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/zan.ApiKeyEcosystemDetail"
                    }
                }
            }
        },
        "zan.ApiKeyDigestInfo": {
            "type": "object",
            "properties": {
                "apiKeyId": {
                    "description": "API KEY ID",
                    "type": "string"
                },
                "createdTime": {
                    "description": "创建时间",
                    "type": "string"
                },
                "name": {
                    "description": "API Key 名字",
                    "type": "string"
                }
            }
        },
        "zan.ApiKeyEcosystemDetail": {
            "type": "object",
            "properties": {
                "ecosystemCode": {
                    "description": "⽣态编码",
                    "type": "string"
                },
                "ecosystemIcon": {
                    "description": "⽣态icon地址",
                    "type": "string"
                },
                "ecosystemName": {
                    "description": "⽣态名",
                    "type": "string"
                },
                "networkDetailInfoList": {
                    "description": "链⽹络信息",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/zan.NetworkDetailInfo"
                    }
                }
            }
        },
        "zan.EcosystemDigestInfo": {
            "type": "object",
            "properties": {
                "ecosystemCode": {
                    "description": "⽣态编码",
                    "type": "string"
                },
                "ecosystemIcon": {
                    "description": "⽣态icon地址",
                    "type": "string"
                },
                "ecosystemName": {
                    "description": "⽣态名",
                    "type": "string"
                },
                "networks": {
                    "description": "链⽹络摘要信息",
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                }
            }
        },
        "zan.NetworkDetailInfo": {
            "type": "object",
            "properties": {
                "code": {
                    "description": "⽹络编码",
                    "type": "string"
                },
                "httpsUrl": {
                    "description": "⽹络https请求地址",
                    "type": "string"
                },
                "name": {
                    "description": "⽹络名称",
                    "type": "string"
                },
                "wssUrl": {
                    "description": "⽹络wss请求地址",
                    "type": "string"
                }
            }
        },
        "zan.StatCreditCostItem": {
            "type": "object",
            "properties": {
                "dataTime": {
                    "description": "数据时间时间戳",
                    "type": "integer"
                },
                "totalCredit": {
                    "description": "credit消耗量",
                    "type": "integer"
                }
            }
        },
        "zan.StatCreditCostOrigin": {
            "type": "object",
            "properties": {
                "dataTime": {
                    "description": "数据时间时间戳",
                    "type": "integer"
                },
                "httpsNum": {
                    "description": "http请求数量",
                    "type": "integer"
                },
                "originIp": {
                    "description": "请求来源Ip",
                    "type": "string"
                },
                "totalNum": {
                    "description": "调⽤总数",
                    "type": "integer"
                },
                "wssNum": {
                    "description": "wss请求数量",
                    "type": "integer"
                }
            }
        },
        "zan.StatMethodCountItem": {
            "type": "object",
            "properties": {
                "dataTime": {
                    "description": "数据时间时间戳",
                    "type": "integer"
                },
                "method": {
                    "description": "rpc 请求的⽅法",
                    "type": "string"
                },
                "num": {
                    "description": "调用次数",
                    "type": "integer"
                }
            }
        },
        "zan.StatMethodRequestActivityDetail": {
            "type": "object",
            "properties": {
                "failedNum": {
                    "description": "失败次数",
                    "type": "integer"
                },
                "method": {
                    "description": "rpc 请求的⽅法",
                    "type": "string"
                },
                "totalNum": {
                    "description": "成功次数",
                    "type": "integer"
                }
            }
        },
        "zan.StatMethodRequestActivityFailedDetailGwInfo": {
            "type": "object",
            "properties": {
                "httpsNum": {
                    "description": "http请求数",
                    "type": "integer"
                },
                "status": {
                    "description": "状态错误码",
                    "type": "string"
                },
                "wssNum": {
                    "description": "wss 请求数",
                    "type": "integer"
                }
            }
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "type": "apiKey",
            "name": "Access-Token",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "0.0.1",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "hamster paas API 接口文档",
	Description:      "提供zan相关接口",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
