package main

import (
	"context"
	"log"
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// NewSponsor - Add a new sponsor
func NewSponsor(collection *mongo.Collection, expiryStr string, name string, logo string, tier string, token string) {
	if !validToken(token) {
		return
	}

	expiryTime, _ := time.Parse(time.RFC3339, expiryStr)
	id := uuid.New()

	sponsor := Sponsor{
		sponsorID:   id,
		sponsorName: name,
		sponsorLogo: logo,
		sponsorTier: tier,
		expiry:      expiryTime.Unix(),
	}

	_, err := collection.InsertOne(context.TODO(), sponsor)
	if err != nil {
		log.Fatal(err)
	}
}

// DeleteSponsor - Delete a sponsor from the database
func DeleteSponsor(collection *mongo.Collection, id string, token string) {
	if !validToken(token) {
		return
	}

	parsedID := uuid.Must(uuid.Parse(id))

	// Find a sponsor by ID and delete it
	filter := bson.D{{"sponsorID", parsedID}}
	_, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
	}
}
