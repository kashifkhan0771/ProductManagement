// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"github.com/kashifkhan0771/ProductManagement/gen/restapi/operations"
)

//go:generate swagger generate server --target ../../gen --name ProductManagement --spec ../../swagger.yml --principal interface{} --exclude-main

func configureFlags(api *operations.ProductManagementAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.ProductManagementAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.UseSwaggerUI()
	// To continue using redoc as your UI, uncomment the following line
	// api.UseRedoc()

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	if api.AddProductHandler == nil {
		api.AddProductHandler = operations.AddProductHandlerFunc(func(params operations.AddProductParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.AddProduct has not yet been implemented")
		})
	}
	if api.DeleteProductHandler == nil {
		api.DeleteProductHandler = operations.DeleteProductHandlerFunc(func(params operations.DeleteProductParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.DeleteProduct has not yet been implemented")
		})
	}
	if api.EditProductHandler == nil {
		api.EditProductHandler = operations.EditProductHandlerFunc(func(params operations.EditProductParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.EditProduct has not yet been implemented")
		})
	}
	if api.GetProductHandler == nil {
		api.GetProductHandler = operations.GetProductHandlerFunc(func(params operations.GetProductParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.GetProduct has not yet been implemented")
		})
	}

	api.PreServerShutdown = func() {}

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix".
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation.
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics.
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
