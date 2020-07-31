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

// GetFaq godoc
// @Summary Return all faq questions and answers pairs
// @Tags faq
// @Success 200 {array} Faq
// @Failure 503 "Service unavailable"
// @Header 503 {string} error "Unable to retrieve FAQs"
// @Router /faq [get]
func GetFaq(c echo.Context) error {
	faqs, err := readFaqJSON()

	if err != nil {
		return c.JSON(http.StatusServiceUnavailable, H{
			"error": "Unable to retrieve FAQs",
		})
	}

	// Validate struct
	if err := c.Validate(faqs); err != nil {
		return c.JSON(http.StatusInternalServerError, H{
			"error": "Missing questions and/or answer fields",
		})
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
