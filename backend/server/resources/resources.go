/*
  Resources
  --
  This module deals with CSE resource links. It creates a database collection
  and inserts the resource links (as JSON objects) read from a static file.

  The API handles directly interact with the database.
*/

package resources

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	. "csesoc.unsw.edu.au/m/v2/server"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var resourceColl *mongo.Collection

// Resource - struct for the list of resources that are displayed
type Resource struct {
	Title       string `json:"title" validate:"required"`
	Description string `json:"description" validate:"required"`
	Link        string `json:"link" validate:"required,url"`
	Source      string `json:"src" validate:"required"`
}

////////
// SETUP
////////

// Setup - Set up the resources collection
func Setup(client *mongo.Client) {
	resourceColl = client.Database("csesoc").Collection("resources")

	// Creating unique index for resource title
	opt := options.Index()
	opt.SetUnique(true)
	index := mongo.IndexModel{
		Keys:    bson.M{"title": 1},
		Options: opt,
	}
	if _, err := resourceColl.Indexes().CreateOne(context.Background(), index); err != nil {
		log.Fatal("Could not create index: ", err)
	}

	// Fetching resource list
	resources, err := readResourceJSON()
	if err != nil {
		log.Fatal("Unable to retrieve Resources from JSON")
	}

	for _, resource := range resources {
		if _, err := resourceColl.InsertOne(context.TODO(), resource); err != nil {
			log.Printf("Could not insert resource " + resource.Title + " " + err.Error())
		}
	}
}

///////////
// HANDLERS
///////////

// HandleGetPreview godoc
// @Summary Get a list of resources stored
// @Tags resources
// @Success 200 {array} Resource
// @Failure 500 "Internal server error"
// @Header 500 {string} error "Unable to retrieve resources from database"
// @Router /resources/preview [get]
func HandleGetPreview(c echo.Context) error {
	var results []*Resource

	// get database pointer
	curr, err := resourceColl.Find(context.TODO(), bson.D{{}}, options.Find())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, H{
			"error": "Unable to retrieve resources from database",
		})
	}

	// decode result into resource array
	for curr.Next(context.TODO()) {
		var elem Resource
		curr.Decode(&elem)
		results = append(results, &elem)
	}

	return c.JSON(http.StatusOK, results)
}

//////////
// HELPERS
//////////

func readResourceJSON() ([]Resource, error) {
	byteValue, err := ReadJSON("resource")
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	var resources []Resource
	json.Unmarshal(byteValue, &resources)

	return resources, nil
}
