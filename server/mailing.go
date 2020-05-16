package main

import (
	"fmt"
	"log"
	"net/http"
	"net/smtp"
	"syscall"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/ssh/terminal"
)

// Message - struct to contain email message data
type Message struct {
	Name  string `validate:"required"`
	Email string `validate:"required,email"`
	Body  string `validate:"required"`
}

// SMTP session variables
var host string
var auth smtp.Auth
var serverEmail string

// InitSMTPClient initialises a session with the Gmail API and stores it in a global variable
func InitSMTPClient() {
	serverEmail = "csesoc@csesoc.org.au"
	password := getPassword()
	host = "smtp.gmail.com:587"
	auth = smtp.PlainAuth("", serverEmail, password, "smtp.gmail.com")
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

		// Format message and targetEmail
		to := []string{targetEmail}
		msg := []byte(composeEmail(message, targetEmail))

		// Send mail to address
		if err := smtp.SendMail(host, auth, serverEmail, to, msg); err != nil {
			return c.JSON(http.StatusServiceUnavailable, H{
				"error": err,
			})
		}

		return c.JSON(http.StatusOK, H{})
	}
}

// Format Message to be of RFC 822-style
func composeEmail(message Message, targetEmail string) string {
	// Define header fields
	header := make(map[string]string)
	header["Resent-From"] = message.Email
	header["Reply-To"] = message.Email
	header["From"] = serverEmail
	header["To"] = targetEmail
	header["Subject"] = fmt.Sprintf("Enquiry from '%s' <%s>", message.Name, message.Email)

	// Stringify header
	headerMsg := ""
	for key, value := range header {
		headerMsg += fmt.Sprintf("%s: %s\r\n", key, value)
	}

	// Return concatenated header and body
	return headerMsg + "\r\n" + message.Body
}

func getPassword() string {
	fmt.Print("Enter Password: ")

	bytePassword, err := terminal.ReadPassword(int(syscall.Stdin))
	if err != nil {
		log.Fatal("Could not process password: ", err)
	}

	return string(bytePassword)
}
