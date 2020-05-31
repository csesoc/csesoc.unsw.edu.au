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
var generalBundle []Enquiry
var sponsorshipBundle []Enquiry
var feedbackBundle []Feedback

// Email addresses
var infoEmail string = "info@csesoc.org.au"
var sponsorshipEmail string = "sponsorship@csesoc.org.au"

// Mailjet session variables
var publicKey string = "8afb96baef07230483a2a5ceca97d55d"
var secretKey string = "424ad90f25487e6be369a1cbb2a34694"
var mailjetClient *mailjet.Client

// MailingSetup initialises a session with the Mailjet API and stores it in a global variable
func MailingSetup() {
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

// HandleGeneralMessage godoc
// @Summary Handle a general enquiry by adding it to a dispatch bundle
// @Tags mailing
// @Param name formData string true "Name"
// @Param email formData string true "Email"
// @Param body formData string true "Message body"
// @Success 202 "Enquiry added to dispatch bundle"
// @Failure 400 "Invalid form"
// @Router /mailing/general [post]
func HandleGeneralMessage(c echo.Context) error {
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
	generalBundle = append(generalBundle, enquiry)

	return c.JSON(http.StatusAccepted, H{})
}

// HandleSponsorshipMessage godoc
// @Summary Handle a sponsorship enquiry by adding it to a dispatch bundle
// @Tags mailing
// @Param name formData string true "Name"
// @Param email formData string true "Email"
// @Param body formData string true "Message body"
// @Success 202 "Enquiry added to dispatch bundle"
// @Failure 400 "Invalid form"
// @Router /mailing/sponsorship [post]
func HandleSponsorshipMessage(c echo.Context) error {
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
	sponsorshipBundle = append(sponsorshipBundle, enquiry)

	return c.JSON(http.StatusAccepted, H{})
}

// HandleFeedbackMessage godoc
// @Summary Handle a feedback by adding it to a dispatch bundle
// @Tags mailing
// @Param name formData string false "Name"
// @Param email formData string false "Email"
// @Param body formData string true "Message body"
// @Success 202 "Feedback added to dispatch bundle"
// @Failure 400 "Invalid form"
// @Router /mailing/feedback [post]
func HandleFeedbackMessage(c echo.Context) error {
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
	if len(generalBundle) > 0 {
		if sendEmail(infoEmail, "Website info enquiry bundle", joinEnquiries(generalBundle)) {
			// If sent successfully, clear bundle
			generalBundle = nil
		}
	}
	if len(sponsorshipBundle) > 0 {
		if sendEmail(sponsorshipEmail, "Website sponsorship enquiry bundle", joinEnquiries(sponsorshipBundle)) {
			// If sent successfully, clear bundle
			sponsorshipBundle = nil
		}
	}
}

// DispatchFeedbackBundle - public trigger for dispatching feedbacks
func DispatchFeedbackBundle() {
	if len(feedbackBundle) > 0 {
		if sendEmail(infoEmail, "Website feedback bundle", joinFeedbacks(feedbackBundle)) {
			// If sent successfully, clear bundle
			feedbackBundle = nil
		}
	}
}

func sendEmail(targetEmail string, subject string, body string) bool {
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

/////////////////
// BUNDLE PARSERS
/////////////////

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
