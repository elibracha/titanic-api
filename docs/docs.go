// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "Eli Bracha"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/health": {
            "get": {
                "description": "Get health check status",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "health"
                ],
                "summary": "Get health check status",
                "operationId": "healthcheck",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/healthcheck.Status"
                        }
                    }
                }
            }
        },
        "/passenger": {
            "get": {
                "description": "Get all passengers",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "passenger"
                ],
                "summary": "Get passengers",
                "operationId": "passenger-get-all",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/passenger.Response"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.Error"
                        }
                    }
                }
            }
        },
        "/passenger/fare/histogram/percentile": {
            "get": {
                "description": "Get histogram represention of number of passengers in each precentile",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "passenger"
                ],
                "summary": "Get fare histogram histogram",
                "operationId": "passenger-fare-histogram",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/histogram.Histogram"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.Error"
                        }
                    }
                }
            }
        },
        "/passenger/{id}": {
            "get": {
                "description": "Get passenger by ID number",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "passenger"
                ],
                "summary": "Get passenger",
                "operationId": "passenger-get",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Passenger ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "array",
                        "items": {
                            "type": "string"
                        },
                        "collectionFormat": "csv",
                        "description": "Allowed: id, age, sex, name, survived, class, siblings-spouses, parents-children, ticket, fare, cabin, embarked",
                        "name": "attributes",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/passenger.Response"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/response.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.Error"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "healthcheck.Status": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "state": {
                    "type": "string"
                }
            }
        },
        "histogram.Entry": {
            "type": "object",
            "properties": {
                "bin": {
                    "type": "integer"
                },
                "count": {
                    "type": "integer"
                }
            }
        },
        "histogram.Histogram": {
            "type": "object",
            "properties": {
                "entries": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/histogram.Entry"
                    }
                }
            }
        },
        "passenger.Response": {
            "type": "object",
            "properties": {
                "age": {
                    "type": "string"
                },
                "cabin": {
                    "type": "string"
                },
                "class": {
                    "type": "integer"
                },
                "embarked": {
                    "type": "string"
                },
                "fare": {
                    "type": "number"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "parents-children": {
                    "type": "integer"
                },
                "sex": {
                    "type": "string"
                },
                "siblings-spouses": {
                    "type": "integer"
                },
                "survived": {
                    "type": "integer"
                },
                "ticket": {
                    "type": "string"
                }
            }
        },
        "response.Error": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "message": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "Titanic API",
	Description:      "This is API provide multiple functionality endpoints over titanic dataset",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
