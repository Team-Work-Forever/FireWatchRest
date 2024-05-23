// Package fireWatch Code generated by swaggo/swag. DO NOT EDIT
package fireWatch

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
            "email": "fiber@swagger.io"
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
        "/auth/forgot_password": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Request a Password Reset",
                "parameters": [
                    {
                        "type": "string",
                        "description": "some description",
                        "name": "accept-language",
                        "in": "header"
                    },
                    {
                        "type": "string",
                        "description": "Email address associated with the account",
                        "name": "email",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/auth/login": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Authenticate with account",
                "parameters": [
                    {
                        "type": "string",
                        "description": "some description",
                        "name": "accept-language",
                        "in": "header"
                    },
                    {
                        "description": "Login Payload",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/contracts.LoginRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/contracts.AuthResponse"
                        }
                    }
                }
            }
        },
        "/auth/refresh_tokens": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Refresh Tokens",
                "parameters": [
                    {
                        "type": "string",
                        "description": "some description",
                        "name": "accept-language",
                        "in": "header"
                    },
                    {
                        "type": "string",
                        "description": "JWT token to be refreshed",
                        "name": "token",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/contracts.AuthResponse"
                        }
                    }
                }
            }
        },
        "/auth/reset_password": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Reset Password",
                "parameters": [
                    {
                        "type": "string",
                        "description": "some description",
                        "name": "accept-language",
                        "in": "header"
                    },
                    {
                        "type": "string",
                        "description": "A unique token sent to the user's email for password reset",
                        "name": "forgot_token",
                        "in": "query",
                        "required": true
                    },
                    {
                        "description": "Reset Password Payload",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/contracts.ResetPasswordRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Password reset successfully",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/auth/signUp": {
            "post": {
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Create an Account",
                "parameters": [
                    {
                        "type": "string",
                        "description": "some description",
                        "name": "accept-language",
                        "in": "header"
                    },
                    {
                        "type": "string",
                        "name": "city",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "name": "email",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "name": "first_name",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "name": "last_name",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "name": "nif",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "name": "password",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "name": "phone_code",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "name": "phone_number",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "name": "street",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "name": "street_port",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "name": "user_name",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "name": "zip_code",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "file",
                        "description": "User avatar",
                        "name": "avatar",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/contracts.AuthResponse"
                        }
                    }
                }
            }
        },
        "/burns": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Burn"
                ],
                "summary": "Fetch burns",
                "parameters": [
                    {
                        "type": "string",
                        "description": "some description",
                        "name": "accept-language",
                        "in": "header"
                    },
                    {
                        "type": "string",
                        "description": "search burn title",
                        "name": "search",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "search by burn state",
                        "name": "state",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "search by an inital date",
                        "name": "start_date",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "search by an end date",
                        "name": "end_date",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "default": 1,
                        "description": "view page",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "default": 10,
                        "description": "number of returned elements",
                        "name": "page_size",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/geojson.GeoJsonFeature"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Burn"
                ],
                "summary": "Create an Burn Request",
                "parameters": [
                    {
                        "type": "string",
                        "description": "some description",
                        "name": "accept-language",
                        "in": "header"
                    },
                    {
                        "type": "boolean",
                        "name": "has_backup_team",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "name": "init_date",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "name": "initial_propose",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "number",
                        "name": "lat",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "number",
                        "name": "lon",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "name": "reason",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "name": "title",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "name": "type",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/contracts.BurnActionResponse"
                        }
                    }
                }
            }
        },
        "/burns/reasons": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Burn"
                ],
                "summary": "Burn Available Reasons",
                "parameters": [
                    {
                        "type": "string",
                        "description": "some description",
                        "name": "accept-language",
                        "in": "header"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/burns/states": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Burn"
                ],
                "summary": "Burn Available States",
                "parameters": [
                    {
                        "type": "string",
                        "description": "some description",
                        "name": "accept-language",
                        "in": "header"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/burns/types": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Burn"
                ],
                "summary": "Burn Available Types",
                "parameters": [
                    {
                        "type": "string",
                        "description": "some description",
                        "name": "accept-language",
                        "in": "header"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/burns/{id}": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Burn"
                ],
                "summary": "Fetch Burn By Id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "some description",
                        "name": "accept-language",
                        "in": "header"
                    },
                    {
                        "type": "string",
                        "description": "Fetch the burn by id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/geojson.GeoJsonFeature"
                        }
                    }
                }
            },
            "put": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Burn"
                ],
                "summary": "Create an Burn Request",
                "parameters": [
                    {
                        "type": "string",
                        "description": "some description",
                        "name": "accept-language",
                        "in": "header"
                    },
                    {
                        "type": "string",
                        "description": "Fetch the burn by id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "boolean",
                        "name": "has_backup_team",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "name": "init_date",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "name": "initial_propose",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "number",
                        "name": "lat",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "number",
                        "name": "lon",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "name": "reason",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "name": "title",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "name": "type",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "202": {
                        "description": "Accepted",
                        "schema": {
                            "$ref": "#/definitions/contracts.BurnActionResponse"
                        }
                    }
                }
            }
        },
        "/profile": {
            "put": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Profile"
                ],
                "summary": "Update Profile Information",
                "parameters": [
                    {
                        "type": "string",
                        "description": "some description",
                        "name": "accept-language",
                        "in": "header"
                    },
                    {
                        "type": "string",
                        "name": "city",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "name": "email",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "name": "phone_code",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "name": "phone_number",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "name": "street",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "name": "street_port",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "name": "user_name",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "name": "zip_code",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "file",
                        "description": "User avatar",
                        "name": "avatar",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "202": {
                        "description": "Accepted",
                        "schema": {
                            "$ref": "#/definitions/contracts.ProfileResponse"
                        }
                    }
                }
            }
        },
        "/whoami": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Profile"
                ],
                "summary": "Fetch Profile Information",
                "parameters": [
                    {
                        "type": "string",
                        "description": "some description",
                        "name": "accept-language",
                        "in": "header"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/contracts.ProfileResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "contracts.AddressResponse": {
            "type": "object",
            "properties": {
                "city": {
                    "type": "string"
                },
                "number": {
                    "type": "integer"
                },
                "street": {
                    "type": "string"
                },
                "zip_code": {
                    "$ref": "#/definitions/contracts.ZipCodeResponse"
                }
            }
        },
        "contracts.AuthResponse": {
            "type": "object",
            "properties": {
                "access_token": {
                    "type": "string"
                },
                "refresh_token": {
                    "type": "string"
                }
            }
        },
        "contracts.BurnActionResponse": {
            "type": "object",
            "properties": {
                "burn_id": {
                    "type": "string"
                },
                "state": {
                    "type": "string"
                }
            }
        },
        "contracts.LoginRequest": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "contracts.PhoneResponse": {
            "type": "object",
            "properties": {
                "country_code": {
                    "type": "string"
                },
                "number": {
                    "type": "string"
                }
            }
        },
        "contracts.ProfileResponse": {
            "type": "object",
            "properties": {
                "address": {
                    "$ref": "#/definitions/contracts.AddressResponse"
                },
                "email": {
                    "type": "string"
                },
                "first_name": {
                    "type": "string"
                },
                "last_name": {
                    "type": "string"
                },
                "phone": {
                    "$ref": "#/definitions/contracts.PhoneResponse"
                },
                "user_name": {
                    "type": "string"
                },
                "user_type": {
                    "type": "string"
                }
            }
        },
        "contracts.ResetPasswordRequest": {
            "type": "object",
            "required": [
                "confirm_password",
                "password"
            ],
            "properties": {
                "confirm_password": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "contracts.ZipCodeResponse": {
            "type": "object",
            "properties": {
                "value": {
                    "type": "string"
                }
            }
        },
        "geojson.GeoJsonFeature": {
            "type": "object",
            "properties": {
                "geometry": {
                    "$ref": "#/definitions/geojson.Geometry"
                },
                "properties": {},
                "type": {
                    "type": "string"
                }
            }
        },
        "geojson.Geometry": {
            "type": "object",
            "properties": {
                "coordinates": {
                    "type": "array",
                    "items": {
                        "type": "number"
                    }
                },
                "type": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "Bearer": {
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
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "FireWatch API",
	Description:      "This is the api for Fire Watch Mobile Application",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
