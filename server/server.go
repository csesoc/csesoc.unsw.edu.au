package main

import (
	"context"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/google/uuid"
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

// User - struct to contain user data
type User struct {
	userID    string //sha256 the zid
	userToken string
	role      string
}

// Sponsor - struct to contain sponsor data
type Sponsor struct {
	sponsorID   uuid.UUID
	sponsorName string
	sponsorLogo string
	sponsorTier string
	expiry      int64
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
	sponsorCollection := client.Database("csesoc").Collection("sponsors")
	userCollection := client.Database("csesoc").Collection("users")

	// Add more API routes here
	e.GET("/api/v1/test", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.POST("/login/", login(userCollection))

	e.GET("/post/:id/", getPost(postsCollection))
	e.GET("/posts/", getAllPosts(postsCollection))
	e.POST("/post/", newPost(postsCollection))
	e.PUT("/post/:id/", updatePost(postsCollection))
	e.DELETE("/post/:id/", deletePost(postsCollection))

	e.GET("/category/:id/", getCat(catCollection))
	e.POST("/category/", newCat(catCollection))
	e.PATCH("/category/", patchCat(catCollection)) // UPDATE category info
	e.DELETE("/category/", deleteCat(catCollection))

	e.POST("/sponsor/", newSponsor(sponsorCollection))
}

// func login(collection *mongo.Collection) echo.HandlerFunc {
// 	return func(c echo.Context) error {

// 	}
// }

func getPost(collection *mongo.Collection) echo.HandlerFunc {
	return func(c echo.Context) error {
		var result *Post
		id, _ := strconv.Atoi(c.QueryParam("id"))
		category := c.QueryParam("category")
		filter := bson.D{{"postID", id}, {"category", category}}
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

func newPost(collection *mongo.Collection) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, _ := strconv.Atoi(c.FormValue("id"))
		category, _ := strconv.Atoi(c.FormValue("category"))
		showinMenu, _ := strconv.ParseBool(c.FormValue("showInMenu"))
		post := Post{
			postID:       id,
			postTitle:    c.FormValue("title"),
			postSubtitle: c.FormValue("subtitle"),
			postType:     c.FormValue("type"),
			postCategory: category,
			createdOn:    time.Now(),
			lastEditedOn: time.Now(),
			postContent:  c.FormValue("content"),
			showInMenu:   showinMenu,
		}

		_, err := collection.InsertOne(context.TODO(), post)
		if err != nil {
			log.Fatal(err)
		}
		return c.JSON(http.StatusOK, H{})
	}
}

func updatePost(collection *mongo.Collection) echo.HandlerFunc {
	return func(c echo.Context) error {
		postID, _ := strconv.Atoi(c.FormValue("id"))
		postTitle := c.FormValue("title")
		postSubtitle := c.FormValue("subtitle")
		postType := c.FormValue("type")
		postContent := c.FormValue("content")
		showinMenu, _ := strconv.ParseBool(c.FormValue("showInMenu"))
		filter := bson.D{{"postID", postID}}
		update := bson.D{
			{"$set", bson.D{
				{"postTitle", postTitle},
				{"postSubtitle", postSubtitle},
				{"postType", postType},
				{"postContent", postContent},
				{"lastEditedOn", time.Now()},
				{"showinMenu", showinMenu},
			}},
		}

		_, err := collection.UpdateOne(context.TODO(), filter, update)
		if err != nil {
			log.Fatal(err)
		}
		return c.JSON(http.StatusOK, H{})
	}
}

func deletePost(collection *mongo.Collection) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, _ := strconv.Atoi(c.FormValue("id"))
		filter := bson.D{{"postID", id}}
		_, err := collection.DeleteOne(context.TODO(), filter)
		if err != nil {
			log.Fatal(err)
		}
		return c.JSON(http.StatusOK, H{})
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

func newCat(collection *mongo.Collection) echo.HandlerFunc {
	return func(c echo.Context) error {
		catID, _ := strconv.Atoi(c.FormValue("id"))
		catName := c.FormValue("name")
		index, _ := strconv.Atoi(c.FormValue("index"))
		category := Category{
			categoryID:   catID,
			categoryName: catName,
			index:        index,
		}

		_, err := collection.InsertOne(context.TODO(), category)
		if err != nil {
			log.Fatal(err)
		}
		return c.JSON(http.StatusOK, H{})
	}
}

func patchCat(collection *mongo.Collection) echo.HandlerFunc {
	return func(c echo.Context) error {
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

		_, err := collection.UpdateOne(context.TODO(), filter, update)
		if err != nil {
			log.Fatal(err)
		}
		return c.JSON(http.StatusOK, H{})
	}
}

func deleteCat(collection *mongo.Collection) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, _ := strconv.Atoi(c.FormValue("id"))
		filter := bson.D{{"categoryID", id}}
		_, err := collection.DeleteOne(context.TODO(), filter)
		if err != nil {
			log.Fatal(err)
		}
		return c.JSON(http.StatusOK, H{})
	}
}

func newSponsor(collection *mongo.Collection) echo.HandlerFunc {
	return func(c echo.Context) error {
		name := c.FormValue("name")
		logo := c.FormValue("logo")
		tier := c.FormValue("tier")
		expiryStr := c.FormValue("expiry")
		expiryTime, _ := time.Parse(time.RFC3339, expiryStr)
		id := uuid.New()

		sponsor := Sponsor{
			sponsorID:   id,
			sponsorName: name,
			sponsorLogo: logo,
			sponsorTier: tier,
			expiry:      expiryTime.Unix(),
		}

		_, err := collection.InsertOne(context.TODO(), sponsor)
		if err != nil {
			log.Fatal(err)
		}
		return c.JSON(http.StatusOK, H{})
	}
}

func deleteSponsor(collection *mongo.Collection) echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.FormValue("id")
		parsedID := uuid.Must(uuid.Parse(id))
		filter := bson.D{{"sponsorID", id}}
		_, err := collection.DeleteOne(context.TODO(), filter)
		if err != nil {
			log.Fatal(err)
		}
		return c.JSON(http.StatusOK, H{})
	}
}
