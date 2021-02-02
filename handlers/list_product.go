package handlers

import (
	"context"

	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"

	runtime "github.com/kashifkhan0771/ProductManagement"
	apiModels "github.com/kashifkhan0771/ProductManagement/gen/models"
	"github.com/kashifkhan0771/ProductManagement/gen/restapi/operations/service"
	domain "github.com/kashifkhan0771/ProductManagement/models"
)

// NewListTasksHandler handles request for listing Tasks
func NewListTasksHandler(ctx context.Context, rt *runtime.Runtime) service.ListTasksHandler {
	return &listTasks{
		ctx: ctx,
		rt:  rt,
	}
}

type listTasks struct {
	ctx context.Context
	rt  *runtime.Runtime
}

// Handle the list Tasks request
func (g *listTasks) Handle(params service.ListTasksParams) middleware.Responder {
	log().Debugf("request:'listTasks' params: %+v", params)
	productList := make([]*apiModels.Product, 0)

	products, err := g.rt.Service().ListProduct(g.ctx, extractTaskFilter(params), nil, int(swag.Int64Value(params.Offset)), int(swag.Int64Value(params.Limit)))
	if err != nil {
		log().Errorf("failed to list tasks: error[500]: %+v ", err)

		return service.NewListTasksInternalServerError()
	}

	for _, product := range products {
		productList = append(productList, ToProductGen(product))
	}

	return service.NewListTasksOK().WithPayload(productList)
}

// extractTaskFilter extracts filters from params
func extractTaskFilter(params service.ListTasksParams) map[string]interface{} {
	filter := make(map[string]interface{})
	if params.Status != nil {
		appendStatusFilter1(filter, swag.StringValue(params.Status))
	}

	return filter
}

// appendStatusFilter - append filter for agents based on status
func appendStatusFilter1(filter map[string]interface{}, param string) {
	switch param {
	case domain.OnSelling:
		filter[domain.OnSelling] = domain.OnSelling
	case domain.OffSelling:
		filter[domain.OffSelling] = domain.OffSelling
	}
}

// ToTaskGen converts domain to Task model
func ToProductGen(product *domain.Product) *apiModels.Product {
	return &apiModels.Product{
		ID:        product.ID,
		ProductID: product.ProductID,
		Name:      string(product.Name),
		Price:     product.Price,
		Status:    product.Status,
		Quantity:  product.Quantity,
	}
}
