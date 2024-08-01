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
        "/Command": {
            "post": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Create a new device",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "device"
                ],
                "summary": "Create a new device",
                "parameters": [
                    {
                        "description": "Device request body",
                        "name": "device",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/deviceproto.CommandRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/deviceproto.CommandRequest"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/device": {
            "put": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Update a device",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "device"
                ],
                "summary": "Update a device",
                "parameters": [
                    {
                        "description": "Device update request body",
                        "name": "device",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/deviceproto.Device"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/deviceproto.DeviceOperationResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
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
                "description": "Create a new device",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "device"
                ],
                "summary": "Create a new device",
                "parameters": [
                    {
                        "description": "Device request body",
                        "name": "device",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/deviceproto.DeviceRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/deviceproto.DeviceResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Delete a device",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "device"
                ],
                "summary": "Delete a device",
                "parameters": [
                    {
                        "description": "Delete device request body",
                        "name": "device",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/deviceproto.DeviceResponse"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/deviceproto.DeviceOperationResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/device/{id}": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Get a device by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "device"
                ],
                "summary": "Get a device by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Device ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/deviceproto.Device"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/login": {
            "post": {
                "description": "Login a user and get a JWT token",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Login a user",
                "parameters": [
                    {
                        "description": "Login request body",
                        "name": "login",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/userproto.LoginRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/userproto.UserResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
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
                        "Bearer": []
                    }
                ],
                "description": "Get all users",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Get all users",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/userproto.ListUser"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
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
                "description": "Update a user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Update a user",
                "parameters": [
                    {
                        "description": "User request body",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/userproto.Request"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/userproto.UpdateResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
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
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Create a new user",
                "parameters": [
                    {
                        "description": "User request body",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/userproto.UserRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/userproto.UserResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Delete a user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Delete a user",
                "parameters": [
                    {
                        "description": "Delete user request body",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/userproto.DeleteUserRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/userproto.UpdateResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/users/email": {
            "put": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Update a user's email",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Update a user's email",
                "parameters": [
                    {
                        "description": "Email update request body",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/userproto.UpdateEmailRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/userproto.UpdateResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/users/password": {
            "put": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Update a user's password",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Update a user's password",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "Password update request body",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/userproto.UpdateUserRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/userproto.UpdateResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/users/{id}": {
            "get": {
                "security": [
                    {
                        "Bearer Auth": []
                    }
                ],
                "description": "Get user by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Get user by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "User ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/userproto.User"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/verifycode": {
            "post": {
                "description": "Login a user and get a JWT token",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "VerifyCode a user",
                "parameters": [
                    {
                        "description": "Login request body",
                        "name": "login",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/userproto.LoginRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/userproto.Res"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "deviceproto.CommandRequest": {
            "type": "object",
            "properties": {
                "command_payload": {
                    "$ref": "#/definitions/deviceproto.Command_Payload"
                },
                "device_id": {
                    "type": "string"
                },
                "status": {
                    "type": "string"
                },
                "timeStamp": {
                    "type": "string"
                },
                "type": {
                    "type": "string"
                },
                "user_id": {
                    "type": "string"
                }
            }
        },
        "deviceproto.Command_Payload": {
            "type": "object",
            "properties": {
                "brightness": {
                    "type": "integer"
                }
            }
        },
        "deviceproto.Configuration_Settings": {
            "type": "object",
            "properties": {
                "brightness": {
                    "type": "integer"
                },
                "color": {
                    "type": "string"
                }
            }
        },
        "deviceproto.Device": {
            "type": "object",
            "properties": {
                "configuration_settings": {
                    "$ref": "#/definitions/deviceproto.Configuration_Settings"
                },
                "connectivity_status": {
                    "type": "string"
                },
                "device_id": {
                    "type": "string"
                },
                "device_name": {
                    "type": "string"
                },
                "device_status": {
                    "type": "string"
                },
                "device_type": {
                    "type": "string"
                },
                "firmware_version": {
                    "type": "string"
                },
                "last_updated": {
                    "type": "string"
                },
                "location": {
                    "type": "string"
                },
                "user_id": {
                    "type": "string"
                }
            }
        },
        "deviceproto.DeviceOperationResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                }
            }
        },
        "deviceproto.DeviceRequest": {
            "type": "object",
            "properties": {
                "configuration_settings": {
                    "$ref": "#/definitions/deviceproto.Configuration_Settings"
                },
                "connectivity_status": {
                    "type": "string"
                },
                "device_name": {
                    "type": "string"
                },
                "device_status": {
                    "type": "string"
                },
                "device_type": {
                    "type": "string"
                },
                "firmware_version": {
                    "type": "string"
                },
                "last_updated": {
                    "type": "string"
                },
                "location": {
                    "type": "string"
                },
                "user_id": {
                    "type": "string"
                }
            }
        },
        "deviceproto.DeviceResponse": {
            "type": "object",
            "properties": {
                "device_id": {
                    "type": "string"
                }
            }
        },
        "userproto.DeleteUserRequest": {
            "type": "object",
            "properties": {
                "user_id": {
                    "type": "string"
                }
            }
        },
        "userproto.ListUser": {
            "type": "object",
            "properties": {
                "user": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/userproto.User"
                    }
                }
            }
        },
        "userproto.LoginRequest": {
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
        "userproto.Profile": {
            "type": "object",
            "properties": {
                "address": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "userproto.Request": {
            "type": "object",
            "properties": {
                "profile": {
                    "$ref": "#/definitions/userproto.Profile"
                },
                "user_id": {
                    "type": "string"
                },
                "user_name": {
                    "type": "string"
                }
            }
        },
        "userproto.Res": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                }
            }
        },
        "userproto.UpdateEmailRequest": {
            "type": "object",
            "properties": {
                "new_email": {
                    "type": "string"
                },
                "old_email": {
                    "type": "string"
                },
                "user_id": {
                    "type": "string"
                }
            }
        },
        "userproto.UpdateResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                }
            }
        },
        "userproto.UpdateUserRequest": {
            "type": "object",
            "properties": {
                "new_password": {
                    "type": "string"
                },
                "old_password": {
                    "type": "string"
                },
                "user_id": {
                    "type": "string"
                }
            }
        },
        "userproto.User": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "profile": {
                    "$ref": "#/definitions/userproto.Profile"
                },
                "user_id": {
                    "type": "string"
                },
                "user_name": {
                    "type": "string"
                }
            }
        },
        "userproto.UserRequest": {
            "type": "object",
            "properties": {
                "confirm_password": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "profile": {
                    "$ref": "#/definitions/userproto.Profile"
                },
                "user_name": {
                    "type": "string"
                }
            }
        },
        "userproto.UserResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                },
                "token": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "Bearer": {
            "description": "Enter the token in the format ` + "`" + `Bearer {token}` + "`" + `",
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
	Title:            "Artisan Connect",
	Description:      "This is a sample server for a restaurant reservation system.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	// LeftDelim:        "{{",
	// RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
