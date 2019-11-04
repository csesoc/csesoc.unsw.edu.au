package main

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetCat(collection *mongo.Collection, id int) *Category {

	var result *Category
	filter := bson.D{{"categoryID", id}}

	// Find a category
	err := collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}

	return result
}

func NewCat(collection *mongo.Collection, catID int, index int, name string) {
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

func PatchCat(collection *mongo.Collection, catID int, name string, index int) {
	filter := bson.D{{"categoryID", catID}}
	update := bson.D{
		{"$set", bson.D{
			{"categoryName", name},
			{"index", index},
		}},
	}

	// Find a category by id and update it
	_, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
}

func DeleteCat(collection *mongo.Collection, id int) {
	filter := bson.D{{"categoryID", id}}

	// Find a category by id and delete it
	_, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
	}
}
