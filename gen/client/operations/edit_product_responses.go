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

// EditProductReader is a Reader for the EditProduct structure.
type EditProductReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EditProductReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewEditProductOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewEditProductBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewEditProductInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewEditProductOK creates a EditProductOK with default headers values
func NewEditProductOK() *EditProductOK {
	return &EditProductOK{}
}

/* EditProductOK describes a response with status code 200, with default header values.

product updated
*/
type EditProductOK struct {
	Payload *models.Product
}

func (o *EditProductOK) Error() string {
	return fmt.Sprintf("[PUT /product/{ID}][%d] editProductOK  %+v", 200, o.Payload)
}
func (o *EditProductOK) GetPayload() *models.Product {
	return o.Payload
}

func (o *EditProductOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Product)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEditProductBadRequest creates a EditProductBadRequest with default headers values
func NewEditProductBadRequest() *EditProductBadRequest {
	return &EditProductBadRequest{}
}

/* EditProductBadRequest describes a response with status code 400, with default header values.

bad request
*/
type EditProductBadRequest struct {
}

func (o *EditProductBadRequest) Error() string {
	return fmt.Sprintf("[PUT /product/{ID}][%d] editProductBadRequest ", 400)
}

func (o *EditProductBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewEditProductInternalServerError creates a EditProductInternalServerError with default headers values
func NewEditProductInternalServerError() *EditProductInternalServerError {
	return &EditProductInternalServerError{}
}

/* EditProductInternalServerError describes a response with status code 500, with default header values.

internal server error
*/
type EditProductInternalServerError struct {
}

func (o *EditProductInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /product/{ID}][%d] editProductInternalServerError ", 500)
}

func (o *EditProductInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
