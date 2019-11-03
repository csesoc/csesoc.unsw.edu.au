package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"

	"go.mongodb.org/mongo-driver/bson"
)

func NewSponsor(c echo.Context) error {
	expiryStr := c.FormValue("expiry")
	expiryTime, _ := time.Parse(time.RFC3339, expiryStr)
	id := uuid.New()

	sponsor := Sponsor{
		sponsorID:   id,
		sponsorName: c.FormValue("name"),
		sponsorLogo: c.FormValue("logo"),
		sponsorTier: c.FormValue("tier"),
		expiry:      expiryTime.Unix(),
	}

	_, err := collection.InsertOne(context.TODO(), sponsor)
	if err != nil {
		log.Fatal(err)
	}
	return c.JSON(http.StatusOK, H{})
}

func DeleteSponsor(c echo.Context) error {
	id := c.FormValue("id")
	parsedID := uuid.Must(uuid.Parse(id))

	// Find a sponsor by ID and delete it
	filter := bson.D{{"sponsorID", parsedID}}
	_, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
	}

	return c.JSON(http.StatusOK, H{})
}
