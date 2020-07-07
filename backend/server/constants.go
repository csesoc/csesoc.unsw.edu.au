package utility

const DEVELOPMENT bool = true
const BASE_URL = "http://localhost:1323/"
const SPONSOR_URL = "api/v1/sponsors"
const MAILING_URL = "api/v1/mailing"
const SOCIAL_URL = "api/v1/social"
const EVENTS_URL = "api/v1/social"
const FAQ_URL = "api/v1/faq"
// JWT used for testing
const AUTH_TOKEN = "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhZG1pbiI6dHJ1ZSwiZXhwIjoxNTk4Mzg5MDg2LCJ6SUQiOiJ6NTEyMzQ1NiJ9.a8869YcfgBZCQcWUFuomF4Fqx9nJzT-rWpP2yRVsBWg"
// Constants for accessing FB API
const FB_API_PATH = "https://graph.facebook.com/v7.0"
const FB_EVENT_PATH = "/110742340691435/events"
const FB_TOKEN = "EAAllxWUTDEABAHJDqpCepgkw3m5BDHWReP8QCseKZBiCeqycR2FDJQ50qZAIxZBAIglbKvGpDIp2Gj6crc2hWZCJ6CvZBkqfBbrGu642W4hyWcxJ3spaqJBQfsvlbK3eCEXvPYGivPidgwwjWIQzXbJZAZCdZAOwjXrhmil3ZC7QJ5WBnlHPVe1qX"
const FB_FETCH_INTERVAL = 120
var JWT_SECRET = []byte("temp_secret_until_proper_secrets_are_implemented")
