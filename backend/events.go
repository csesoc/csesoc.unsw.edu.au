package main

import (
	"time"
	"fmt"
	"encoding/json"
	// "os"
	"log"
	"io/ioutil"
	"net/http"
	. "csesoc.unsw.edu.au/m/v2/server"
)

type FbResponse struct {
	Data  []FbEvent `json:"data"`
	Error FbError `json:"error"` 
}

type FbEvent struct {
	Description string `json:"string"`
	Name string `json:"name"`
	Start string `json:"start_time"`
	End string `json:"end_time"`
	Id string `json:"id"`
	Place FbLocation `json:"place"`
}

type FbLocation struct {
	Name string `json:"name"`
}

type FbError struct {
	ErrorType int `json:"type"`
	Message string `json:"message"`
}

func callInterval(d time.Duration, f func()) {
	for range time.Tick(d) {
		f()
	}
}


func getEvents() {
	resp, err := http.Get(
						fmt.Sprintf("%s?access_token=%s&since=%d", 
									FB_EVENT_PATH,
									FB_TOKEN,
									time.Now().Unix()))
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	
	if resp.StatusCode == http.StatusOK {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			// error
		} 
		var result FbResponse
		json.Unmarshal(body, &result)
		
		if result.Error != (FbError{}) {
			// error
		}

		for _, element := range result.Data {
			fmt.Println("Event found: ", element.Name)
		}

		
		
	}
}
