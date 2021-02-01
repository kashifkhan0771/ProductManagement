package handlers

import (
	"github.com/go-openapi/runtime/middleware"

	runtime "github.com/kashifkhan0771/ProductManagement"
	domainErr "github.com/kashifkhan0771/ProductManagement/errors"
	"github.com/kashifkhan0771/ProductManagement/gen/restapi/operations"
)

// NewDeleteHost function will delete the Host.
func NewDeleteHost(rt *runtime.Runtime) operations.DeleteProductHandler {
	return &deleteProduct{rt: rt}
}

type deleteProduct struct {
	rt *runtime.Runtime
}

// Handle the delete entry request.
func (d *deleteProduct) Handle(params operations.DeleteProductParams) middleware.Responder {
	if err := d.rt.Service().DeleteProduct(params.ID); err != nil {
		switch apiErr := err.(*domainErr.APIError); {
		case apiErr.IsError(domainErr.NotFound):
			return operations.NewDeleteProductBadRequest()
		default:
			return operations.NewDeleteProductInternalServerError()
		}
	}

	return operations.NewDeleteProductNoContent()
}
