/*
  Social
  --
  This module handles social related API requests.

  It sends a JSON array of objects containing information about the various
  social platforms of CSESoc, read from a static file.
*/

package social

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	. "csesoc.unsw.edu.au/m/v2/server"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/labstack/echo/v4"
)

// Social - struct to contain social links data
type Social struct {
	Title  string `json:"title" validate:"required"`
	Link   string `json:"link" validate:"required,url"`
	Source string `json:"src"`
}

var socialColl *mongo.Collection

////////
// SETUP
////////

// Setup - setup the collection to be used for social links
func Setup(client *mongo.Client) {
	socialColl = client.Database("csesoc").Collection("socials")

	// Creating unique index for sponsor name
	opt := options.Index()
	opt.SetUnique(true)
	index := mongo.IndexModel{
		Keys:    bson.M{"title": 1},
		Options: opt,
	}
	if _, err := socialColl.Indexes().CreateOne(context.Background(), index); err != nil {
		log.Fatal("Could not create index: ", err)
	}

	// Fetching faq list
	socials, err := readSocialJSON()
	if err != nil {
		log.Fatal("Could not retrive social links from JSON")
	}

	for _, social := range socials {
		if _, err := socialColl.InsertOne(context.TODO(), social); err != nil {
			log.Printf("Could not insert social link " + social.Title + " " + err.Error())
		}
	}
}

///////////
// HANDLERS
///////////

// HandleGet godoc
// @Summary Return all social media links
// @Tags social
// @Success 200 {array} Social
// @Failure 500 "Service unavailable"
// @Header 500 {string} error "Missing fields"
// @Failure 503 "Service unavailable"
// @Header 503 {string} error "Unable to retrieve social media links"
// @Router /social [get]
func HandleGet(c echo.Context) error {
	socials, err := retrieveSocials()

	if err != nil {
		return c.JSON(http.StatusServiceUnavailable, H{
			"error": "Unable to retrieve social links from database",
		})
	}

	// Validate structss
	for _, social := range socials {
		if err := c.Validate(social); err != nil {
			return c.JSON(http.StatusInternalServerError, H{
				"error": fmt.Sprintf("Missing fields on: %v", social),
			})
		}
	}

	return c.JSON(http.StatusOK, socials)
}

//////////
// HELPERS
//////////

func retrieveSocials() ([]*Social, error) {
	var results []*Social

	curr, err := socialColl.Find(context.TODO(), bson.M{})
	// decode result into social links array
	if err == nil {
		for curr.Next(context.TODO()) {
			var elem Social
			curr.Decode(&elem)
			results = append(results, &elem)
		}
	}
	return results, err
}

func readSocialJSON() ([]Social, error) {
	byteValue, err := ReadJSON("social")
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	var socials []Social
	json.Unmarshal(byteValue, &socials)

	return socials, nil
}
