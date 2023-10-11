package Connect

import (
	"context"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// ConnectDB connects to MongoDB using the URI from the environment variable
func ConnectDB(mongoURI string) (*mongo.Client, error) {
	// If mongoURI is not provided, get it from the environment variable
	if mongoURI == "" {
		mongoURI = os.Getenv("MONGODB_URI")
		if mongoURI == "" {
			return nil, fmt.Errorf("MONGODB_URI environment variable not set")
		}
	}

	clientOptions := options.Client().ApplyURI(mongoURI)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return nil, err
	}

	// Test the database connection
	err = client.Ping(context.Background(), nil)
	if err != nil {
		return nil, err
	}

	return client, nil
}
