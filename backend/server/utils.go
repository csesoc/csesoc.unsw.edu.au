/*
  Utils
  --
  This file contains general helper functions that are used in multiple modules.
  The categories of utilities are:
  - JSON
  - Testing
*/

package utility

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

///////
// JSON
///////

// H - interface for sending JSON
type H map[string]interface{}

// ReadJSON - returns JSONs from static files
func ReadJSON(name string) ([]byte, error) {
	var filename string = "static/" + name + ".json"
	abspath, _ := filepath.Abs(filename)
	jsonFile, err := os.Open(abspath)
	defer jsonFile.Close()

	if err != nil {
		return nil, fmt.Errorf("Could not open file: %v", err)
	}

	return ioutil.ReadAll(jsonFile)
}

//////////
// TESTING
//////////

// AssertStatus - wrapper function for testing the response status of HTTP requests
func AssertStatus(t *testing.T, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("Wrong status: got '%d', want '%d'", got, want)
	}
}

// AssertResponseBody - wrapper function for testing the response body of HTTP requests
func AssertResponseBody(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("Wrong response body, got '%s', want '%s'", got, want)
	}
}
