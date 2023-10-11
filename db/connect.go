package connect

import (
	"context"
	"os"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// ConnectDB connects to MongoDB using the URI from the environment variable
func ConnectDB() (*mongo.Client, error) {
	// Get the MongoDB URI from the environment variable
	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		return nil, fmt.Errorf("MONGODB_URI environment variable not set")
	}
	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return nil, err
	}

	// Test the database connection
	err = client.Ping(context.Background(), nil)
	if err != nil {
		return nil, err
	}

	// Set the database variable
	// database = client.Database("brand_search")

	return client, nil
}
