// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/kashifkhan0771/ProductManagement/gen/models"
)

// GetProductReader is a Reader for the GetProduct structure.
type GetProductReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetProductReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetProductOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetProductNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetProductInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetProductOK creates a GetProductOK with default headers values
func NewGetProductOK() *GetProductOK {
	return &GetProductOK{}
}

/* GetProductOK describes a response with status code 200, with default header values.

product response
*/
type GetProductOK struct {
	Payload *models.Product
}

func (o *GetProductOK) Error() string {
	return fmt.Sprintf("[GET /product/{ID}][%d] getProductOK  %+v", 200, o.Payload)
}
func (o *GetProductOK) GetPayload() *models.Product {
	return o.Payload
}

func (o *GetProductOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Product)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetProductNotFound creates a GetProductNotFound with default headers values
func NewGetProductNotFound() *GetProductNotFound {
	return &GetProductNotFound{}
}

/* GetProductNotFound describes a response with status code 404, with default header values.

product not found
*/
type GetProductNotFound struct {
}

func (o *GetProductNotFound) Error() string {
	return fmt.Sprintf("[GET /product/{ID}][%d] getProductNotFound ", 404)
}

func (o *GetProductNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetProductInternalServerError creates a GetProductInternalServerError with default headers values
func NewGetProductInternalServerError() *GetProductInternalServerError {
	return &GetProductInternalServerError{}
}

/* GetProductInternalServerError describes a response with status code 500, with default header values.

internal server error
*/
type GetProductInternalServerError struct {
}

func (o *GetProductInternalServerError) Error() string {
	return fmt.Sprintf("[GET /product/{ID}][%d] getProductInternalServerError ", 500)
}

func (o *GetProductInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
