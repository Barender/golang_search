package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	Controllers "mygo/controllers"
	Connect "mygo/db"
	Middleware "mygo/middleware"
	Seed "mygo/utils"
)

type Item struct {
	BrandName   string `bson:"brand_name" json:"brand_name"`
	ProductName string `bson:"product_name" json:"product_name"`
	Category    string `bson:"category" json:"category"`
	Location    string `bson:"location" json:"location"`
}

func init() {
	// Get the current working directory
	path_dir, err := os.Getwd()
	if err != nil {
		log.Fatalf("Error getting current directory: %v", err)
	}

	// Load environment variables from the .env file
	if err := godotenv.Load(filepath.Join(path_dir, ".env")); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func main() {
	// Initialize the Gin router
	router := gin.Default()
	router.Use(CORSMiddleware())

	// Middleware to log network requests
	router.Use(Middleware.RequestLoggerMiddleware())

	// Connect to MongoDB
	client, err := Connect.ConnectDB()
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}
	defer client.Disconnect(context.Background())

	// Get the collection
	collection := client.Database("golang_search").Collection("items")

	// Create a text index on the fields you want to search
	_, err = collection.Indexes().CreateOne(
		context.Background(),
		mongo.IndexModel{
			Keys: bson.D{
				{"brand_name", "text"},
				{"product_name", "text"},
				{"category", "text"},
				{"location", "text"},
			},
		},
	)
	if err != nil {
		log.Fatal(err)
	}

	// check for seeding Initial Data.
	count, err := collection.CountDocuments(context.Background(), bson.M{})
	if err != nil {
		log.Fatalf("Failed to count documents in collection: %v", err)
	}
	// If there is no data in the collection, seed it
	if count == 0 {
		log.Printf("Seeding JSON data into database for first time. %v", count)
		err = Seed.SeedData(collection)
		if err != nil {
			log.Fatalf("Failed to seed data: %v", err)
		}
	}

	// Define API routes
	// Root route
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome to the API!",
		})
	})

	// Example route: Get all items from an "items" collection
	router.GET("/getAllProducts", Controllers.GetAllProducts(collection))

	// filteredProducts route: Get all items from an "items" collection
	router.POST("/getFilteredProducts", Controllers.GetFilteredProducts(collection))

	// Start the server
	port := os.Getenv("PORT")
	if port == "" {
		port = "4000"
	}

	addr := fmt.Sprintf(":%s", port)
	err = router.Run(addr)
	if err != nil {
		log.Fatal(err)
	}
}
