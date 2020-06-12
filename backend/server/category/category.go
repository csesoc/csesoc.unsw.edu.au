package category

import (
	"context"
	"net/http"
	"strconv"
	"log"

	. "csesoc.unsw.edu.au/m/v2/server"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// Category - struct to contain category data
type Category struct {
	CategoryID   int
	CategoryName string
	Index        int
}

func getCats(collection *mongo.Collection) echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.FormValue("token")
		id, _ := strconv.Atoi(c.QueryParam("id"))
		result := GetCats(collection, id, token)
		
		return c.JSON(http.StatusAccepted, result)
	}
}

func newCats(collection *mongo.Collection) echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.FormValue("token")
		catID, _ := strconv.Atoi(c.FormValue("id"))
		index, _ := strconv.Atoi(c.FormValue("index"))
		name := c.FormValue("name")
		NewCats(collection, catID, index, name, token)
		return c.JSON(http.StatusOK, H{})
	}
}

func patchCats(collection *mongo.Collection) echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.FormValue("token")
		catID, _ := strconv.Atoi(c.FormValue("id"))
		name := c.FormValue("name")
		index, _ := strconv.Atoi(c.FormValue("index"))
		PatchCats(collection, catID, name, index, token)
		return c.JSON(http.StatusOK, H{})
	}
}

func deleteCats(collection *mongo.Collection) echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.FormValue("token")
		id, _ := strconv.Atoi(c.FormValue("id"))
		DeleteCats(collection, id, token)
		return c.JSON(http.StatusOK, H{})
	}
}

// GetCats - Retrieve a category from the database
func GetCats(collection *mongo.Collection, id int, token string) Category {
	// if !validToken(token) {
	// 	return nil
	// }

	var result Category
	filter := bson.D{{Key: "categoryid", Value: id}}

	// Find a category
	err := collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}

	return result
}

// NewCats - Add a new category
func NewCats(collection *mongo.Collection, catID int, index int, name string, token string) {
	// if !validToken(token) {
	// 	return
	// }

	category := Category{
		CategoryID:   catID,
		CategoryName: name,
		Index:        index,
	}

	_, err := collection.InsertOne(context.TODO(), category)
	if err != nil {
		log.Fatal(err)
	}
}

// PatchCats - Update a category with new information
func PatchCats(collection *mongo.Collection, catID int, name string, index int, token string) {
	// if !validToken(token) {
	// 	return
	// }

	filter := bson.D{{Key: "categoryid", Value: catID}}
	update := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "categoryname", Value: name},
			{Key: "index", Value: index},
		}},
	}

	// Find a category by id and update it
	_, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
}

// DeleteCats - Delete a category from the database
func DeleteCats(collection *mongo.Collection, id int, token string) {
	// if !validToken(token) {
	// 	return
	// }

	filter := bson.D{{Key: "categoryid", Value: id}}

	// Find a category by id and delete it
	_, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
	}
}
