package main

import (
	"context"
	"crypto/sha256"
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/ldap.v2"
)

var jwtKey = []byte("secret_text")

// Auth - to login
func Auth(collection *mongo.Collection, zid string, password string, permissions string) string {
	// Connect to UNSW LDAP server
	l, err := ldap.Dial("tcp", "ad.unsw.edu.au")
	if err != nil {
		log.Fatal(err)
	}

	// Attempt to sign in using credentials
	hashedZID := sha256.Sum256([]byte(zid))
	stringZID := string(hashedZID[:])
	username := zid + "ad.unsw.edu.au"

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

	userFound := searchResult.Entries[0]
	expirationTime := time.Now().Add(time.Hour * 24)
	claims := &Claims{
		hashedZID:   hashedZID,
		firstName:   userFound.GetAttributeValue("firstName"),
		permissions: permissions,
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

	return tokenString
}

// validToken - returns true if a token is valid and false otherwise.
func validToken(tokenString string) bool {
	claims := &Claims{}
	tkn, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return false
		}
	}

	if !tkn.Valid || claims.permissions != "staff" {
		return false
	}

	return true
}
