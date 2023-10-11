package seed

import (
	// ...
	"context"
	"encoding/json"
	"io"
	"log"
	"os"
	"path/filepath"

	"go.mongodb.org/mongo-driver/mongo"
)

// func seedDatabase() error {
// 	// Read the JSON file
// 	data, err := os.ReadFile("brands.json")
// 	if err != nil {
// 		return err
// 	}

// 	// Parse the JSON data into an array of Brand objects
// 	var brands []Brand
// 	if err := json.Unmarshal(data, &brands); err != nil {
// 		return err
// 	}

// 	// Insert the records into the database
// 	collection := database.Collection("brands")
// 	_, err = collection.InsertMany(context.Background(), brands)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }


func SeedData(collection *mongo.Collection) error {

	// Get the directory path of the seed.go file
	// dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	dir, err := os.Getwd()
	if err != nil {
		return err
	}
	// Construct the path to the seed.json file
	log.Printf("this is current dir : %s",dir)
	filePath := filepath.Join(dir, "utils/seed.json")
	

	// Read the JSON file
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	// Read the JSON data from the file
	data, err := io.ReadAll(file)
	if err != nil {
	 return err
	}

	  

    // // Parse the JSON data
    // var items []interface{}
    // err = json.NewDecoder(file).Decode(&items)
    // if err != nil {
    //     return err
    // }

    // Parse the JSON data into an array of items
    var items []interface{}
    if err := json.Unmarshal(data, &items); err != nil {
        return err
    }

    // Insert the items into the collection
    _, err = collection.InsertMany(context.Background(), items)
    if err != nil {
        return err
    }

    return nil
}