package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mailjet/mailjet-apiv3-go"
)

// Message - struct to contain email message data
type Message struct {
	Name  string `validate:"required"`
	Email string `validate:"required,email"`
	Body  string `validate:"required"`
}

// Mailjet session variables
var publicKey string = "MJ_APIKEY_PUBLIC"
var secretKey string = "MJ_APIKEY_PRIVATE"
var mailjetClient *mailjet.Client

// InitMailClient initialises a session with the Mailjet API and stores it in a global variable
func InitMailClient() {
	mailjetClient = mailjet.NewMailjetClient(publicKey, secretKey)
}

// HandleEnquiry by forwarding emails to relevant inboxes
func HandleEnquiry(targetEmail string) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Extract fields from form
		message := Message{
			Name:  c.FormValue("name"),
			Email: c.FormValue("email"),
			Body:  c.FormValue("body"),
		}

		// Validate struct
		if err := c.Validate(message); err != nil {
			return c.JSON(http.StatusBadRequest, H{
				"error": err,
			})
		}

		// Format message payload
		payload := []mailjet.InfoMessagesV31{
			mailjet.InfoMessagesV31{
				From: &mailjet.RecipientV31{
					Email: "projects.website@csesoc.org.au",
					Name:  "CSESoc Website",
				},
				To: &mailjet.RecipientsV31{
					mailjet.RecipientV31{
						Email: targetEmail,
					},
				},
				Subject:  fmt.Sprintf("Enquiry from '%s' <%s>", message.Name, message.Email),
				TextPart: message.Body,
			},
		}

		// Send query
		messages := mailjet.MessagesV31{Info: payload}
		_, err := mailjetClient.SendMailV31(&messages)
		if err != nil {
			return c.JSON(http.StatusServiceUnavailable, H{
				"error": err,
			})
		}

		return c.JSON(http.StatusOK, H{})
	}
}
