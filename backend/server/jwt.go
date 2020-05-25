package main

import (
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

var jwtSecret = []byte("temp_secret_until_proper_secrets_are_implemented")

var tempUsers = map[string]string{
	"testUser": "t3stP@ssw0rd",
}

// handleAuthRequest handles an authentication request made to the server.
func handleAuthRequest(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	// Handle incorrect password/non-existent user.
	if password != tempUsers[username] {
		return echo.ErrUnauthorized
	}

	// Create a token.
	authToken := jwt.New(jwt.SigningMethodHS256)

	// Set JWT claims.
	claims := authToken.Claims.(jwt.MapClaims)
	claims["name"] = username
	claims["admin"] = true
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	tok, err := authToken.SignedString(jwtSecret)

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]string{
		"token": tok,
	})
}
