package social

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

// Social - struct to contain social links data
type Social struct {
	SocialID int    `json:"id"`
	Title    string `json:"title"`
	Link     string `json:"link"`
	Source   string `json:"src"`
}

// GetSocial godoc
// @Summary Return all social media links
// @Tags social
// @Success 200 {array} Social
// @Failure 503 "Service unavailable"
// @Header 503 {string} error "Unable to retrieve social media links"
// @Router /social [get]
func GetSocial(c echo.Context) error {
	results, err := retriveSocialJSON()

	if err != nil {
		return c.JSON(http.StatusServiceUnavailable, H{
			"error": "Unable to retrieve socials",
		})
	}

	return c.JSON(http.StatusOK, results)
}

// retriveFaqJSON - returns a list of questions and answers from a json file in /static
func retriveSocialJSON() ([]Social, error) {
	abspath, _ := filepath.Abs("static/social.json")
	jsonFile, err := os.Open(abspath)

	if err != nil {
		return nil, fmt.Errorf("Cound not open file faq.json: %v", err)
	}

	byteValue, _ := ioutil.ReadAll(jsonFile)
	var socials []Social
	json.Unmarshal(byteValue, &socials)

	defer jsonFile.Close()
	return socials, nil
}
