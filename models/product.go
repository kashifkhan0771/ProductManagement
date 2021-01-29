package models

import (
	"github.com/fatih/structs"
)

type metadata map[string]interface{}

// Host Model.
type Product struct {
	ID       string   `structs:"id" json:"id" bson:"_id" db:"id"`
	Name     string   `structs:"name" json:"name" bson:"name" db:"name"`
	Price 	 int 	  `structs:"price" json:"price" bson:"price" db:"price"`
	Status   string   `structs:"status" json:"status" bson:"status" db:"status"` // Can be : OnSelling, OffSelling
	Quantity int      `structs:"quantity" json:"quantity" bson:"quantity" db:"quantity"`
}

// Map convert struct into map.
func (prd *Product) Map() map[string]interface{} {
	return structs.Map(prd)
}

// Names return field names of Host model.
func (prd *Product) Names() []string {
	fields := structs.Fields(prd)

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