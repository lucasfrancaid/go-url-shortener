// Code generated by swaggo/swag. DO NOT EDIT.

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
            "name": "API Support",
            "url": "http://www.swagger.io/support",
            "email": "support@swagger.io"
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
        "/shorten": {
            "post": {
                "description": "Shorten an URL",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Shortener"
                ],
                "summary": "Shorten",
                "parameters": [
                    {
                        "description": "URL to Shorten",
                        "name": "payload",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.ShortenDTO"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.ShortenedDTO"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/presenter.ErrorResponseHTTP"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/presenter.ErrorResponseHTTP"
                        }
                    }
                }
            }
        },
        "/stats/{hashedURL}": {
            "get": {
                "description": "Show statistics about a shortened URL",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Shortener"
                ],
                "summary": "Stats",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Last block of Shortened URL, the value after /u/ part",
                        "name": "hashedURL",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.ShortenerStatsDTO"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/presenter.ErrorResponseHTTP"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/presenter.ErrorResponseHTTP"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/presenter.ErrorResponseHTTP"
                        }
                    }
                }
            }
        },
        "/u/{hashedURL}": {
            "get": {
                "description": "Redirect a Shortened URL to Original URL",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Shortener"
                ],
                "summary": "Redirect",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Last block of Shortened URL, the value after /u/ part",
                        "name": "hashedURL",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "303": {
                        "description": "See Other"
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/presenter.ErrorResponseHTTP"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/presenter.ErrorResponseHTTP"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/presenter.ErrorResponseHTTP"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "dto.ShortenDTO": {
            "type": "object",
            "properties": {
                "url": {
                    "type": "string"
                }
            }
        },
        "dto.ShortenedDTO": {
            "type": "object",
            "properties": {
                "shortenedURL": {
                    "type": "string"
                }
            }
        },
        "dto.ShortenerStatsDTO": {
            "type": "object",
            "properties": {
                "counter": {
                    "type": "integer"
                }
            }
        },
        "presenter.ErrorResponseHTTP": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:3000",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "URL Shortener API",
	Description:      "Go URL Shortener implemented using Clean Architecture with Echo and Fiber as HTTP Adapters.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
