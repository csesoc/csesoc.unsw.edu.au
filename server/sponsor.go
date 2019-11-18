package main

import (
	"context"
	"log"
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// GetSponsors - Retrieve a sponsor from the database
func GetSponsors(collection *mongo.Collection, id string, token string) Sponsor {
	parsedID := uuid.Must(uuid.Parse(id))

	var result Sponsor
	filter := bson.D{{Key: "sponsorid", Value: parsedID}}
	err := collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}

	return result
}

// NewSponsors - Add a new sponsor
func NewSponsors(collection *mongo.Collection, expiryStr string, name string, logo string, tier string, token string) {
	// if !validToken(token) {
	// 	return
	// }

	expiryTime, _ := time.Parse(time.RFC3339, expiryStr)
	id := uuid.New()

	sponsor := Sponsor{
		SponsorID:   id,
		SponsorName: name,
		SponsorLogo: logo,
		SponsorTier: tier,
		Expiry:      expiryTime.Unix(),
	}

	_, err := collection.InsertOne(context.TODO(), sponsor)
	if err != nil {
		log.Fatal(err)
	}
}

// DeleteSponsors - Delete a sponsor from the database
func DeleteSponsors(collection *mongo.Collection, id string, token string) {
	// if !validToken(token) {
	// 	return
	// }

	parsedID := uuid.Must(uuid.Parse(id))

	// Find a sponsor by ID and delete it
	filter := bson.D{{Key: "sponsorid", Value: parsedID}}
	_, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
	}
}
