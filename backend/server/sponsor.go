package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

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
	Detail string `json:"detail" validate:"required"`
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

	sponsors, err := retriveSponsorsJSON()

	if err != nil {
		log.Fatal("Could not retrive sponsors from JSON")
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
func NewSponsor() echo.HandlerFunc {
	return func(c echo.Context) error {
		tier, err := strconv.Atoi(c.FormValue("tier"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, H{
				"error": "Tier has to be a number",
			})
		}
		sponsor := Sponsor{
			Name:   c.FormValue("name"),
			Logo:   c.FormValue("logo"),
			Tier:   tier,
			Detail: c.FormValue("detail"),
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

}

// GetSponsor - find entry for a specific sponsor.
func GetSponsor() echo.HandlerFunc {
	return func(c echo.Context) error {
		var result Sponsor
		// token := c.FormValue("token")
		filter := bson.D{{Key: "name", Value: c.FormValue("name")}}
		if err := sponsorColl.FindOne(context.TODO(), filter).Decode(&result); err != nil {
			return c.JSON(http.StatusNotFound, H{
				"response": "No such sponsor.",
			})
		}
		return c.JSON(http.StatusOK, result)
	}
}

// GetSponsors - gives a list of sponsors stored.
func GetSponsors() echo.HandlerFunc {
	return func(c echo.Context) error {
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
}

// DeleteSponsor - Delete a sponsor
func DeleteSponsor() echo.HandlerFunc {
	return func(c echo.Context) error {
		// token := c.FormValue("token")
		filter := bson.D{{Key: "name", Value: c.FormValue("name")}}
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

// helper function

func retriveSponsorsJSON() ([]Sponsor, error) {
	abspath, _ := filepath.Abs("static/sponsor.json")
	jsonFile, err := os.Open(abspath)

	if err != nil {
		return nil, fmt.Errorf("Cound not open file faq.json: %v", err)
	}

	byteValue, _ := ioutil.ReadAll(jsonFile)
	var sponsors []Sponsor
	json.Unmarshal(byteValue, &sponsors)

	defer jsonFile.Close()
	return sponsors, nil
}
