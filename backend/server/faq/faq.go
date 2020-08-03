/*
  FAQ
  --
  This module handles FAQ related API requests.
  It responds with a JSON array of question/answer pairs read from a static file.
*/

package faq

import (
	"encoding/json"
	"fmt"
	"net/http"

	. "csesoc.unsw.edu.au/m/v2/server"

	"github.com/labstack/echo/v4"
)

// Faq - struct to store faq pairs
type Faq struct {
	Question string `json:"question" validate:"required"`
	Answer   string `json:"answer" validate:"required"`
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
	faqs, err := readFaqJSON()

	if err != nil {
		return c.JSON(http.StatusServiceUnavailable, H{
			"error": "Unable to retrieve FAQs",
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

func readFaqJSON() ([]Faq, error) {
	byteValue, err := ReadJSON("faq")
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	var faqs []Faq
	json.Unmarshal(byteValue, &faqs)

	return faqs, nil
}
