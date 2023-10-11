package main

import "go.mongodb.org/mongo-driver/bson/primitive"

// Brand represents a brand record in the database.
type Brand struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	BrandName   string             `bson:"brand_name"`
	ProductName string             `bson:"product_name"`
	Category    string             `bson:"category"`
	Location    string             `bson:"location"`
}

// NewBrand creates a new Brand instance with the given details.
func NewBrand(brandName, productName, category, location string) *Brand {
	return &Brand{
		BrandName:   brandName,
		ProductName: productName,
		Category:    category,
		Location:    location,
	}
}
