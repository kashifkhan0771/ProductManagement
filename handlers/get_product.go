package handlers

import (
	"github.com/go-openapi/runtime/middleware"

	runtime "github.com/kashifkhan0771/ProductManagement"
	domainErr "github.com/kashifkhan0771/ProductManagement/errors"
	"github.com/kashifkhan0771/ProductManagement/gen/models"
	"github.com/kashifkhan0771/ProductManagement/gen/restapi/operations"
)

// NewGetHost handles a request for retrieving Host.
func NewGetProduct(rt *runtime.Runtime) operations.GetProductHandler {
	return &getProduct{rt: rt}
}

type getProduct struct {
	rt *runtime.Runtime
}

// Handle the get Host request.
func (d *getProduct) Handle(params operations.GetProductParams) middleware.Responder {
	product, err := d.rt.Service().RetrieveProduct(params.ID)
	if err != nil {
		switch apiErr := err.(*domainErr.APIError); {
		case apiErr.IsError(domainErr.NotFound):
			return operations.NewGetProductNotFound()
		default:
			return operations.NewGetProductInternalServerError()
		}
	}

	return operations.NewGetProductOK().WithPayload(&models.Product{
		ID:        product.ID,
		ProductID: product.ProductID,
		Name:      product.Name,
		Price:     product.Price,
		Status:    product.Status,
		Quantity:  product.Quantity,
	})
}
