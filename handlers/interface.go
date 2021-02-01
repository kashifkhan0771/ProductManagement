package handlers

import (
	"github.com/go-openapi/loads"

	runtime "github.com/kashifkhan0771/ProductManagement"
	"github.com/kashifkhan0771/ProductManagement/gen/restapi/operations"
)

// Handler replaces swagger handler.
type Handler *operations.ProductManagementAPI

// NewHandler overrides swagger api handlers.
func NewHandler(rt *runtime.Runtime, spec *loads.Document) Handler {
	handler := operations.NewProductManagementAPI(spec)

	// Host handlers
	handler.AddProductHandler = NewAddProduct(rt)
	handler.GetProductHandler = NewGetProduct(rt)
	handler.EditProductHandler = NewEditHost(rt)
	handler.DeleteProductHandler = NewDeleteHost(rt)

	return handler
}
