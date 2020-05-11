package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Sponsor - struct to contain sponsor data
type Sponsor struct {
	Name   string `validate:"required"`
	Logo   string `validate:"required,url"`
	Tier   int    `validate:"required,numeric,eq=10|eq=100|eq=1000"`
	Expiry int64  // check because if this will be null most of the time then what would this number be?
}

var sponsorColl *mongo.Collection

/* Setup */

// SponsorSetup - Setup the collection to be used for sponsors
func SponsorSetup(client *mongo.Client) {
	sponsorColl = client.Database("csesoc").Collection("sponsors")
	opt := options.Index()
	opt.SetUnique(true)
	opt.SetCollation(&options.Collation{
		Locale:    "en",
		CaseLevel: true,
	})
	index := mongo.IndexModel{
		Keys:    bson.M{"name": 1},
		Options: opt,
	}
	if _, err := sponsorColl.Indexes().CreateOne(context.Background(), index); err != nil {
		log.Fatal("Could not create index: ", err)
	}

	resp, err := http.Get("https://gistcdn.githack.com/esyw/4e35cd5fe73fa024020e67855ca733fb/raw/e85c9ae58a6323a4214ffa4ad89b0a5ebe404e31/sponsors.json")
	if err != nil {
		log.Fatal("Could not get sponsor list: ", err)
	}
	defer resp.Body.Close()

	var sponsors []Sponsor
	if err = json.NewDecoder(resp.Body).Decode(&sponsors); err != nil {
		log.Printf("Could not convert JSON response to Sponsors")
	}

	fmt.Println(sponsors)
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
				"error": err,
			})
		}
		expiryTime, _ := time.Parse(time.RFC3339, c.FormValue("expiry"))
		sponsor := Sponsor{
			Name:   c.FormValue("name"),
			Logo:   c.FormValue("logo"),
			Tier:   tier,
			Expiry: expiryTime.Unix(),
		}
		if err := c.Validate(sponsor); err != nil {
			return c.JSON(http.StatusBadRequest, H{
				"error": err,
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
		name := strings.ToLower(c.FormValue("name"))
		filter := bson.D{{Key: "name", Value: name}}
		if err := sponsorColl.FindOne(context.TODO(), filter).Decode(&result); err != nil {
			return c.JSON(http.StatusNotFound, H{
				"response": "No such sponsor.",
			})
		}
		return c.JSON(http.StatusOK, H{
			"sponsor": result,
		})
	}
}

// GetSponsors - gives a list of sponsors stored.
func GetSponsors() echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.FormValue("token")
		tier := c.FormValue("tier")
		results, err := getSponsors(token, tier)
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

// getSponsors - Retrieve a sponsor from the database
func getSponsors(token string, tier string) ([]*Sponsor, error) {
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
