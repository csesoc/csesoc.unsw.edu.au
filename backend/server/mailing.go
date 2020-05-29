package main

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/mailjet/mailjet-apiv3-go"
)

type messageType int

const (
	// InfoType = 0
	InfoType messageType = iota
	// SponsorshipType = 1
	SponsorshipType
	// FeedbackType = 2
	FeedbackType
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
var sponsorshipBundle []Enquiry

// Email addresses
var infoEmail string = "info@csesoc.org.au"
var sponsorshipEmail string = "sponsorship@csesoc.org.au"

// Mailjet session variables
var publicKey string = "8afb96baef07230483a2a5ceca97d55d"
var secretKey string = "424ad90f25487e6be369a1cbb2a34694"
var mailjetClient *mailjet.Client

// InitMailClient initialises a session with the Mailjet API and stores it in a global variable
func InitMailClient() {
	if InDevelopment {
		infoEmail = "projects.website+info@csesoc.org.au"
		sponsorshipEmail = "projects.website+sponsorship@csesoc.org.au"
	}

	mailjetClient = mailjet.NewMailjetClient(publicKey, secretKey)

	// Start mailing timers
	go mailingTimer()
}

///////////
// HANDLERS
///////////

// HandleMessage by forwarding emails to relevant inboxes
func HandleMessage(mt messageType) echo.HandlerFunc {
	return func(c echo.Context) error {
		var enquiry Enquiry
		var feedback Feedback

		// Extract fields from form
		if mt == InfoType || mt == SponsorshipType {
			enquiry = Enquiry{
				Name:  c.FormValue("name"),
				Email: c.FormValue("email"),
				Body:  c.FormValue("body"),
			}
			if err := c.Validate(enquiry); err != nil {
				return c.JSON(http.StatusBadRequest, H{
					"error": err,
				})
			}
		} else if mt == FeedbackType {
			feedback = Feedback{
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
		}

		// Add to bundle
		switch mt {
		case InfoType:
			infoBundle = append(infoBundle, enquiry)
		case SponsorshipType:
			sponsorshipBundle = append(sponsorshipBundle, enquiry)
		case FeedbackType:
			feedbackBundle = append(feedbackBundle, feedback)
		}

		return c.JSON(http.StatusAccepted, H{})
	}
}

////////
// TIMER
////////

// This function is executed once in a subroutine and triggers every 15 minutes
func mailingTimer() {
	const minutesInDay int = 24 * 60
	const mailingInterval int = 15

	var intervalCounter int = 0
	for {
		time.Sleep(time.Duration(mailingInterval) * time.Minute)
		intervalCounter++
		DispatchEnquiryBundles()

		// Since feedback triggers every day, a different interval is used
		if intervalCounter == minutesInDay/mailingInterval {
			intervalCounter = 0
			DispatchFeedbackBundle()
		}
	}
}

//////////////
// DISPATCHERS
//////////////

// DispatchEnquiryBundles - public trigger for dispatching enquiries
func DispatchEnquiryBundles() {
	if len(infoBundle) > 0 {
		if sendBundle(infoEmail, "Website info enquiry bundle", joinEnquiries(infoBundle)) {
			// If sent successfully, clear bundle
			infoBundle = nil
		}
	}
	if len(sponsorshipBundle) > 0 {
		if sendBundle(sponsorshipEmail, "Website sponsorship enquiry bundle", joinEnquiries(sponsorshipBundle)) {
			// If sent successfully, clear bundle
			sponsorshipBundle = nil
		}
	}
}

// DispatchFeedbackBundle - public trigger for dispatching feedbacks
func DispatchFeedbackBundle() {
	if len(feedbackBundle) > 0 {
		if sendBundle(infoEmail, "Website feedback bundle", joinFeedbacks(feedbackBundle)) {
			// If sent successfully, clear bundle
			feedbackBundle = nil
		}
	}
}

/////////////////
// BUNDLE SENDERS
/////////////////

func sendBundle(targetEmail string, subject string, body string) bool {
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
			Subject:  subject,
			HTMLPart: body,
		},
	}

	// Send query
	messages := mailjet.MessagesV31{Info: payload}
	_, err := mailjetClient.SendMailV31(&messages)
	return err == nil
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
