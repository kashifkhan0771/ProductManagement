package handlers

import (
	"github.com/go-openapi/runtime/middleware"

	runtime "github.com/kashifkhan0771/ProductManagement"
	"github.com/kashifkhan0771/ProductManagement/gen/restapi/operations"
	"github.com/kashifkhan0771/ProductManagement/models"
)

// NewAddHost function will add new host.
func NewAddProduct(rt *runtime.Runtime) operations.AddProductHandler {
	return &addProduct{rt: rt}
}

type addProduct struct {
	rt *runtime.Runtime
}

// Handle the add student request.
func (d *addProduct) Handle(params operations.AddProductParams) middleware.Responder {
	host := models.Product{
		ID:        params.Product.ID,
		ProductID: params.Product.ProductID,
		Name:      params.Product.Name,
		Price:     params.Product.Price,
		Status:    params.Product.Status,
		Quantity:  params.Product.Quantity,
	}

	id, err := d.rt.Service().AddProduct(&host)
	if err != nil {
		log().Errorf("failed to add Product: %s", err)

		return operations.NewAddProductBadRequest()
	}

	params.Product.ID = id

	return operations.NewAddProductCreated().WithPayload(params.Product)
}
