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
	SocialID int    `json:"id" validate:"required"`
	Title    string `json:"title" validate:"required"`
	Link     string `json:"link" validate:"required"`
	Source   string `json:"src"`
}

///////////
// HANDLERS
///////////

// GetSocial godoc
// @Summary Return all social media links
// @Tags social
// @Success 200 {array} Social
// @Failure 503 "Service unavailable"
// @Header 503 {string} error "Unable to retrieve social media links"
// @Router /social [get]
func GetSocial(c echo.Context) error {
	socials, err := readSocialJSON()

	if err != nil {
		return c.JSON(http.StatusServiceUnavailable, H{
			"error": "Unable to retrieve socials",
		})
	}

	// Validate struct
	if err := c.Validate(socials); err != nil {
		return c.JSON(http.StatusInternalServerError, H{
			"error": "Missing fields on one or more social link",
		})
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
