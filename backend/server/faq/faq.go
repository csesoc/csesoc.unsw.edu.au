/*
  FAQ
  --
  This module handles FAQ related API requests.
  It responds with a JSON array of question/answer pairs read from a static file.
*/

package faq

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

// Faq - struct to store faq pairs
type Faq struct {
	Question string `json:"question" validate:"required"`
	Answer   string `json:"answer" validate:"required"`
}

var faqColl *mongo.Collection

////////
// SETUP
////////

// Setup - setup the collection to be used for faq
func Setup(client *mongo.Client) {
	faqColl = client.Database("csesoc").Collection("faqs")

	// Creating unique index for sponsor name
	opt := options.Index()
	opt.SetUnique(true)
	index := mongo.IndexModel{
		Keys:    bson.M{"question": 1},
		Options: opt,
	}
	if _, err := faqColl.Indexes().CreateOne(context.Background(), index); err != nil {
		log.Fatal("Could not create index: ", err)
	}

	// Fetching faq list
	faqs, err := readFaqJSON()
	if err != nil {
		log.Fatal("Could not retrive Faqs from JSON")
	}

	// Try to update; insert if document is not found
	optUpsert := options.Update().SetUpsert(true)
	for _, faq := range faqs {
		filter := bson.M{"name": faq.Question}
		update := bson.M{"$set": faq}
		if _, err := faqColl.UpdateOne(context.TODO(), filter, update, optUpsert); err != nil {
			log.Printf("Could not insert faqs " + faq.Question + " " + err.Error())
		}
	}
}

///////////
// HANDLERS
///////////

// HandleGet godoc
// @Summary Return all faq questions and answers pairs
// @Tags faq
// @Success 200 {array} Faq
// @Failure 500 "Service unavailable"
// @Header 500 {string} error "Missing questions and/or answer fields"
// @Failure 503 "Service unavailable"
// @Header 503 {string} error "Unable to retrieve FAQs"
// @Router /faq [get]
func HandleGet(c echo.Context) error {
	faqs, err := retrieveFaqs()

	if err != nil {
		return c.JSON(http.StatusServiceUnavailable, H{
			"error": "Unable to retrieve FAQs from database",
		})
	}

	// Validate structs
	for _, faq := range faqs {
		if err := c.Validate(faq); err != nil {
			return c.JSON(http.StatusInternalServerError, H{
				"error": "Missing questions and/or answer fields",
			})
		}
	}

	return c.JSON(http.StatusOK, faqs)
}

//////////
// HELPERS
//////////

func retrieveFaqs() ([]*Faq, error) {
	var results []*Faq

	curr, err := faqColl.Find(context.TODO(), bson.M{})
	// decode result into faq array
	if err == nil {
		for curr.Next(context.TODO()) {
			var elem Faq
			curr.Decode(&elem)
			results = append(results, &elem)
		}
	}
	return results, err
}

func readFaqJSON() ([]Faq, error) {
	byteValue, err := ReadJSON("faq")
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	var faqs []Faq
	json.Unmarshal(byteValue, &faqs)

	return faqs, nil
}
