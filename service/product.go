package service

import (
	"context"

	"github.com/kashifkhan0771/ProductManagement/models"
)

// AddHost add into db.
func (s *Service) AddProduct(product *models.Product) (string, error) {
	return s.db.AddProduct(product)
}

// DeleteHost delete from db.
func (s *Service) DeleteProduct(id string) error {
	if _, err := s.db.GetProductByID(id); err != nil {
		return err
	}

	return s.db.DeleteProduct(id)
}

// RetrieveHost get host from db.
func (s *Service) RetrieveProduct(id string) (*models.Product, error) {
	product, err := s.db.GetProductByID(id)
	if err != nil {
		return nil, err
	}

	return product, nil
}

// ListTasks retrieves list of task
func (s *Service) ListProduct(ctx context.Context, filter map[string]interface{}, sortBy map[string]bool, offset, limit int) ([]*models.Product, error) {
	return s.db.ListProducts(ctx, filter, sortBy, offset, limit)
}

// UpdateHost update in db.
func (s *Service) UpdateHost(product *models.Product, id string) error {
	if _, err := s.RetrieveProduct(id); err != nil {
		return err
	}

	return s.db.UpdateProduct(product)
}
