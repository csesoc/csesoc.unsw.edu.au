package login

import (
	"net/http"
	"time"

	. "csesoc.unsw.edu.au/m/v2/server"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

var tempUsers = map[string]string{
	"z5123456": "t3stP@ssw0rd",
}

// createJwtToken creates a new JWT and returns it as a string.
func createJwtToken(zID string, admin bool) (string, time.Time, error) {
	unsignedToken := jwt.New(jwt.SigningMethodHS256)
	claims := unsignedToken.Claims.(jwt.MapClaims)
	expTime := time.Now().Add(time.Hour * 72)
	claims["zID"] = zID
	claims["admin"] = true
	claims["exp"] = expTime.Unix()

	token, err := unsignedToken.SignedString(JWT_SECRET)
	if err != nil {
		return "", time.Now(), err
	}
	return token, expTime, nil
}

// TempLogin allows auth functionality to be tested before LDAP is implemented
func TempLogin(c echo.Context) error {
	userzID := c.QueryParam("zID")
	password := c.QueryParam("password")
	expectedPass, ok := tempUsers[userzID]
	if !ok || password != expectedPass {
		return c.String(http.StatusUnauthorized, "Your username or password was incorrect.")
	}

	// Create a new token.
	token, expTime, err := createJwtToken(userzID, true)
	if err != nil {
		return c.String(http.StatusInternalServerError, "500 Internal Error")
	}

	// Create a cookie to store the JWT.
	tokenCookie := new(http.Cookie)
	tokenCookie.Name = "activeToken"
	tokenCookie.Value = token
	tokenCookie.Expires = expTime
	tokenCookie.HttpOnly = true
	// c.SetCookie(tokenCookie)

	return c.JSON(http.StatusOK, map[string]string{
		"message": "Success!!",
		"token":   token,
	})

}
