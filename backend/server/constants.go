package utility

const DEVELOPMENT bool = true
const BASE_URL = "http://localhost:1323/"
const SPONSOR_URL = "api/v1/sponsors"
const MAILING_URL = "api/v1/mailing"
const SOCIAL_URL = "api/v1/social"
const FAQ_URL = "api/v1/faq"
// JWT used for testing
const AUTH_TOKEN = "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhZG1pbiI6dHJ1ZSwiZXhwIjoxNTk4Mzg5MDg2LCJ6SUQiOiJ6NTEyMzQ1NiJ9.a8869YcfgBZCQcWUFuomF4Fqx9nJzT-rWpP2yRVsBWg"
// Constants for accessing FB API
const FB_EVENT_PATH = "https://graph.facebook.com/v7.0/110742340691435/events"
const FB_TOKEN = ""
var JWT_SECRET = []byte("temp_secret_until_proper_secrets_are_implemented")
