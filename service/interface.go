package service

import (
	"github.com/kashifkhan0771/ProductManagement/db"
)

// Service initializes our database instance.
type Service struct {
	db db.DataStore
}

// NewService creates a connection to our database.
func NewService(ds db.DataStore) *Service {
	return &Service{db: ds}
}
