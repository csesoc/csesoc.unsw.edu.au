package faq

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"

	. "csesoc.unsw.edu.au/m/v2/server"

	"github.com/labstack/echo/v4"
)

// Faq - struct to store faq pairs
type Faq struct {
	Question string `json:"question"`
	Answer   string `json:"answer"`
}

// GetFaq godoc
// @Summary Return all faq questions and answers pairs
// @Tags faq
// @Success 200 {array} Faq
// @Failure 503 "Service unavailable"
// @Header 503 {string} error "Unable to retrieve FAQs"
// @Router /faq [get]
func GetFaq(c echo.Context) error {
	results, err := retriveFaqJSON()

	if err != nil {
		return c.JSON(http.StatusServiceUnavailable, H{
			"error": "Unable to retrieve FAQs",
		})
	}

	return c.JSON(http.StatusOK, results)
}

// retriveFaqJSON - returns a list of questions and answers from a json file in /static
func retriveFaqJSON() ([]Faq, error) {
	abspath, _ := filepath.Abs("static/faq.json")
	jsonFile, err := os.Open(abspath)

	if err != nil {
		return nil, fmt.Errorf("Cound not open file faq.json: %v", err)
	}

	byteValue, _ := ioutil.ReadAll(jsonFile)
	var faqs []Faq
	json.Unmarshal(byteValue, &faqs)

	defer jsonFile.Close()
	return faqs, nil
}
