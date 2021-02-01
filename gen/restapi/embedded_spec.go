// Code generated by go-swagger; DO NOT EDIT.

package restapi

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

var (
	// SwaggerJSON embedded version of the swagger document used at generation time
	SwaggerJSON json.RawMessage
	// FlatSwaggerJSON embedded flattened version of the swagger document used at generation time
	FlatSwaggerJSON json.RawMessage
)

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
    "description": "A Product Management System",
    "title": "Product Management",
    "version": "1.0"
  },
  "host": "localhost:8080",
  "basePath": "/",
  "paths": {
    "/Product/{ID}": {
      "get": {
        "description": "return product based on UUID",
        "operationId": "getProduct",
        "parameters": [
          {
            "type": "string",
            "description": "UUID of the product",
            "name": "ID",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "product response",
            "schema": {
              "$ref": "#/definitions/Product"
            }
          },
          "404": {
            "description": "product not found"
          },
          "500": {
            "description": "internal server error"
          }
        }
      },
      "put": {
        "operationId": "editProduct",
        "parameters": [
          {
            "type": "string",
            "description": "UUID of the product",
            "name": "ID",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "product updated",
            "schema": {
              "$ref": "#/definitions/Product"
            }
          },
          "400": {
            "description": "bad request"
          },
          "500": {
            "description": "internal server error"
          }
        }
      },
      "delete": {
        "operationId": "deleteProduct",
        "parameters": [
          {
            "type": "string",
            "description": "UUID of the product",
            "name": "ID",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "204": {
            "description": "product deleted"
          },
          "400": {
            "description": "bad request"
          },
          "500": {
            "description": "internal server error"
          }
        }
      }
    },
    "/addProduct": {
      "post": {
        "operationId": "addProduct",
        "parameters": [
          {
            "description": "product's detail",
            "name": "product",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/Product"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "Product added",
            "schema": {
              "$ref": "#/definitions/Product"
            }
          },
          "400": {
            "description": "bad request"
          },
          "409": {
            "description": "product already exist"
          },
          "500": {
            "description": "internal server error"
          }
        }
      }
    }
  },
  "definitions": {
    "Product": {
      "type": "object",
      "properties": {
        "ID": {
          "type": "string"
        },
        "Name": {
          "type": "string"
        },
        "Price": {
          "type": "integer"
        },
        "ProductID": {
          "type": "string"
        },
        "Quantity": {
          "type": "integer"
        },
        "Status": {
          "type": "integer"
        }
      }
    }
  }
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
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
    "description": "A Product Management System",
    "title": "Product Management",
    "version": "1.0"
  },
  "host": "localhost:8080",
  "basePath": "/",
  "paths": {
    "/Product/{ID}": {
      "get": {
        "description": "return product based on UUID",
        "operationId": "getProduct",
        "parameters": [
          {
            "type": "string",
            "description": "UUID of the product",
            "name": "ID",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "product response",
            "schema": {
              "$ref": "#/definitions/Product"
            }
          },
          "404": {
            "description": "product not found"
          },
          "500": {
            "description": "internal server error"
          }
        }
      },
      "put": {
        "operationId": "editProduct",
        "parameters": [
          {
            "type": "string",
            "description": "UUID of the product",
            "name": "ID",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "product updated",
            "schema": {
              "$ref": "#/definitions/Product"
            }
          },
          "400": {
            "description": "bad request"
          },
          "500": {
            "description": "internal server error"
          }
        }
      },
      "delete": {
        "operationId": "deleteProduct",
        "parameters": [
          {
            "type": "string",
            "description": "UUID of the product",
            "name": "ID",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "204": {
            "description": "product deleted"
          },
          "400": {
            "description": "bad request"
          },
          "500": {
            "description": "internal server error"
          }
        }
      }
    },
    "/addProduct": {
      "post": {
        "operationId": "addProduct",
        "parameters": [
          {
            "description": "product's detail",
            "name": "product",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/Product"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "Product added",
            "schema": {
              "$ref": "#/definitions/Product"
            }
          },
          "400": {
            "description": "bad request"
          },
          "409": {
            "description": "product already exist"
          },
          "500": {
            "description": "internal server error"
          }
        }
      }
    }
  },
  "definitions": {
    "Product": {
      "type": "object",
      "properties": {
        "ID": {
          "type": "string"
        },
        "Name": {
          "type": "string"
        },
        "Price": {
          "type": "integer"
        },
        "ProductID": {
          "type": "string"
        },
        "Quantity": {
          "type": "integer"
        },
        "Status": {
          "type": "integer"
        }
      }
    }
  }
}`))
}
