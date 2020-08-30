/*
  Events
  --
  This module takes care of grabbing event data from the Facebook Graph API
  and handling related API requests.

  Since the API returns an extensive amount of data there are a lot of
  required structs to parse and decompose the response payload.
*/

package events

import (
	"bytes"
	"encoding/json"
	"fmt"
	"sort"
	"time"
	"os"

	"io/ioutil"
	"net/http"
	"path/filepath"

	. "csesoc.unsw.edu.au/m/v2/server"

	"github.com/labstack/echo/v4"
	"github.com/relvacode/iso8601"
)

// FB response expects data (array of events), and paging, which we ignore.
// However, should FB provide an error, we capture it.
type (
	// FbResponse - Facebook API response container
	FbResponse struct {
		Data  []FbRespEvent `json:"data"`
		Cover FbRespCover   `json:"cover"`
		Error FbRespError   `json:"error"`
	}

	// FbRespEvent - struct to unmarshal event specifics
	FbRespEvent struct {
		Description string        `json:"description"`
		Name        string        `json:"name"`
		Start       string        `json:"start_time"`
		End         string        `json:"end_time"`
		EventTimes  []FbRespTimes `json:"event_times"`
		ID          string        `json:"id"`
		Place       FbRespPlace   `json:"place"`
	}

	// FbRespPlace - event location can come with added information, so we only take the name
	FbRespPlace struct {
		Name string `json:"name"`
	}

	// FbRespTimes - struct to deal with recurring events
	FbRespTimes struct {
		Start string `json:"start_time"`
		End   string `json:"end_time"`
	}

	// FbRespError - struct to unmarshal any error response
	FbRespError struct {
		ErrorType int    `json:"type"`
		Message   string `json:"message"`
	}

	// FbRespCover - struct to unmarshal the URI of the cover image
	FbRespCover struct {
		CoverURI string `json:"source"`
	}

	// MarshalledEvents - struct to pack up events with the last update time to be marshalled.
	MarshalledEvents struct {
		LastUpdate int64   `json:"updated"`
		Events     []Event `json:"events"`
	}

	// Event - struct to store an individual event with all the info we want
	Event struct {
		Name        string `json:"name"`
		Description string `json:"description"`
		Start       int64  `json:"start_time"`
		End         int64  `json:"end_time"`
		ID          string `json:"fb_event_id"`
		Place       string `json:"place"`
		CoverURL    string `json:"fb_cover_img"`
	}
)

///////////
// HANDLERS
///////////

// HandleGet godoc
// @Summary Get a list of upcoming events
// @Tags events
// @Success 200 {array} Event
// @Failure 500 "Internal server error"
// @Header 500 {string} error "Unable to retrieve events from file"
// @Router /events [get]
func HandleGet(c echo.Context) error {
	fp, err := filepath.Abs("static/events.json")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, H{
			"error": "Unable to retrieve events",
		})
	}
	return c.File(fp)
}

/////////
// TIMERS
/////////

// FetchTimer - sets up a ticker to fetch events at an interval
func FetchTimer() {
	saveEvents()
	time.Sleep(time.Duration(FB_FETCH_INTERVAL * time.Second))
}

//////////
// HELPERS
//////////

// Fetch events from FB
func fetchEvents(response *FbResponse) error {
	// Make a request to FB
	resp, err := http.Get(
		fmt.Sprintf(
			"%s%s?access_token=%s&since=%d",
			FB_API_PATH,
			FB_EVENT_PATH,
			os.Getenv("FB_TOKEN"),
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

// Fetch the cover image for an event from Facebook.
func fetchCoverImage(id string) (string, error) {
	resp, err := http.Get(
		fmt.Sprintf(
			"%s/%s?fields=cover&access_token=%s",
			FB_API_PATH,
			id,
			os.Getenv("FB_TOKEN"),
		),
	)
	if err != nil {
		return "", fmt.Errorf("There was an error making a request to FB")
	}

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("There was an error making a request to FB")
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("Something went wrong with parsing the FB request")
	}

	// Unmarshal the response body to pull the cover image from it
	var result FbResponse
	err = json.Unmarshal(body, &result)
	if err != nil {
		return "", fmt.Errorf("There was an issue parsing JSON")
	}
	if result.Error != (FbRespError{}) {
		return "", fmt.Errorf("Something went wrong with the FB request")
	}
	return result.Cover.CoverURI, nil
}

// Process the facebook event information before saving it to a new file.
func saveEvents() {
	var result FbResponse
	err := fetchEvents(&result)
	if err != nil {
		// do something
	}

	// Store processed events
	var processedEvents []Event

	for _, element := range result.Data {
		if len(element.EventTimes) != 0 {
			for _, occurrence := range element.EventTimes {
				cover, err := fetchCoverImage(element.ID)
				if err != nil {
					// do something
				}
				start, err := iso8601.ParseString(occurrence.Start)
				if err != nil {
					// do something
				}
				end, err := iso8601.ParseString(occurrence.End)
				if err != nil {
					// do something
				}

				processedEvents = append(processedEvents, Event{
					Name:        element.Name,
					Description: element.Description,
					Start:       start.Unix(),
					End:         end.Unix(),
					ID:          element.ID,
					CoverURL:    cover,
				})
			}
		} else {
			cover, err := fetchCoverImage(element.ID)
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

			processedEvents = append(processedEvents, Event{
				Name:        element.Name,
				Description: element.Description,
				Start:       start.Unix(),
				End:         end.Unix(),
				ID:          element.ID,
				Place:       element.Place.Name,
				CoverURL:    cover,
			})
		}
	}

	// Sort the parsed events by starting date
	sort.Slice(processedEvents, func(i, j int) bool {
		return processedEvents[i].Start < processedEvents[j].Start
	})

	buf := bytes.NewBuffer([]byte{})
	// Using json.NewEncoder allows for escaping ampersands.
	jsonEncoder := json.NewEncoder(buf)
	jsonEncoder.SetEscapeHTML(false)
	jsonEncoder.Encode(MarshalledEvents{
		LastUpdate: time.Now().Unix(),
		Events:     processedEvents,
	})
	fp, _ := filepath.Abs("static/events.json")
	// 644 = rw--w--w-
	err = ioutil.WriteFile(fp, buf.Bytes(), 0644)

	if err != nil {
		// error handling
	}
}
