package models

import (
	"github.com/fatih/structs"
)

// Product Model.
type Product struct {
	ID        string `structs:"id" json:"id" bson:"_id" db:"id"`
	ProductID string `structs:"product_id" json:"product_id" bson:"product_id" db:"product_id"`
	Name      string `structs:"name" json:"name" bson:"name" db:"name"`
	Price     int64  `structs:"price" json:"price" bson:"price" db:"price"`
	Status    string `structs:"status" json:"status" bson:"status" db:"status"` // Can be : OnSelling, OffSelling
	Quantity  int64  `structs:"quantity" json:"quantity" bson:"quantity" db:"quantity"`
}

// Map convert struct into map.
func (product *Product) Map() map[string]interface{} {
	return structs.Map(product)
}

// Names return field names of Host model.
func (product *Product) Names() []string {
	fields := structs.Fields(product)

	names := make([]string, len(fields))
	for i, field := range fields {
		name := field.Name()
		tagName := field.Tag(structs.DefaultTagName)
		if tagName != "" {
			name = tagName
		}
		names[i] = name
	}

	return names
}
