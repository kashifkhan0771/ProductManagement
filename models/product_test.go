package models

import (
	"reflect"
	"testing"
)

func TestProduct_Map(t *testing.T) {
	t.Parallel()
	type fields struct {
		ID        string
		ProductID string
		Name      string
		Price     int
		Status    string
		Quantity  int
	}
	tests := []struct {
		name   string
		fields fields
		want   map[string]interface{}
	}{
		{
			name: "Success - Get Product 1 Map",
			fields: fields{
<<<<<<< HEAD
				ID: "001",
				ProductID: "Prd 001",
				Name: "Product Name 1",
				Price: 2350,
				Status: "OnSelling",
				Quantity: 23,
			},
			want: map[string]interface{}{
				"id": "001",
				"product_id": "Prd 001",
				"name": "Product Name 1",
				"price": 2350,
				"status": "OnSelling",
				"quantity": 23,
=======
				ID:        "001",
				ProductID: "Prd 001",
				Name:      "Product Name 1",
				Price:     2350,
				Status:    "OnSelling",
				Quantity:  23,
			},
			want: map[string]interface{}{
				"id":         "001",
				"product_id": "Prd 001",
				"name":       "Product Name 1",
				"price":      2350,
				"status":     "OnSelling",
				"quantity":   23,
>>>>>>> feature/docker
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			prd := &Product{
				ID:        tt.fields.ID,
				ProductID: tt.fields.ProductID,
				Name:      tt.fields.Name,
				Price:     tt.fields.Price,
				Status:    tt.fields.Status,
				Quantity:  tt.fields.Quantity,
			}
			if got := prd.Map(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Map() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProduct_Names(t *testing.T) {
	t.Parallel()
	type fields struct {
		ID        string
		ProductID string
		Name      string
		Price     int
		Status    string
		Quantity  int
	}
	tests := []struct {
		name   string
		fields fields
		want   []string
	}{
		{
			name: "Success - Get Product Field Names",
			fields: fields{
<<<<<<< HEAD
				ID: "001",
				ProductID: "Prod 001",
				Name: "Product Name 1",
				Price: 2350,
				Status: "OnSelling",
				Quantity: 23,
=======
				ID:        "001",
				ProductID: "Prod 001",
				Name:      "Product Name 1",
				Price:     2350,
				Status:    "OnSelling",
				Quantity:  23,
>>>>>>> feature/docker
			},
			want: []string{"id", "product_id", "name", "price", "status", "quantity"},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			prd := &Product{
				ProductID: tt.fields.ProductID,
				Name:      tt.fields.Name,
				Price:     tt.fields.Price,
				Status:    tt.fields.Status,
				Quantity:  tt.fields.Quantity,
			}
			if got := prd.Names(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Names() = %v, want %v", got, tt.want)
			}
		})
	}
}
