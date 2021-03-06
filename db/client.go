package db

import (
	"context"
	"log"

	"github.com/kashifkhan0771/ProductManagement/models"
)

// DataStore is an interface for query ops.
type DataStore interface {
	AddProduct(Product *models.Product) (string, error)
	GetProductByID(id string) (*models.Product, error)
	ListProducts(ctx context.Context, filter map[string]interface{}, sortBy map[string]bool, offset, limit int) ([]*models.Product, error)
	DeleteProduct(id string) error
	UpdateProduct(Product *models.Product) error
}

// Option holds configuration for data store clients.
type Option struct {
	TestMode bool
}

// DataStoreFactory holds configuration for data store.
type DataStoreFactory func(conf Option) (DataStore, error)

var datastoreFactories = make(map[string]DataStoreFactory)

// Register saves data store into a data store factory.
func Register(name string, factory DataStoreFactory) {
	if factory == nil {
		log.Fatalf("Datastore factory %s does not exist.", name)

		return
	}
	_, ok := datastoreFactories[name]
	if ok {
		log.Fatalf("Datastore factory %s already registered. Ignoring.", name)

		return
	}
	datastoreFactories[name] = factory
}
