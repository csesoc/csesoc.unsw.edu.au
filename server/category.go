package main

import (
	"context"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
)

func GetCat(c echo.Context) error {
	id, _ := strconv.Atoi(c.QueryParam("id"))
	var result *Category
	filter := bson.D{{"categoryID", id}}

	// Find a category
	err := collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}

	return c.JSON(http.StatusOK, H{
		"category": result,
	})
}

func NewCat(c echo.Context) error {
	catID, _ := strconv.Atoi(c.FormValue("id"))
	index, _ := strconv.Atoi(c.FormValue("index"))

	category := Category{
		categoryID:   catID,
		categoryName: c.FormValue("name"),
		index:        index,
	}

	_, err := collection.InsertOne(context.TODO(), category)
	if err != nil {
		log.Fatal(err)
	}
	return c.JSON(http.StatusOK, H{})
}

func PatchCat(c echo.Context) error {
	categoryID, _ := strconv.Atoi(c.FormValue("id"))
	categoryName := c.FormValue("name")
	index, _ := strconv.Atoi(c.FormValue("index"))
	filter := bson.D{{"categoryID", categoryID}}
	update := bson.D{
		{"$set", bson.D{
			{"categoryName", categoryName},
			{"index", index},
		}},
	}

	// Find a category by id and update it
	_, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Fatal(err)
	}

	return c.JSON(http.StatusOK, H{})
}

func DeleteCat(c echo.Context) error {
	id, _ := strconv.Atoi(c.FormValue("id"))
	filter := bson.D{{"categoryID", id}}

	// Find a category by id and delete it
	_, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
	}

	return c.JSON(http.StatusOK, H{})
}
