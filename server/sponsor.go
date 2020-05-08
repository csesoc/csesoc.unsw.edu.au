package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var sponsorColl *mongo.Collection

// Setup
func SponsorSetup(client *mongo.Client) {
	sponsorColl = client.Database("csesoc").Collection("sponsors")
	opt := options.Index()
	opt.SetUnique(true)
	index := mongo.IndexModel{Keys: bson.M{"sponsorname": 1}, Options: opt}
	if _, err := sponsorColl.Indexes().CreateOne(context.Background(), index); err != nil {
		log.Println("Could not create index:", err)
	}
}

// NewSponsors - Add a sponsor
func NewSponsors() echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.FormValue("token")
		expiryStr := c.FormValue("expiry")
		name := c.FormValue("name")
		logo := c.FormValue("logo")
		tier := c.FormValue("tier")
		err := addSponsors(name, logo, tier, expiryStr, token)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, H{
				"error": err,
			})
		}
		return c.JSON(http.StatusCreated, H{})
	}

}

// DeleteSponsors - Delete a sponsor
func DeleteSponsors() echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.FormValue("token")
		sponsorName := c.FormValue("name")
		err := removeSponsors(sponsorName, token)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, H{
				"error": err,
			})
		}
		return c.JSON(http.StatusOK, H{})
	}
}

// addSponsors - Add a new sponsor
func addSponsors(name string, logo string, tier string, expiryStr string, token string) error {
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

	insertResult, err := sponsorColl.InsertOne(context.TODO(), sponsor)
	fmt.Println("Inserted a single document: ", insertResult.InsertedID)
	return err
}

// GetSponsor - Retrieve a list of sponsors from the database
func GetSponsor(sponsorName string, token string) (Sponsor, error) {
	var result Sponsor
	filter := bson.D{{Key: "sponsorname", Value: sponsorName}}
	err := sponsorColl.FindOne(context.TODO(), filter).Decode(&result)

	return result, err
}

// GetSponsors - Retrieve a sponsor from the database
func GetSponsors(token string) ([]*Sponsor, error) {
	var results []*Sponsor

	curr, err := sponsorColl.Find(context.TODO(), bson.D{{}}, options.Find())
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
func removeSponsors(sponsorName string, token string) error {
	// if !validToken(token) {
	// 	return
	// }

	// Find a sponsor by ID and delete it
	filter := bson.D{{Key: "sponsorname", Value: sponsorName}}
	_, err := sponsorColl.DeleteOne(context.TODO(), filter)
	return err
}
