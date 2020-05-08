package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Handlers
func NewSponsors(collection *mongo.Collection) echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.FormValue("token")
		expiryStr := c.FormValue("expiry")
		name := c.FormValue("name")
		logo := c.FormValue("logo")
		tier := c.FormValue("tier")
		err := addSponsors(collection, name, logo, tier, expiryStr, token)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, H{
				"error": err,
			})
		}
		return c.JSON(http.StatusCreated, H{})
	}

}

// DeleteSponsors - Delete a sponsor from the database
func DeleteSponsors(collection *mongo.Collection) echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.FormValue("token")
		sponsorName := c.FormValue("name")
		err := removeSponsors(collection, sponsorName, token)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, H{
				"error": err,
			})
		}
		return c.JSON(http.StatusOK, H{})
	}
}

// addSponsors - Add a new sponsor
func addSponsors(collection *mongo.Collection, name string, logo string, tier string, expiryStr string, token string) error {
	// if !validToken(token) {
	// 	return
	// }

	expiryTime, _ := time.Parse(time.RFC3339, expiryStr)
	sponsor := Sponsor{
		SponsorName: name,
		SponsorLogo: logo,
		SponsorTier: tier,
		Expiry:      expiryTime.Unix(),
	}

	insertResult, err := collection.InsertOne(context.TODO(), sponsor)
	fmt.Println("Inserted a single document: ", insertResult.InsertedID)
	return err
}

// GetSponsor - Retrieve a list of sponsors from the database
func GetSponsor(collection *mongo.Collection, sponsorName string, token string) (Sponsor, error) {
	var result Sponsor
	filter := bson.D{{Key: "sponsorname", Value: sponsorName}}
	err := collection.FindOne(context.TODO(), filter).Decode(&result)

	return result, err
}

// GetSponsors - Retrieve a sponsor from the database
func GetSponsors(collection *mongo.Collection, token string) ([]*Sponsor, error) {
	var results []*Sponsor

	curr, err := collection.Find(context.TODO(), bson.D{{}}, options.Find())
	if err == nil {
		for curr.Next(context.TODO()) {

			var elem Sponsor
			curr.Decode(&elem)
			results = append(results, &elem)
		}
	}
	return results, err
}

// removeSponsors - Remove a sponsor from the database
func removeSponsors(collection *mongo.Collection, sponsorName string, token string) error {
	// if !validToken(token) {
	// 	return
	// }

	// Find a sponsor by ID and delete it
	filter := bson.D{{Key: "sponsorname", Value: sponsorName}}
	_, err := collection.DeleteOne(context.TODO(), filter)
	return err
}
