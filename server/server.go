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

// H - interface for sending JSON
type H map[string]interface{}

// Post - struct to contain post data
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

// Category - struct to contain category data
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

	e.GET("/post/:id", getPost(postsCollection))
	e.GET("/posts", getAllPosts(postsCollection))
	e.GET("/category/:id/", getCat(catCollection))
	e.POST("/category/", makeCat(catCollection))
	// e.PATCH("/category/", patchCat)
	e.DELETE("/category/", delCat(catCollection))
}

func getPost(collection *mongo.Collection) echo.HandlerFunc {
	return func(c echo.Context) error {
		var result *Post
		id, _ := strconv.Atoi(c.QueryParam("id"))
		filter := bson.D{{"postID", id}}
		err := collection.FindOne(context.TODO(), filter).Decode(&result)
		if err != nil {
			log.Fatal(err)
		}
		return c.JSON(http.StatusOK, H{
			"post": result,
		})
	}
}

func getAllPosts(collection *mongo.Collection) echo.HandlerFunc {
	return func(c echo.Context) error {
		count, _ := strconv.Atoi(c.QueryParam("id"))

		name := c.QueryParam("category")
		findOptions := options.Find()
		if count != 10 {
			findOptions.SetLimit(int64(count))
		} else {
			findOptions.SetLimit(10)
		}

		var posts []*Post
		var cur *mongo.Cursor
		var err error

		if name == "" {
			cur, err = collection.Find(context.TODO(), bson.D{{}}, findOptions)
		} else {
			filter := bson.D{{"post_category", name}}
			cur, err = collection.Find(context.TODO(), filter, findOptions)
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
		return c.JSON(http.StatusOK, H{
			"posts": posts,
		})
	}
}

func getCat(collection *mongo.Collection) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, _ := strconv.Atoi(c.QueryParam("id"))
		var result *Category
		filter := bson.D{{"categoryID", id}}
		err := collection.FindOne(context.TODO(), filter).Decode(&result)
		if err != nil {
			log.Fatal(err)
		}
		return c.JSON(http.StatusOK, H{
			"category": result,
		})
	}
}

func makeCat(collection *mongo.Collection) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, _ := strconv.Atoi(c.QueryParam("id"))
		newName := c.QueryParam("name")
		filter := bson.D{{"categoryID", id}}
		update := bson.D{
			{"$set", bson.D{
				{"name", newName},
			}},
		}
		_, err := collection.UpdateOne(context.TODO(), filter, update)
		if err != nil {
			log.Fatal(err)
		}
		return c.JSON(http.StatusOK, H{})
	}
}

func delCat(collection *mongo.Collection) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, _ := strconv.Atoi(c.QueryParam("id"))
		filter := bson.D{{"categoryID", id}}
		_, err := collection.DeleteOne(context.TODO(), filter)
		if err != nil {
			log.Fatal(err)
		}
		return c.JSON(http.StatusOK, H{})
	}
}
