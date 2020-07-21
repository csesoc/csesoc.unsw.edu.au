package sponsor

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

	. "csesoc.unsw.edu.au/m/v2/server"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Sponsor - struct to contain sponsor data
type Sponsor struct {
	Name   string `json:"name" validate:"required"`
	Logo   string `json:"logo" validate:"required"`
	Tier   int    `json:"tier" validate:"required,numeric,eq=0|eq=1|eq=2"`
	Detail string `json:"detail" validate:"required"`
	Url    string `json:"url" validate:"required,url"`
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
	sponsors, err := retriveSponsorsJSON()
	if err != nil {
		log.Fatal("Could not retrive sponsors from JSON")
	}

	optUpsert := options.Update().SetUpsert(true)
	for _, sponsor := range sponsors {
		filter := bson.M{"name": sponsor.Name}
		update := bson.M{"$set": sponsor}
		if _, err := sponsorColl.UpdateOne(context.TODO(), filter, update, optUpsert); err != nil {
			log.Printf("Could not insert sponsor " + sponsor.Name + " " + err.Error())
		}
	}
}

///////////
// HANDLERS
///////////

// NewSponsor godoc
// @Summary Add a new sponsor
// @Tags sponsors
// @accept Content-Type application/x-www-form-urlencoded
// @Param Authorization header string true "Bearer <token>"
// @Param name formData string true "Name"
// @Param logo formData string true "Logo in base64"
// @Param tier formData integer true "Valid tier" mininum(0) maxinum(2)
// @Param detail formData string true "Detail"
// @Success 201 "Created"
// @Header 201 {string} response "Sponsor added"
// @Failure 400 "Bad request"
// @Header 400 {string} error "Invalid form"
// @Failure 409 "Conflict"
// @Header 409 {string} error "Sponsor already exists on database"
// @Router /sponsors [post]
// @Security BearerAuthKey
func NewSponsor(c echo.Context) error {
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
		Url:    c.FormValue("url"),
	}

	// Validate the struct with golang validator package
	if err := c.Validate(sponsor); err != nil {
		return c.JSON(http.StatusBadRequest, H{
			"error": "Invalid form",
		})
	}

	if _, err := sponsorColl.InsertOne(context.TODO(), sponsor); err != nil {
		return c.JSON(http.StatusConflict, H{
			"error": "Sponsor already exists on database",
		})
	}

	return c.JSON(http.StatusCreated, H{
		"response": "Sponsor added",
	})
}

// GetSponsor godoc
// @Summary Find entry for a specific sponsor
// @Tags sponsors
// @Param name path string true "Sponsor name"
// @Success 200 {object} Sponsor
// @Failure 404 "Not found"
// @Header 404 {string} error "No such sponsor"
// @Router /sponsors/{name} [get]
func GetSponsor(c echo.Context) error {
	var result Sponsor
	filter := bson.D{{Key: "name", Value: c.Param("name")}}
	if err := sponsorColl.FindOne(context.TODO(), filter).Decode(&result); err != nil {
		return c.JSON(http.StatusNotFound, H{
			"error": "No such sponsor",
		})
	}
	return c.JSON(http.StatusOK, result)
}

// GetSponsors godoc
// @Summary Get a list of sponsors stored
// @Tags sponsors
// @Param tier query integer false "Valid sponsor tier, 0-2 inclusive" mininum(0) maxinum(2)
// @Success 200 {array} Sponsor
// @Failure 500 "Internal server error"
// @Header 500 {string} error "Unable to retrieve sponsors from database"
// @Router /sponsors [get]
func GetSponsors(c echo.Context) error {
	tier := c.QueryParam("tier")
	results, err := retrieveSponsors(tier)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, H{
			"error": "Unable to retrieve sponsors from database",
		})
	}
	return c.JSON(http.StatusOK, results)
}

// DeleteSponsor godoc
// @Summary Delete a sponsor
// @Tags sponsors
// @Param Authorization header string true "Bearer <token>"
// @Param name path string true "Sponsor name"
// @Success 204 "No content"
// @Header 204 {string} response "Sponsor deleted"
// @Failure 500 "Internal server error"
// @Header 500 {string} error "Unable to delete sponsor from database"
// @Router /sponsors/{name} [delete]
// @Security BearerAuthKey
func DeleteSponsor(c echo.Context) error {
	filter := bson.D{{Key: "name", Value: c.Param("name")}}
	if _, err := sponsorColl.DeleteOne(context.TODO(), filter); err != nil {
		return c.JSON(http.StatusInternalServerError, H{
			"error": "Unable to delete sponsor from database",
		})
	}
	return c.JSON(http.StatusNoContent, H{
		"response": "Sponsor deleted",
	})
}

//////////
// HELPERS
//////////

// retrieveSponsors - Retrieve a sponsor from the database
func retrieveSponsors(tierString string) ([]*Sponsor, error) {
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
