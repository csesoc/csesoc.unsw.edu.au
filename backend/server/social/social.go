/*
  Social
  --
  This module handles social related API requests.

  It sends a JSON array of objects containing information about the various
  social platforms of CSESoc, read from a static file.
*/

package social

import (
	"encoding/json"
	"fmt"
	"net/http"

	. "csesoc.unsw.edu.au/m/v2/server"

	"github.com/labstack/echo/v4"
)

// Social - struct to contain social links data
type Social struct {
	SocialID int    `json:"id" validate:"min=0"`
	Title    string `json:"title" validate:"required"`
	Link     string `json:"link" validate:"required,url"`
	Source   string `json:"src"`
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
	socials, err := readSocialJSON()

	if err != nil {
		return c.JSON(http.StatusServiceUnavailable, H{
			"error": "Unable to retrieve socials",
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

func readSocialJSON() ([]Social, error) {
	byteValue, err := ReadJSON("social")
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	var socials []Social
	json.Unmarshal(byteValue, &socials)

	return socials, nil
}
