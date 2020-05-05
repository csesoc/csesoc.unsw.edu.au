package main

import (
	"context"
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// GetSponsors - Retrieve a sponsor from the database
func GetSponsors(collection *mongo.Collection, id string, token string) (Sponsor, error) {
	parsedID := uuid.Must(uuid.Parse(id))

	var result Sponsor
	filter := bson.D{{Key: "sponsorid", Value: parsedID}}
	err := collection.FindOne(context.TODO(), filter).Decode(&result)

	return result, err
}

// NewSponsors - Add a new sponsor
func NewSponsors(collection *mongo.Collection, expiryStr string, name string, logo string, tier string, token string) error {
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
	return err
}

// DeleteSponsors - Delete a sponsor from the database
func DeleteSponsors(collection *mongo.Collection, id string, token string) error {
	// if !validToken(token) {
	// 	return
	// }

	parsedID := uuid.Must(uuid.Parse(id))

	// Find a sponsor by ID and delete it
	filter := bson.D{{Key: "sponsorid", Value: parsedID}}
	_, err := collection.DeleteOne(context.TODO(), filter)
	return err
}
