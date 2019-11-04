package main

import (
	"context"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// H - interface for sending JSON
type H map[string]interface{}

// Post - struct to contain post data
type Post struct {
	postID           int
	postTitle        string
	postSubtitle     string
	postType         string
	postCategory     int
	createdOn        time.Time
	lastEditedOn     time.Time
	postContent      string
	postLinkGithub   string
	postLinkFacebook string
	showInMenu       bool
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

// Claims - struct to store jwt data
type Claims struct {
	hashedZID [32]byte
	firstName string
	jwt.StandardClaims
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

	// Creating collections
	postsCollection := client.Database("csesoc").Collection("posts")
	catCollection := client.Database("csesoc").Collection("categories")
	sponsorCollection := client.Database("csesoc").Collection("sponsors")
	userCollection := client.Database("csesoc").Collection("users")

	// Add more API routes here
	e.GET("/api/v1/test", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.POST("/login/", login(userCollection))

	// Routes for posts
	e.GET("/post/:id/", getPost(postsCollection))
	e.GET("/posts/", getAllPosts(postsCollection))
	e.POST("/post/", newPost(postsCollection))
	e.PUT("/post/:id/", updatePost(postsCollection))
	e.DELETE("/post/:id/", deletePost(postsCollection))

	// Routes for categories
	e.GET("/category/:id/", getCat(catCollection))
	e.POST("/category/", newCat(catCollection))
	e.PATCH("/category/", patchCat(catCollection))
	e.DELETE("/category/", deleteCat(catCollection))

	// Routes for sponsors
	e.POST("/sponsor/", newSponsor(sponsorCollection))
	e.DELETE("/sponsor/", deleteSponsor(sponsorCollection))
}

func login(collection *mongo.Collection) echo.HandlerFunc {
	return func(c echo.Context) error {
		zid := c.FormValue("zid")
		password := c.FormValue("password")
		tokenString := auth(collection, zid, password)
		return c.JSON(http.StatusOK, H{
			"token": tokenString,
		})
	}
}

func getPost(collection *mongo.Collection) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, _ := strconv.Atoi(c.QueryParam("id"))
		category := c.QueryParam("category")
		result := GetPost(collection, id, category)
		return c.JSON(http.StatusOK, H{
			"post": result,
		})
	}
}

func getAllPosts(collection *mongo.Collection) echo.HandlerFunc {
	return func(c echo.Context) error {
		count, _ := strconv.Atoi(c.QueryParam("id"))
		cat := c.QueryParam("category")
		posts := GetAllPosts(collection, count, cat)
		return c.JSON(http.StatusOK, H{
			"posts": posts,
		})
	}
}

func newPost(collection *mongo.Collection) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, _ := strconv.Atoi(c.FormValue("id"))
		category, _ := strconv.Atoi(c.FormValue("category"))
		showInMenu, _ := strconv.ParseBool(c.FormValue("showInMenu"))
		title := c.FormValue("title")
		subtitle := c.FormValue("subtitle")
		postType := c.FormValue("type")
		content := c.FormValue("content")
		github := c.FormValue("linkGithub")
		fb := c.FormValue("linkFacebook")
		NewPost(collection, id, category, showInMenu, title, subtitle, postType, content, github, fb)
		return c.JSON(http.StatusOK, H{})
	}
}

func updatePost(collection *mongo.Collection) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, _ := strconv.Atoi(c.FormValue("id"))
		category, _ := strconv.Atoi(c.FormValue("category"))
		showInMenu, _ := strconv.ParseBool(c.FormValue("showInMenu"))
		title := c.FormValue("title")
		subtitle := c.FormValue("subtitle")
		postType := c.FormValue("type")
		content := c.FormValue("content")
		github := c.FormValue("linkGithub")
		fb := c.FormValue("linkFacebook")
		UpdatePost(collection, id, category, showInMenu, title, subtitle, postType, content, github, fb)
		return c.JSON(http.StatusOK, H{})
	}
}

func deletePost(collection *mongo.Collection) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, _ := strconv.Atoi(c.FormValue("id"))
		DeletePost(collection, id)
		return c.JSON(http.StatusOK, H{})
	}
}

func getCat(collection *mongo.Collection) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, _ := strconv.Atoi(c.QueryParam("id"))
		result := GetCat(collection, id)
		return c.JSON(http.StatusOK, H{
			"category": result,
		})
	}
}

func newCat(collection *mongo.Collection) echo.HandlerFunc {
	return func(c echo.Context) error {
		catID, _ := strconv.Atoi(c.FormValue("id"))
		index, _ := strconv.Atoi(c.FormValue("index"))
		name := c.FormValue("name")
		NewCat(collection, catID, index, name)
		return c.JSON(http.StatusOK, H{})
	}
}

func patchCat(collection *mongo.Collection) echo.HandlerFunc {
	return func(c echo.Context) error {
		catID, _ := strconv.Atoi(c.FormValue("id"))
		name := c.FormValue("name")
		index, _ := strconv.Atoi(c.FormValue("index"))
		PatchCat(collection, catID, name, index)
		return c.JSON(http.StatusOK, H{})
	}
}

func deleteCat(collection *mongo.Collection) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, _ := strconv.Atoi(c.FormValue("id"))
		DeleteCat(collection, id)
		return c.JSON(http.StatusOK, H{})
	}
}

func newSponsor(collection *mongo.Collection) echo.HandlerFunc {
	return func(c echo.Context) error {
		expiryStr := c.FormValue("expiry")
		name := c.FormValue("name")
		logo := c.FormValue("logo")
		tier := c.FormValue("tier")
		NewSponsor(collection, expiryStr, name, logo, tier)
		return c.JSON(http.StatusOK, H{})
	}
}

func deleteSponsor(collection *mongo.Collection) echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.FormValue("id")
		DeleteSponsor(collection, id)
		return c.JSON(http.StatusOK, H{})
	}
}
