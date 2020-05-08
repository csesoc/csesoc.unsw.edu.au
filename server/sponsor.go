package main

import (
	"context"
	"log"
	"net/http"
	"strings"
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
		expiryTime, _ := time.Parse(time.RFC3339, c.FormValue("expiry"))
		sponsor := Sponsor{
			SponsorName: strings.ToLower(c.FormValue("name")),
			SponsorLogo: c.FormValue("logo"),
			SponsorTier: c.FormValue("tier"),
			Expiry:      expiryTime.Unix(),
		}
		if err := c.Validate(sponsor); err != nil {
			return c.JSON(http.StatusBadRequest, H{})
		}
		// token := c.FormValue("token")
		if _, err := sponsorColl.InsertOne(context.TODO(), sponsor); err != nil {
			return c.JSON(http.StatusConflict, H{})
		}

		return c.JSON(http.StatusCreated, H{
			"response": "Created",
		})
	}

}

// GetSponsor - find entry for a specific sponsor.
func GetSponsor() echo.HandlerFunc {
	return func(c echo.Context) error {
		var result Sponsor
		// token := c.FormValue("token")
		name := strings.ToLower(c.FormValue("name"))
		filter := bson.D{{Key: "sponsorname", Value: name}}
		if err := sponsorColl.FindOne(context.TODO(), filter).Decode(&result); err != nil {
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
		// token := c.FormValue("token")
		name := strings.ToLower(c.FormValue("name"))
		filter := bson.D{{Key: "sponsorname", Value: name}}
		if _, err := sponsorColl.DeleteOne(context.TODO(), filter); err != nil {
			return c.JSON(http.StatusInternalServerError, H{
				"error": err,
			})
		}
		return c.JSON(http.StatusOK, H{
			"response": "Deleted",
		})
	}
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
