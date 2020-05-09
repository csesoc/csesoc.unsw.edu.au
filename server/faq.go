package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/labstack/echo/v4"
)

// Faq - struct to store faq pairs
type faq struct {
	Question string `json:"Question"`
	Answer   string `json:"Answer"`
}

// GetFaq - Returns all faq questions and answers pairs
func GetFaq() echo.HandlerFunc {
	return func(c echo.Context) error {
		result := getFaq()
		return c.JSON(http.StatusOK, H{
			"faq": result,
		})
	}
}

// GetFaq - returns a list of questions and answers from a json file in /static
func getFaq() []faq {
	abspath, _ := filepath.Abs("static/faq.json")
	jsonFile, err := os.Open(abspath)

	if err != nil {
		log.Fatal(err)
	}

	byteValue, _ := ioutil.ReadAll(jsonFile)
	var faqs []faq
	json.Unmarshal(byteValue, &faqs)

	defer jsonFile.Close()
	return faqs
}

/*
///// IF USING THE DATABASE TO RETRIVE FAQ//////
// GetFaq - returns a jsonfile of the faq
func GetFaq(collection *mongo.Collection) echo.HandlerFunc {
	return func(c echo.Context) error {
		result := getFaq(collection)
		return c.JSON(http.StatusOK, H{
			"faq": result,
		})
	}
}

// getFaq - returns a list of questions and answers from the database
func getFaq(collection *mongo.Collection) []*Faq {
	var results []*Faq
	findOptions := options.Find()
	// TODO: Get rid of limit of fix limit
	findOptions.SetLimit(100)

	// finding all Q&A
	// note: collection is already just the faq,
	// so we are finding all information in the collection
	cur, err := collection.Find(context.TODO(), bson.D{{}}, findOptions)
	if err != nil {
		log.Fatal(err)
	}

	// decoding each element in 'findOptions'
	for cur.Next(context.TODO()) {

		// create a value into which the single document can be decoded
		var elem Faq
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}

		results = append(results, &elem)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	// Close the cursor once finished
	cur.Close(context.TODO())

	return results
}
*/
