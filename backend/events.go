package main

import (
	"time"
	"fmt"
	"encoding/json"
	// "os"
	// "log"
	"path/filepath"
	"io/ioutil"
	"net/http"
	. "csesoc.unsw.edu.au/m/v2/server"
	"github.com/relvacode/iso8601"
)


// FB response expects data (array of events), and paging, which we ignore
// However, should FB provide an error, we capture it.
type FbResponse struct {
	Data  []FbRespEvent `json:"data"`
	Cover FbCover `json:"cover"`
	Error FbRespError `json:"error"` 
}

// Unmarshal event specifics
type FbRespEvent struct {
	Description string `json:"description"`
	Name string `json:"name"`
	Start string `json:"start_time"`
	End string `json:"end_time"`
	Id string `json:"id"`
	Place FbRespPlace `json:"place"`
}

// Event location can come with added information, so we only take the name
type FbRespPlace struct {
	Name string `json:"name"`
}

// Unmarshall any error response
type FbRespError struct {
	ErrorType int `json:"type"`
	Message string `json:"message"`
}

// Get a cover image
type FbCover struct {
	CoverUri string `json:"source"`
}

type MarshalledEvents struct {
	LastUpdate int64 `json:"updated"`
	Events []Event `json:"events"`
}

type Event struct {
	Name string `json:"name"`
	Description string `json:"description"`
	Start int64 `json:"start_time"`
	End int64 `json:"end_time"`
	Id string `json:"fb_event_id"`
	CoverUrl string `json:"fb_cover_img"`
}

func callInterval(d time.Duration, f func()) {
	for range time.Tick(d) {
		f()
	}
}

func fetchEvents(response *FbResponse) (error) {
	// Make a request to FB
	resp, err := http.Get(
		fmt.Sprintf(
			"%s%s?access_token=%s&since=%d",
			FB_API_PATH, 
			FB_EVENT_PATH,
			FB_TOKEN,
			time.Now().Unix()),
		)
	if err != nil {
		return fmt.Errorf("There was an error making a request to FB")
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("There was an error making a request to FB")	
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("Something went wrong with parsing the FB request")
	} 

	json.Unmarshal(body, &response)

	// Check to make sure no error was captured.
	if (*response).Error != (FbRespError{}) {
		return fmt.Errorf("Something went wrong with the FB request")
	}
	return nil
}

func fetchCoverImage(id string) (string, error) {
		resp, err := http.Get(
			fmt.Sprintf(
				"%s/%s?fields=cover&access_token=%s",
				FB_API_PATH,
				id,
				FB_TOKEN,
			),
		)
		if err != nil {
			return "", fmt.Errorf("There was an error making a request to FB")
		}

		if resp.StatusCode != http.StatusOK {
			return "", fmt.Errorf("There was an error making a request to FB.")	
		}
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return "", fmt.Errorf("Something went wrong with parsing the FB request.")
		}

		// Unmarshal the response body to pull the cover image from it
		var result FbResponse
		err = json.Unmarshal(body, &result)
		if err != nil {
			return "", fmt.Errorf("There was an issue parsing JSON.")
		}
		if result.Error != (FbRespError{}) {
			return "", fmt.Errorf("Something went wrong with the FB request.")
		}
		return result.Cover.CoverUri, nil
}

func getEvents() {
		var result FbResponse
		err := fetchEvents(&result)
		if err != nil {
			// do something
		}

		// Store processed events
		var processedEvents []Event

		for _, element := range result.Data {
			cover, err := fetchCoverImage(element.Id)
			if err != nil {
				// do something
			}
			start, err := iso8601.ParseString(element.Start)
			if err != nil {
				// do something
			}
			end, err := iso8601.ParseString(element.End) 
			if err != nil {
				// do something
			}
			
			processedEvents = append(processedEvents, Event {
				Name: element.Name,
				Description: element.Description,
				Start: start.Unix(),
				End: end.Unix(),
				Id: element.Id,
				CoverUrl: cover,
			})
		}


		contents, _ := json.Marshal(MarshalledEvents {
										LastUpdate: time.Now().Unix(),
										Events: processedEvents,
									})
		fp, _ := filepath.Abs("static/events.json")
		// chmod a+rwx,u-x,g-wx,o-wx
		err = ioutil.WriteFile(fp, contents, 0644)

		if err != nil {
			// error handling
		}				

}
