package data

import (
	"brain-api/models"
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
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

// GetContentItem get one content item from Content interface
func GetContentItem(id string) (bson.M, error) {
	objID, _ := primitive.ObjectIDFromHex(id)
	result := db.Collection(contentCollection).FindOne(context.Background(), bson.M{"_id": objID})
	if result.Err() != nil {
		return nil, result.Err()
	}
	doc := bson.M{}
	err := result.Decode(&doc)
	return doc, err
}

// UpdateContentItem update one content item from Content interface
func UpdateContentItem(id string, item models.Content) (bson.M, error) {
	item.SetUpdateDate(time.Now())
	objID, _ := primitive.ObjectIDFromHex(id)
	update := bson.M{"$set": item}
	upsert := true
	after := options.After
	opt := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
		Upsert:         &upsert,
	}
	result := db.Collection(contentCollection).FindOneAndUpdate(context.Background(), bson.M{"_id": objID}, update, &opt)
	if result.Err() != nil {
		log.Fatal(result.Err())
		return nil, result.Err()
	}
	doc := bson.M{}
	err := result.Decode(&doc)
	return doc, err
}

// InsertContentItem inserts one content item from Content interface
func InsertContentItem(item models.Content) (bson.M, error) {
	item.SetCreationDate(time.Now())
	item.SetUpdateDate(time.Now())
	result, err := db.Collection(contentCollection).InsertOne(context.Background(), item)
	if err != nil {
		log.Fatal(err)
	}
	newID := result.InsertedID
	id := string(newID.(primitive.ObjectID).Hex())
	return GetContentItem(id)
}
