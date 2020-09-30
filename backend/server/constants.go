/*
  Constants
  --
  These constants can be dot imported into any module to get direct access to them.
*/

package utility

import "os"

const DEVELOPMENT bool = true
const BASE_URL = "http://localhost:1323/"
const SPONSOR_URL = "api/v1/sponsors"
const MAILING_URL = "api/v1/mailing"
const SOCIAL_URL = "api/v1/social"
const EVENTS_URL = "api/v1/events"
const FAQ_URL = "api/v1/faq"
const RESOURCES_URL = "api/v1/resources"

// JWT used for testing
var AUTH_TOKEN = "Bearer " + os.Getenv("TESTING_JWT")

var JWT_SECRET = []byte(os.Getenv("JWT_SECRET"))

// Constants for accessing FB API
const FB_API_PATH = "https://graph.facebook.com/v7.0"
const FB_EVENT_PATH = "/csesoc/events"
const FB_FETCH_INTERVAL = 120
