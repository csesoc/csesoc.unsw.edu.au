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
// name is not required, email must be valid and body is required.
type Feedback struct {
	Name  string
	Email string `validate:"email"`
	Body  string `validate:"required"`
}

// FeedbackNoEmail - struct to contain feedback message when there is no email given
// name, email are not required, body is.
type FeedbackNoEmail struct {
	Name  string
	Email string
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
		name := c.FormValue("name")
		email := c.FormValue("email")
		body := c.FormValue("body")

		// if there is no email, use the struct where emails are not required (FeedbackNoEmail)
		if email == "" {
			feedback := FeedbackNoEmail{
				Name:  name,
				Email: email,
				Body:  body,
			}

			if err := c.Validate(feedback); err != nil {
				return c.JSON(http.StatusBadRequest, H{
					"error": err,
				})
			}

			// if there is an email, check that the email is valid by using the Feedback Struct
		} else {
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

		}

		//TODO: Compose email and send
		return c.JSON(http.StatusOK, H{})
	}
}
