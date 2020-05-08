package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Sponsor - struct to contain sponsor data
type Sponsor struct {
	SponsorName string `validate:"required"`
	SponsorLogo string `validate:"required"`
	SponsorTier string `validate:"required"`
	Expiry      int64
}

var sponsorColl *mongo.Collection

/* Setup */

// SponsorSetup - Setup the collection to be used for sponsors
func SponsorSetup(client *mongo.Client) {
	sponsorColl = client.Database("csesoc").Collection("sponsors")
	opt := options.Index()
	opt.SetUnique(true)
	index := mongo.IndexModel{Keys: bson.M{"sponsorname": 1}, Options: opt}
	if _, err := sponsorColl.Indexes().CreateOne(context.Background(), index); err != nil {
		log.Fatal("Could not create index: ", err)
	}
}

/* Handles */

// NewSponsor - Add a sponsor
func NewSponsor() echo.HandlerFunc {
	return func(c echo.Context) error {
		name := c.FormValue("name")
		logo := c.FormValue("logo")
		tier := c.FormValue("tier")
		expiryStr := c.FormValue("expiry")
		token := c.FormValue("token")
		err := addSponsors(name, logo, tier, expiryStr, token)
		if err != nil {
			return c.JSON(http.StatusConflict, H{})
		}
		return c.JSON(http.StatusCreated, H{})
	}

}

// GetSponsor - find entry for a specific sponsor.
func GetSponsor() echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.FormValue("token")
		name := c.FormValue("name")
		result, err := getSponsor(name, token)
		if err != nil {
			return c.JSON(http.StatusNotFound, H{
				"response": "No such sponsor.",
			})
		}
		return c.JSON(http.StatusOK, H{
			"sponsors": result,
		})
	}
}

// GetSponsors - gives a list of sponsors stored.
func GetSponsors() echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.FormValue("token")
		results, err := getSponsors(token)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, H{})
		}
		content := H{}
		if results != nil {
			content = H{
				"sponsors": results,
			}
		}
		return c.JSON(http.StatusOK, content)
	}
}

// DeleteSponsor - Delete a sponsor
func DeleteSponsor() echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.FormValue("token")
		sponsorName := c.FormValue("name")
		err := removeSponsors(sponsorName, token)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, H{
				"error": err,
			})
		}
		return c.JSON(http.StatusOK, H{
			"response": "Deleted " + sponsorName,
		})
	}
}

/* Database Queries */

// addSponsors - Add a new sponsor
func addSponsors(name string, logo string, tier string, expiryStr string, token string) error {
	// should validating be done by the handler or database function
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

	_, err := sponsorColl.InsertOne(context.TODO(), sponsor)
	return err
}

// getSponsor - Retrieve a list of sponsors from the database
func getSponsor(sponsorName string, token string) (Sponsor, error) {
	var result Sponsor
	filter := bson.D{{Key: "sponsorname", Value: sponsorName}}
	err := sponsorColl.FindOne(context.TODO(), filter).Decode(&result)

	return result, err
}

// getSponsors - Retrieve a sponsor from the database
func getSponsors(token string) ([]*Sponsor, error) {
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
