package main

import (
	"context"
	"crypto/sha256"
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"gopkg.in/ldap.v2"
)

func auth(c echo.Context) error {
	// Connect to UNSW LDAP server
	l, err := ldap.Dial("tcp", "ad.unsw.edu.au")
	if err != nil {
		log.Fatal(err)
	}

	// Attempt to sign in using credentials
	zid := c.FormValue("zid")
	hashedZID := sha256.Sum256([]byte(zid))
	stringZID := string(hashedZID[:])
	username := zid + "ad.unsw.edu.au"
	password := c.FormValue("password")

	err = l.Bind(username, password)
	if err != nil {
		log.Fatal(err)
	}

	// Retrieve first name from Identity Manager
	baseDN := "OU=IDM_People,OU=IDM,DC=ad,DC=unsw,DC=edu,DC=au"
	searchScope := ldap.ScopeWholeSubtree
	aliases := ldap.NeverDerefAliases
	retrieveAttributes := []string{"givenName"}
	searchFilter := "cn=" + username //cn = common name

	searchRequest := ldap.NewSearchRequest(
		baseDN, searchScope, aliases, 0, 0, false,
		searchFilter, retrieveAttributes, nil,
	)

	searchResult, err := l.Search(searchRequest)
	if err != nil {
		log.Fatal(err)
	}

	// Encode user details into a JWT and turn it into a string
	jwtKey := []byte("secret_text")
	userFound := searchResult.Entries[0]
	expirationTime := time.Now().Add(time.Hour * 24)
	claims := &Claims{
		hashedZID: hashedZID,
		firstName: userFound.GetAttributeValue("firstName"),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	tokenJWT := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, _ := tokenJWT.SignedString(jwtKey)

	// Insert a new user into the collection if user has never logged in before
	// Or update the existing token if it has expired
	user := User{
		userID:    stringZID,
		userToken: tokenString,
		role:      "user", // Change this???
	}

	var isValidUser *User
	userFilter := bson.D{{"userID", stringZID}}
	err = collection.FindOne(context.TODO(), userFilter).Decode(&isValidUser)

	if isValidUser == nil { // Never logged in before
		_, err = collection.InsertOne(context.TODO(), user)
		if err != nil {
			log.Fatal(err)
		}
	} else { // Logged in before - check validity of token
		claims = &Claims{}
		decodedToken, _ := jwt.ParseWithClaims(isValidUser.userToken, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})
		decodedTokenString, _ := decodedToken.SignedString(jwtKey)

		if !decodedToken.Valid { // Logged in before but token is invalid - replace with new token
			filter := bson.D{{"userID", stringZID}}
			update := bson.D{
				{"$set", bson.D{
					{"userToken", decodedTokenString},
				}},
			}
			_, err = collection.UpdateOne(context.TODO(), filter, update)
			if err != nil {
				log.Fatal(err)
			}
		}
	}

	return c.JSON(http.StatusOK, H{
		"token": tokenString,
	})
}
