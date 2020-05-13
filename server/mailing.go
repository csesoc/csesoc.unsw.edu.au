package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/context"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/gmail/v1"
)

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

// InitMailingClient initialises a session with the Gmail API and stores it in a global variable
func InitMailingClient() error {
	b, err := ioutil.ReadFile("credentials.json")
	if err != nil {
		return fmt.Errorf("Unable to read client secret file: %v", err)
	}

	// If modifying these scopes, delete your previously saved token.json.
	config, err := google.ConfigFromJSON(b, gmail.GmailReadonlyScope)
	if err != nil {
		return fmt.Errorf("Unable to parse client secret file to config: %v", err)
	}
	client, err = getClient(config)

	return nil
}

// SendEmail composes the email to send
func SendEmail(name string, email string, message string, targetEmail string) error {
	srv, err := gmail.New(client)
	if err != nil {
		return fmt.Errorf("Unable to retrieve Gmail client: %v", err)
	}

	var msg gmail.Message

	temp := []byte("From: '" + name + "' " + "<" + email + ">" + "\r\n" +
		"To: " + targetEmail + "\r\n" +
		"Subject: Enquiry from CSESoc Website\r\n" +
		"\r\n" + message)

	msg.Raw = base64.StdEncoding.EncodeToString(temp)
	msg.Raw = strings.Replace(msg.Raw, "/", "_", -1)
	msg.Raw = strings.Replace(msg.Raw, "+", "-", -1)
	msg.Raw = strings.Replace(msg.Raw, "=", "", -1)

	_, err = srv.Users.Messages.Send(name, &msg).Do()
	if err != nil {
		return fmt.Errorf("Unable to send: %v", err)
	}

	return nil
}
