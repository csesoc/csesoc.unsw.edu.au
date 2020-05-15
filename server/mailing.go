package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// Message - struct to contain email message data
type Message struct {
	Name  string `validate:"required"`
	Email string `validate:"required,email"`
	Body  string `validate:"required"`
}

// Feedback - struct to contain feedback message data
// name and email are not required
type Feedback struct {
	Name  string
	Email string `validate:"email"`
	Body  string `validate:"required"`
}

// HandleEnquiry by forwarding emails to relevant inboxes
func HandleEnquiry(targetEmail string) echo.HandlerFunc {
	return func(c echo.Context) error {
		message := Message{
			Name:  c.FormValue("name"),
			Email: c.FormValue("email"),
			Body:  c.FormValue("body"),
		}

		if err := c.Validate(message); err != nil {
			return c.JSON(http.StatusBadRequest, H{
				"error": err,
			})
		}

		// TODO: Compose and send email
		return c.JSON(http.StatusOK, H{})
	}
}

// FeedbackEnquiry - validates feedback input and forwards as an email to the revelant inbox
func FeedbackEnquiry(targetEmail string) echo.HandlerFunc {
	return func(c echo.Context) error {
		feedback := Feedback{
			Name:  c.FormValue("name"),
			Email: c.FormValue("email"),
			Body:  c.FormValue("body"),
		}

		if err := c.Validate(feedback); err != nil {
			return c.JSON(http.StatusBadRequest, H{
				"error": err,
			})
		}

		//TODO: Compose email and send
		return c.JSON(http.StatusOK, H{})
	}
}
