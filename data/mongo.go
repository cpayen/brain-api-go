package data

import (
	"brain-api/models"
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const dbName = "Brain01"
const contentCollection = "Content"

var db *mongo.Database

// Connect establish a connection to database
func init() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("BRAIN_DB")))
	if err != nil {
		log.Fatal(err)
	}
	db = client.Database(dbName)
}

// InsertOneFolder inserts one folder from Folder model
func InsertOneFolder(folder *models.Folder) string {
	folder.CreatedAt = time.Now()
	folder.UpdatedAt = time.Now()
	fmt.Println(folder)
	result, err := db.Collection(contentCollection).InsertOne(context.Background(), folder)
	if err != nil {
		log.Fatal(err)
	}
	newID := result.InsertedID
	return string(newID.(primitive.ObjectID).Hex())
}
