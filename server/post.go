package main

import (
	"context"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetPost(c echo.Context) error {
	var result *Post
	id, _ := strconv.Atoi(c.QueryParam("id"))
	category := c.QueryParam("category")

	// Search for post by id and category
	filter := bson.D{{"postID", id}, {"category", category}}
	err := collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}
	return c.JSON(http.StatusOK, H{
		"post": result,
	})
}

func GetAllPosts(c echo.Context) error {
	count, _ := strconv.Atoi(c.QueryParam("id"))
	cat := c.QueryParam("category")

	findOptions := options.Find()
	if count != 10 {
		findOptions.SetLimit(int64(count))
	} else {
		findOptions.SetLimit(10)
	}

	var posts []*Post
	var cur *mongo.Cursor
	var err error

	if cat == "" { // No specified category
		cur, err = collection.Find(context.TODO(), bson.D{{}}, findOptions)
	} else {
		filter := bson.D{{"post_category", cat}}
		cur, err = collection.Find(context.TODO(), filter, findOptions)
	}

	if err != nil {
		log.Fatal(err)
	}

	// Iterate through all results
	for cur.Next(context.TODO()) {
		var elem Post
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}

		posts = append(posts, &elem)
	}

	return c.JSON(http.StatusOK, H{
		"posts": posts,
	})
}

func NewPost(c echo.Context) error {
	id, _ := strconv.Atoi(c.FormValue("id"))
	category, _ := strconv.Atoi(c.FormValue("category"))
	showinMenu, _ := strconv.ParseBool(c.FormValue("showInMenu"))

	post := Post{
		postID:           id,
		postTitle:        c.FormValue("title"),
		postSubtitle:     c.FormValue("subtitle"),
		postType:         c.FormValue("type"),
		postCategory:     category,
		createdOn:        time.Now(),
		lastEditedOn:     time.Now(),
		postContent:      c.FormValue("content"),
		postLinkGithub:   c.FormValue("linkGithub"),
		postLinkFacebook: c.FormValue("linkFacebook"),
		showInMenu:       showinMenu,
	}

	_, err := collection.InsertOne(context.TODO(), post)
	if err != nil {
		log.Fatal(err)
	}

	return c.JSON(http.StatusOK, H{})
}

func UpdatePost(c echo.Context) error {
	postID, _ := strconv.Atoi(c.FormValue("id"))
	postTitle := c.FormValue("title")
	postSubtitle := c.FormValue("subtitle")
	postType := c.FormValue("type")
	postCategory := c.FormValue("category")
	postContent := c.FormValue("content")
	postLinkGithub := c.FormValue("linkGithub")
	postLinkFacebook := c.FormValue("linkFacebook")
	showinMenu, _ := strconv.ParseBool(c.FormValue("showInMenu"))

	filter := bson.D{{"postID", postID}}
	update := bson.D{
		{"$set", bson.D{
			{"postTitle", postTitle},
			{"postSubtitle", postSubtitle},
			{"postType", postType},
			{"postCategory", postCategory},
			{"lastEditedOn", time.Now()},
			{"postContent", postContent},
			{"postLinkGithub", postLinkGithub},
			{"postLinkFacebook", postLinkFacebook},
			{"showinMenu", showinMenu},
		}},
	}

	// Find a post by id and update it
	_, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Fatal(err)
	}

	return c.JSON(http.StatusOK, H{})
}

func DeletePost(c echo.Context) error {
	id, _ := strconv.Atoi(c.FormValue("id"))
	filter := bson.D{{"postID", id}}

	// Find a post by id and delete it
	_, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
	}

	return c.JSON(http.StatusOK, H{})
}
