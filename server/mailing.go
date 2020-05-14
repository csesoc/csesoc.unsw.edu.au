package main

import (
	"fmt"
	"net/http"
	"net/smtp"

	"github.com/labstack/echo/v4"
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
	password := "gmail-app-password"
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

// Gmail API approach
/*
// Current client session
var client *http.Client

// Retrieve a token, saves the token, then returns the generated client.
func getClient(config *oauth2.Config) (*http.Client, error) {
	// The file token.json stores the user's access and refresh tokens, and is
	// created automatically when the authorization flow completes for the first
	// time.
	tokFile := "token.json"
	tok, errTokenNotFound := tokenFromFile(tokFile)
	if errTokenNotFound != nil {
		tok, errCannotReachWeb := getTokenFromWeb(config)
		if errCannotReachWeb != nil {
			return nil, errCannotReachWeb
		}
		errCannotSave := saveToken(tokFile, tok)
		if errCannotSave != nil {
			return nil, errCannotSave
		}
	}
	return config.Client(context.Background(), tok), nil
}

// Request a token from the web, then returns the retrieved token.
func getTokenFromWeb(config *oauth2.Config) (*oauth2.Token, error) {
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	fmt.Printf("Go to the following link in your browser then type the "+
		"authorization code: \n%v\n", authURL)

	var authCode string
	if _, err := fmt.Scan(&authCode); err != nil {
		return nil, fmt.Errorf("Unable to read authorization code: %v", err)
	}

	tok, err := config.Exchange(context.TODO(), authCode)
	if err != nil {
		return nil, fmt.Errorf("Unable to retrieve token from web: %v", err)
	}
	return tok, nil
}

// Retrieves a token from a local file.
func tokenFromFile(file string) (*oauth2.Token, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	tok := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(tok)
	return tok, err
}

// Saves a token to a file path.
func saveToken(path string, token *oauth2.Token) error {
	fmt.Printf("Saving credential file to: %s\n", path)
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		return fmt.Errorf("Unable to cache oauth token: %v", err)
	}
	defer f.Close()
	json.NewEncoder(f).Encode(token)

	return nil
}

// Composes the email to send
func sendEmail(message Message, targetEmail string) error {
	srv, err := gmail.New(client)
	if err != nil {
		return fmt.Errorf("Unable to retrieve Gmail client: %v", err)
	}

	var msg gmail.Message

	temp := []byte("From: '" + message.Name + "' " + "<" + message.Email + ">" + "\r\n" +
		"To: " + targetEmail + "\r\n" +
		"Subject: Enquiry from CSESoc Website\r\n" +
		"\r\n" + message.Body)

	msg.Raw = base64.StdEncoding.EncodeToString(temp)
	msg.Raw = strings.Replace(msg.Raw, "/", "_", -1)
	msg.Raw = strings.Replace(msg.Raw, "+", "-", -1)
	msg.Raw = strings.Replace(msg.Raw, "=", "", -1)

	_, err = srv.Users.Messages.Send(message.Name, &msg).Do()
	if err != nil {
		return fmt.Errorf("Unable to send: %v", err)
	}

	return nil
}

// InitMailingClient initialises a session with the Gmail API and stores it in a global variable
func InitMailingClient() {
	b, err := ioutil.ReadFile("credentials.json")
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
	}

	// If modifying these scopes, delete your previously saved token.json.
	config, err := google.ConfigFromJSON(b, gmail.GmailReadonlyScope)
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}
	client, err = getClient(config)
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

		if err := sendEmail(message, targetEmail); err != nil {
			return c.JSON(http.StatusServiceUnavailable, H{
				"error": err,
			})
		}

		return c.JSON(http.StatusOK, H{})
	}
}
*/
