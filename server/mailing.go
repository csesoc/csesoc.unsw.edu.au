package main

import (
	"net/http"
	"regexp"

	"github.com/labstack/echo/v4"
)

// HandleEnquiry by forwarding emails to relevant inboxes
func HandleEnquiry(targetEmail string) echo.HandlerFunc {
	return func(c echo.Context) error {
		name := c.FormValue("name")
		email := c.FormValue("email")
		message := c.FormValue("message")

		// Email validation
		emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)

		if emailRegex.MatchString(email) && len(name) > 0 && len(message) > 0 {
			// TODO: Compose and send email
			return c.JSON(http.StatusOK, H{})
		} else {
			return c.JSON(http.StatusBadRequest, H{})
		}
	}
}
