package mongo

import (
	"os"
	"reflect"
	"testing"

	"github.com/kashifkhan0771/ProductManagement/db"
	"github.com/kashifkhan0771/ProductManagement/models"
)

func Test_client_AddProduct(t *testing.T) {
	t.Parallel()
	_ = os.Setenv("DB_PORT", "27017")
	_ = os.Setenv("DB_HOST", "product-management-mongo-db")

	type args struct {
		product *models.Product
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Success - Add Product to db",
			args: args{product: &models.Product{
				ProductID: "Product 001",
				Name: "product name 1",
				Price: 2300,
				Status: "OnSelling",
				Quantity: 34,
			}},
			wantErr: false,

		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			c, _ := NewClient(db.Option{})
			if c != nil {
				_, err := c.AddProduct(tt.args.product)
				if (err != nil) != tt.wantErr {
					t.Errorf("AddProduct() error = %v, wantErr %v", err, tt.wantErr)

					return
				}
			}
		})
	}
}

func Test_client_DeleteHost(t *testing.T) {
	t.Parallel()
	_ = os.Setenv("DB_PORT", "27017")
	_ = os.Setenv("DB_HOST", "product-management-mongo-db")
	c, _ := NewClient(db.Option{})
	product := &models.Product{ProductID: "Product 002", Name: "Product Name 2",
		Price: 4500, Status: "OffSelling", Quantity: 12}
	if c != nil {
		_, _ = c.AddProduct(product)
	}
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "Success - delete product from db",
			args:    args{id: product.ID},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if err := c.DeleteProduct(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("DeleteProduct() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_client_GetProductByID(t *testing.T) {
	t.Parallel()
	_ = os.Setenv("DB_PORT", "27017")
	_ = os.Setenv("DB_HOST", "product-management-mongo-db")

	c, _ := NewClient(db.Option{})
	product := &models.Product{ProductID: "Product 003", Name: "Product Name 3", Price: 1399, Status: "OnSelling",
		Quantity: 340}
	if c != nil {
		_, _ = c.AddProduct(product)
	}
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		args    args
		want    *models.Product
		wantErr bool
	}{
		{
			name:    "Success - get product from db",
			args:    args{product.ID},
			want:    product,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, err := c.GetProductByID(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetProduct() error = %v, wantErr %v", err, tt.wantErr)

				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetProduct() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_client_UpdateProduct(t *testing.T) {
	t.Parallel()
	_ = os.Setenv("DB_PORT", "27017")
	_ = os.Setenv("DB_HOST", "product-management-mongo-db")

	c, _ := NewClient(db.Option{})
	product := &models.Product{ProductID: "Product 004", Name: "Product Update", Price: 799, Status: "OffSelling",
		Quantity: 78}

	updateProductID, _ := c.AddProduct(product)

	type args struct {
		host *models.Product
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Success - update product in db",
			args: args{host: &models.Product{ID: updateProductID, ProductID: "Product 004",
				Name: "Product Update", Price: 750, Status: "OnSelling", Quantity: 100}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		if c != nil {
			if err := c.UpdateProduct(tt.args.host); (err != nil) != tt.wantErr {
				t.Errorf("UpdateProduct() error = %v, wantErr %v", err, tt.wantErr)
			}
		}
	}
}
