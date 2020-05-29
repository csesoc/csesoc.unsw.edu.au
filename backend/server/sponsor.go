package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Sponsor - struct to contain sponsor data
type Sponsor struct {
	Name   string `json:"name" validate:"required"`
	Logo   string `json:"logo" validate:"required,url"`
	Tier   int    `json:"tier" validate:"required,numeric,eq=10|eq=100|eq=1000"`
	Expiry int64  `json:"expiry"` // check because if this will be null most of the time then what would this number be?
}

var sponsorColl *mongo.Collection

/* Setup */

// SponsorSetup - Setup the collection to be used for sponsors
func SponsorSetup(client *mongo.Client) {
	sponsorColl = client.Database("csesoc").Collection("sponsors")

	// Creating unique index for sponsor name
	opt := options.Index()
	opt.SetUnique(true)
	index := mongo.IndexModel{
		Keys:    bson.M{"name": 1},
		Options: opt,
	}
	if _, err := sponsorColl.Indexes().CreateOne(context.Background(), index); err != nil {
		log.Fatal("Could not create index: ", err)
	}

	// Fetching sponsor list
	resp, err := http.Get("https://gistcdn.githack.com/esyw/4e35cd5fe73fa024020e67855ca733fb/raw/e85c9ae58a6323a4214ffa4ad89b0a5ebe404e31/sponsors.json")
	if err != nil {
		log.Fatal("Could not get sponsor list: ", err)
	}
	defer resp.Body.Close()

	var sponsors []Sponsor
	if err = json.NewDecoder(resp.Body).Decode(&sponsors); err != nil {
		log.Printf("Could not convert JSON response to Sponsors")
	}

	for _, sponsor := range sponsors {
		log.Println(sponsor)
		if _, err := sponsorColl.InsertOne(context.TODO(), sponsor); err != nil {
			log.Printf("Could not insert sponsor " + sponsor.Name + " " + err.Error())
		}
	}

}

/* Handles */

// NewSponsor - Add a sponsor
func NewSponsor(c echo.Context) error {
	tier, err := strconv.Atoi(c.FormValue("tier"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, H{
			"error": "Tier has to be a number",
		})
	}
	expiryTime, _ := time.Parse(time.RFC3339, c.FormValue("expiry"))
	sponsor := Sponsor{
		Name:   c.FormValue("name"),
		Logo:   c.FormValue("logo"),
		Tier:   tier,
		Expiry: expiryTime.Unix(),
	}

	// validate the struct with golang validator package
	if err := c.Validate(sponsor); err != nil {
		return c.JSON(http.StatusBadRequest, H{
			"error": "Bad request",
		})
	}
	// token := c.FormValue("token")

	if _, err := sponsorColl.InsertOne(context.TODO(), sponsor); err != nil {
		return c.JSON(http.StatusConflict, H{})
	}

	return c.JSON(http.StatusCreated, H{
		"response": "Created",
	})
}

// GetSponsor - find entry for a specific sponsor.
func GetSponsor(c echo.Context) error {
	var result Sponsor
	// token := c.FormValue("token")
	filter := bson.D{{Key: "name", Value: c.Param("name")}}
	if err := sponsorColl.FindOne(context.TODO(), filter).Decode(&result); err != nil {
		return c.JSON(http.StatusNotFound, H{
			"response": "No such sponsor.",
		})
	}
	return c.JSON(http.StatusOK, result)
}

// GetSponsors - gives a list of sponsors stored.
func GetSponsors(c echo.Context) error {
	token := c.FormValue("token")
	tier := c.FormValue("tier")
	results, err := retrieveSponsors(token, tier)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, H{})
	}
	if results != nil {
		return c.JSON(http.StatusOK, results)
	}
	return c.JSON(http.StatusOK, H{})
}

// DeleteSponsor - Delete a sponsor
func DeleteSponsor(c echo.Context) error {
	// token := c.FormValue("token")
	filter := bson.D{{Key: "name", Value: c.Param("name")}}
	if _, err := sponsorColl.DeleteOne(context.TODO(), filter); err != nil {
		return c.JSON(http.StatusInternalServerError, H{
			"error": err,
		})
	}
	return c.JSON(http.StatusOK, H{
		"response": "Deleted",
	})
}

// retrieveSponsors - Retrieve a sponsor from the database
func retrieveSponsors(token string, tierString string) ([]*Sponsor, error) {
	var results []*Sponsor

	filter := bson.D{{}}
	if tierString != "" {
		tier, err := strconv.Atoi(tierString)
		if err != nil {
			return results, err
		}
		filter = bson.D{{Key: "tier", Value: tier}}
	}
	curr, err := sponsorColl.Find(context.TODO(), filter, options.Find())
	// decode result into sponsor array
	if err == nil {
		for curr.Next(context.TODO()) {
			var elem Sponsor
			curr.Decode(&elem)
			results = append(results, &elem)
		}
	}
	return results, err
}
