// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// DeleteProductHandlerFunc turns a function with the right signature into a delete product handler
type DeleteProductHandlerFunc func(DeleteProductParams) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteProductHandlerFunc) Handle(params DeleteProductParams) middleware.Responder {
	return fn(params)
}

// DeleteProductHandler interface for that can handle valid delete product params
type DeleteProductHandler interface {
	Handle(DeleteProductParams) middleware.Responder
}

// NewDeleteProduct creates a new http.Handler for the delete product operation
func NewDeleteProduct(ctx *middleware.Context, handler DeleteProductHandler) *DeleteProduct {
	return &DeleteProduct{Context: ctx, Handler: handler}
}

/* DeleteProduct swagger:route DELETE /product/{ID} deleteProduct

DeleteProduct delete product API

*/
type DeleteProduct struct {
	Context *middleware.Context
	Handler DeleteProductHandler
}

func (o *DeleteProduct) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewDeleteProductParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
