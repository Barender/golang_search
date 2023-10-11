package controller

import (
	"context"
	"log"
	"net/http"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type product struct {
	BrandName   string `bson:"brand_name" json:"brand_name"`
	ProductName string `bson:"product_name" json:"product_name"`
	Category    string `bson:"category" json:"category"`
	Location    string `bson:"location" json:"location"`
	Score       int    `bson:"-" json:"-"` // Add this line
}

func GetFilteredProducts(collection *mongo.Collection) func(c *gin.Context) {
	return func(c *gin.Context) {
		var filterParams bson.M
		err := c.BindJSON(&filterParams)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
			return
		}

		searchString := filterParams["search"].(string)
		searchWords := strings.Fields(searchString)
		items := make(map[string]*product)

		cursor, err := collection.Find(context.Background(), bson.M{})
		if err != nil {
			log.Fatalf("Failed to retrieve items from MongoDB: %v", err)
		}
		defer cursor.Close(context.Background())

		for cursor.Next(context.Background()) {
			var item product
			if err := cursor.Decode(&item); err != nil {
				log.Fatalf("Failed to decode item: %v", err)
			}
			item.Score = 0
			matchedFields := make(map[string]int)

			itemValue := reflect.ValueOf(item)
			for i := 0; i < itemValue.NumField(); i++ {
				fieldValue := itemValue.Field(i).String()
				for _, word := range searchWords {
					if strings.Contains(strings.ToLower(fieldValue), strings.ToLower(word)) {
						matchedFields[itemValue.Type().Field(i).Name]++
					}
				}
			}

			for _, count := range matchedFields {
				if count > 0 {
					item.Score++
				}
			}

			if len(searchWords) > 1 && item.Score >= 2 {
				items[item.ProductName] = &item
			} else if len(searchWords) == 1 && item.Score > 0 {
				items[item.ProductName] = &item
			}
		}

		var filteredItems []product
		for _, item := range items {
			filteredItems = append(filteredItems, *item)
		}

		if len(filteredItems) == 0 {
			c.JSON(http.StatusNotFound, gin.H{"error": "No matching documents found"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"product": filteredItems})
	}
}

func GetAllProducts(collection *mongo.Collection) func(c *gin.Context) {

	return func(c *gin.Context) {
		// collection := Connect.ConnectDB().Database("mydb").Collection("items")

		// Implement your MongoDB query logic here
		// For example, you can use the Find method to retrieve all items from the collection
		cursor, err := collection.Find(context.Background(), bson.M{})
		if err != nil {
			log.Fatalf("Failed to retrieve items from MongoDB: %v", err)
		}
		defer cursor.Close(context.Background())

		// Iterate over the cursor and process each item
		var items []product
		for cursor.Next(context.Background()) {
			var item product
			if err := cursor.Decode(&item); err != nil {
				log.Fatalf("Failed to decode item: %v", err)
			}
			items = append(items, item)
		}

		// Example response
		c.JSON(http.StatusOK, gin.H{"products": items})

	}
}