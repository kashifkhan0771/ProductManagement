package handlers

import (
	"github.com/go-openapi/runtime/middleware"

	runtime "github.com/kashifkhan0771/ProductManagement"
	domainErr "github.com/kashifkhan0771/ProductManagement/errors"
	"github.com/kashifkhan0771/ProductManagement/gen/restapi/operations"
)

// NewEditHost function is for edit Host.
func NewEditHost(rt *runtime.Runtime) operations.EditProductHandler {
	return &editProduct{rt: rt}
}

type editProduct struct {
	rt *runtime.Runtime
}

// Handle the put Host request.
func (d *editProduct) Handle(params operations.EditProductParams) middleware.Responder {
	host, _ := d.rt.Service().RetrieveProduct(params.ID)
	if err := d.rt.Service().UpdateHost(host, params.ID); err != nil {
		switch apiErr := err.(*domainErr.APIError); {
		case apiErr.IsError(domainErr.NotFound):
			return operations.NewDeleteProductBadRequest()
		default:
			return operations.NewDeleteProductInternalServerError()
		}
	}

	return operations.NewEditProductOK()
}
