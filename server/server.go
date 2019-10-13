package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Post struct {
	postID       int
	postTitle    string
	postSubtitle string
	postType     string
	postCategory int
	createdOn    time.Time
	lastEditedOn time.Time
	postContent  string
	showInMenu   bool
}

type Category struct {
	categoryID   int
	categoryName string
	index        int
}

func main() {
	// Create new instance of echo
	e := echo.New()

	servePages(e)
	serveAPI(e)

	// Start echo instance on 1323 port
	e.Logger.Fatal(e.Start(":1323"))
}

func servePages(e *echo.Echo) {
	// Setup our assetHandler and point it to our static build location
	assetHandler := http.FileServer(http.Dir("../dist/"))

	// Setup a new echo route to load the build as our base path
	e.GET("/", echo.WrapHandler(assetHandler))

	// Serve our static assists under the /static/ endpoint
	e.GET("/js/*", echo.WrapHandler(assetHandler))
	e.GET("/css/*", echo.WrapHandler(assetHandler))

	echo.NotFoundHandler = func(c echo.Context) error {
		// render your 404 page
		return c.String(http.StatusNotFound, "not found page")
	}
}

func serveAPI(e *echo.Echo) {
	//Set client options
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	//Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	//Check connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	postsCollection := client.Database("csesoc").Collection("posts")
	catCollection := client.Database("csesoc").Collection("categories")

	// Add more API routes here
	e.GET("/api/v1/test", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.GET("/post/:id", getPost)
	e.GET("/posts", getAllPosts)
	e.GET("/category/:id/", getCat)
	e.POST("/category/", makeCat)
	e.PATCH("/category/", patchCat)
	e.DELETE("/category/", delCat)
}

func getPost(id int, collection Collection) *Post {
	var result *Post
	filter := bson.D{{"postID", id}}
	err := collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}
	return result
}

func getAllPosts(count int64, category string, collection Collection) []*Post {
	findOptions := options.Find()
	if count != 10 {
		findOptions.SetLimit(count)
	} else {
		findOptions.SetLimit(10)
	}

	var posts []*Post
	var cur *Cursor
	var err error
	if category == "" {
		cur, err := collection.Find(context.TODO(), bson.D{{}}, findOptions)
	} else {
		filter := bson.D{{"post_category", category}}
		cur, err := collection.Find(context.TODO(), filter, findOptions)
	}

	if err != nil {
		log.Fatal(err)
	}

	for cur.Next(context.TODO()) {
		var elem Post
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}

		posts = append(posts, &elem)
	}

	return posts
}

func getCat(id int, collection Collection) *Category {
	var result *Category
	filter := bson.D{{"categoryID", id}}
	err := collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}
	return result
}

func makeCat(id int, collection Collection, newName string) {
	filter := bson.D{{"categoryID", id}}
	update := bson.D{
		{"$set", bson.D{
			{"name", newName},
		}},
	}
	updateResult, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
}

func delCat(id int, collection Collection) {
	filter := bson.D{{"categoryID", id}}
	deleteResult, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
	}
}
