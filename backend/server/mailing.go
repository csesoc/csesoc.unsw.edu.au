package main

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/mailjet/mailjet-apiv3-go"
)

// Enquiry - struct to contain email enquiry data
type Enquiry struct {
	Name  string `validate:"required"`
	Email string `validate:"required,email"`
	Body  string `validate:"required"`
}

// Feedback - struct to contain feedback message data
// name is not required, email must be valid (or empty) and body is required.
type Feedback struct {
	Name  string
	Email string `validate:"omitempty,email"`
	Body  string `validate:"required"`
}

// Message bundles
var feedbackBundle []Feedback
var infoBundle []Enquiry
var infoEmail string = "info@csesoc.org.au"
var sponsorshipBundle []Enquiry
var sponsorshipEmail string = "sponsorship@csesoc.org.au"

// Mailjet session variables
var publicKey string = "8afb96baef07230483a2a5ceca97d55d"
var secretKey string = "424ad90f25487e6be369a1cbb2a34694"
var mailjetClient *mailjet.Client

// InitMailClient initialises a session with the Mailjet API and stores it in a global variable
func InitMailClient() {
	mailjetClient = mailjet.NewMailjetClient(publicKey, secretKey)

	// Start mailing timers
	go enquiryMailingTimer()
	go feedbackMailingTimer()
}

///////////
// HANDLERS
///////////

// HandleEnquiry by forwarding emails to relevant inboxes
func HandleEnquiry(targetEmail string) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Extract fields from form
		enquiry := Enquiry{
			Name:  c.FormValue("name"),
			Email: c.FormValue("email"),
			Body:  c.FormValue("body"),
		}

		// Validate struct
		if err := c.Validate(enquiry); err != nil {
			return c.JSON(http.StatusBadRequest, H{
				"error": err,
			})
		}

		// Add to bundle
		switch targetEmail {
		case sponsorshipEmail:
			sponsorshipBundle = append(sponsorshipBundle, enquiry)
		case infoEmail:
			infoBundle = append(infoBundle, enquiry)
		}

		return c.JSON(http.StatusAccepted, H{})
	}
}

// HandleFeedback - validates feedback input and forwards as an email to the revelant inbox
func HandleFeedback() echo.HandlerFunc {
	return func(c echo.Context) error {
		// Extract fields from form
		feedback := Feedback{
			Name:  c.FormValue("name"),
			Email: c.FormValue("email"),
			Body:  c.FormValue("body"),
		}

		// Validate struct
		if err := c.Validate(feedback); err != nil {
			return c.JSON(http.StatusBadRequest, H{
				"error": err,
			})
		}

		// Add to bundle
		feedbackBundle = append(feedbackBundle, feedback)

		return c.JSON(http.StatusAccepted, H{})
	}
}

/////////
// TIMERS
/////////

// This function is executed once in a subroutine and triggers every 15 minutes
func enquiryMailingTimer() {
	for {
		time.Sleep(15 * time.Minute)
		DispatchEnquiryBundles()
	}
}

// This function is executed once in a subroutine and triggers once every day
func feedbackMailingTimer() {
	for {
		time.Sleep(24 * time.Hour)
		DispatchFeedbackBundle()
	}
}

//////////////
// DISPATCHERS
//////////////

// DispatchEnquiryBundles - public trigger for dispatching enquiries
func DispatchEnquiryBundles() {
	if len(infoBundle) > 0 {
		sendEnquiryBundle(infoEmail, &infoBundle)
	}
	if len(sponsorshipBundle) > 0 {
		sendEnquiryBundle(sponsorshipEmail, &sponsorshipBundle)
	}
}

// DispatchFeedbackBundle - public trigger for dispatching feedbacks
func DispatchFeedbackBundle() {
	if len(feedbackBundle) > 0 {
		sendFeedbackBundle(infoEmail, &feedbackBundle)
	}
}

/////////////////
// BUNDLE SENDERS
/////////////////

func sendEnquiryBundle(targetEmail string, bundle *[]Enquiry) {
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
			Subject:  "Website enquiry bundle",
			HTMLPart: joinEnquiries(*bundle),
		},
	}

	// Send query
	messages := mailjet.MessagesV31{Info: payload}
	_, err := mailjetClient.SendMailV31(&messages)
	if err == nil {
		// Only dump the bundle if email was successfully sent
		*bundle = nil
	}
}

func sendFeedbackBundle(targetEmail string, bundle *[]Feedback) {
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
			Subject:  "Website feedback bundle",
			HTMLPart: joinFeedbacks(*bundle),
		},
	}

	// Send query
	messages := mailjet.MessagesV31{Info: payload}
	_, err := mailjetClient.SendMailV31(&messages)
	if err == nil {
		// Only dump the bundle if email was successfully sent
		*bundle = nil
	}
}

func joinEnquiries(bundle []Enquiry) string {
	var message string = ""

	for _, msg := range bundle {
		message += "<hr />"
		message += "<h3>" + "Enquiry from " + msg.Name + " &lt;" + msg.Email + "&gt;" + "</h3>"
		message += "<p>" + msg.Body + "</p>"
	}
	message += "<hr />"

	return message
}

func joinFeedbacks(bundle []Feedback) string {
	var message string = ""

	for _, msg := range bundle {
		message += "<hr />"
		message += "<p>" + msg.Body + "</p>"
		if len(msg.Name) != 0 || len(msg.Email) != 0 {
			message += "<i>" + "Feedback from "
			if len(msg.Name) != 0 {
				message += msg.Name
				if len(msg.Email) != 0 {
					message += " &lt;" + msg.Email + "&gt;"
				}
			} else if len(msg.Email) != 0 {
				message += "&lt;" + msg.Email + "&gt;"
			}
			message += "</i>"
		}
	}
	message += "<hr />"

	return message
}
