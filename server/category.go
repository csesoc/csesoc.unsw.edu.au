package main

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// GetCats - Retrieve a category from the database
func GetCats(collection *mongo.Collection, id int, token string) Category {
	// if !validToken(token) {
	// 	return nil
	// }

	var result Category
	filter := bson.D{{Key: "categoryID", Value: id}}

	// Find a category
	err := collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}

	return result
}

// NewCats - Add a new category
func NewCats(collection *mongo.Collection, catID int, index int, name string, token string) {
	// if !validToken(token) {
	// 	return
	// }

	category := Category{
		categoryID:   catID,
		categoryName: name,
		index:        index,
	}

	_, err := collection.InsertOne(context.TODO(), category)
	if err != nil {
		log.Fatal(err)
	}
}

// PatchCats - Update a category with new information
func PatchCats(collection *mongo.Collection, catID int, name string, index int, token string) {
	// if !validToken(token) {
	// 	return
	// }

	filter := bson.D{{Key: "categoryID", Value: catID}}
	update := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "categoryName", Value: name},
			{Key: "index", Value: index},
		}},
	}

	// Find a category by id and update it
	_, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
}

// DeleteCats - Delete a category from the database
func DeleteCats(collection *mongo.Collection, id int, token string) {
	// if !validToken(token) {
	// 	return
	// }

	filter := bson.D{{Key: "categoryID", Value: id}}

	// Find a category by id and delete it
	_, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
	}
}
