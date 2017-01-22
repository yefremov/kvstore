package restapi

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import "encoding/json"

// SwaggerJSON embedded version of the swagger document used at generation time
var SwaggerJSON json.RawMessage

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "K/V store is a simple single node store for retrieving key/value information",
    "title": "K/V store",
    "contact": {
      "name": "Ivan Porto Carrero",
      "email": "ivan@flanders.co.nz"
    },
    "license": {
      "name": "MIT",
      "url": "https://github.com/go-openapi/kvstore/blob/master/LICENSE"
    },
    "version": "0.0.1"
  },
  "paths": {
    "/kv": {
      "get": {
        "description": "lists all the keys",
        "tags": [
          "kv"
        ],
        "operationId": "findKeys",
        "parameters": [
          {
            "type": "string",
            "name": "prefix",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "list the keys known to this datastore",
            "schema": {
              "type": "array",
              "items": {
                "type": "string"
              }
            },
            "headers": {
              "X-Request-Id": {
                "type": "string",
                "description": "The request id this is a response to"
              }
            }
          },
          "default": {
            "$ref": "#/responses/errorResponse"
          }
        }
      },
      "parameters": [
        {
          "$ref": "#/parameters/requestId"
        }
      ]
    },
    "/kv/{key}": {
      "get": {
        "produces": [
          "application/octet-stream"
        ],
        "tags": [
          "kv"
        ],
        "operationId": "getEntry",
        "parameters": [
          {
            "type": "string",
            "name": "If-None-Match",
            "in": "header"
          }
        ],
        "responses": {
          "200": {
            "description": "entry was found",
            "schema": {
              "type": "string",
              "format": "binary"
            },
            "headers": {
              "ETag": {
                "type": "string",
                "description": "The version of this entry"
              },
              "Last-Modified": {
                "type": "string",
                "description": "The time this entry was last modified"
              },
              "X-Request-Id": {
                "type": "string",
                "description": "The request id this is a response to"
              }
            }
          },
          "304": {
            "description": "entry was found but not modified",
            "headers": {
              "ETag": {
                "type": "string",
                "description": "The version of this entry"
              },
              "Last-Modified": {
                "type": "string",
                "description": "The time this entry was last modified"
              },
              "X-Request-Id": {
                "type": "string",
                "description": "The request id this is a response to"
              }
            }
          },
          "404": {
            "$ref": "#/responses/errorNotFound"
          },
          "default": {
            "$ref": "#/responses/errorResponse"
          }
        }
      },
      "put": {
        "consumes": [
          "application/octet-stream"
        ],
        "tags": [
          "kv"
        ],
        "operationId": "putEntry",
        "parameters": [
          {
            "pattern": "[0-9]*",
            "type": "string",
            "description": "when this is an update to an entry, then this field needs to be present",
            "name": "If-Match",
            "in": "header"
          },
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "type": "string",
              "format": "binary",
              "maxLength": 536870912
            }
          }
        ],
        "responses": {
          "201": {
            "description": "entry was created",
            "headers": {
              "Etag": {
                "type": "string",
                "description": "The version of this entry"
              },
              "Location": {
                "type": "string",
                "format": "uri",
                "description": "the location to get the newly created entry"
              },
              "X-Request-Id": {
                "type": "string",
                "description": "The request id this is a response to"
              }
            }
          },
          "204": {
            "description": "entry was updated",
            "headers": {
              "ETag": {
                "type": "string",
                "description": "The version of this entry"
              },
              "X-Request-Id": {
                "type": "string",
                "description": "The request id this is a response to"
              }
            }
          },
          "404": {
            "$ref": "#/responses/errorNotFound"
          },
          "409": {
            "description": "there is a version mismatch for the entry",
            "schema": {
              "$ref": "#/definitions/error"
            },
            "headers": {
              "X-Request-Id": {
                "type": "string",
                "description": "The request id this is a response to"
              }
            }
          },
          "410": {
            "description": "The entry is deleted",
            "schema": {
              "$ref": "#/definitions/error"
            },
            "headers": {
              "X-Request-Id": {
                "type": "string",
                "description": "The request id this is a response to"
              }
            }
          },
          "default": {
            "$ref": "#/responses/errorResponse"
          }
        }
      },
      "delete": {
        "tags": [
          "kv"
        ],
        "operationId": "deleteEntry",
        "responses": {
          "204": {
            "description": "the delete was successful",
            "headers": {
              "X-Request-Id": {
                "type": "string",
                "description": "The request id this is a response to"
              }
            }
          },
          "default": {
            "$ref": "#/responses/errorResponse"
          }
        }
      },
      "parameters": [
        {
          "$ref": "#/parameters/requestId"
        },
        {
          "$ref": "#/parameters/entryKey"
        }
      ]
    }
  },
  "definitions": {
    "error": {
      "description": "the error model is a model for all the error responses coming from kvstore\n",
      "type": "object",
      "required": [
        "message",
        "code"
      ],
      "properties": {
        "cause": {
          "$ref": "#/definitions/error"
        },
        "code": {
          "description": "The error code",
          "type": "integer",
          "format": "int64"
        },
        "helpUrl": {
          "description": "link to help page explaining the error in more detail",
          "type": "string",
          "format": "uri"
        },
        "message": {
          "description": "The error message",
          "type": "string"
        }
      }
    }
  },
  "parameters": {
    "entryKey": {
      "minLength": 1,
      "type": "string",
      "description": "The key for a given entry",
      "name": "key",
      "in": "path",
      "required": true
    },
    "requestId": {
      "minLength": 1,
      "type": "string",
      "description": "A unique UUID for the request",
      "name": "X-Request-Id",
      "in": "header"
    }
  },
  "responses": {
    "errorNotFound": {
      "description": "The entry was not found",
      "schema": {
        "$ref": "#/definitions/error"
      },
      "headers": {
        "X-Request-Id": {
          "type": "string",
          "description": "The request id this is a response to"
        }
      }
    },
    "errorResponse": {
      "description": "Error",
      "schema": {
        "$ref": "#/definitions/error"
      },
      "headers": {
        "X-Request-Id": {
          "type": "string",
          "description": "The request id this is a response to"
        }
      }
    }
  }
}`))
}
