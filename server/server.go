package main

import (
	"context"
	"crypto/sha256"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"gopkg.in/ldap.v2"

	"go.mongodb.org/mongo-driver/bson"
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
	return auth(c echo.Context)
}

func getPost(collection *mongo.Collection) echo.HandlerFunc {
	return GetPosts(c echo.Context)
}

func getAllPosts(collection *mongo.Collection) echo.HandlerFunc {
	return GetAllPosts(c echo.Context)
}

func newPost(collection *mongo.Collection) echo.HandlerFunc {
	return NewPost(c echo.Context)
}

func updatePost(collection *mongo.Collection) echo.HandlerFunc {
	return UpdatePost(c echo.Context)
}

func deletePost(collection *mongo.Collection) echo.HandlerFunc {
	return DeletePost(c echo.Context)
}

func getCat(collection *mongo.Collection) echo.HandlerFunc {
	return GetCat(c echo.Context)
}

func newCat(collection *mongo.Collection) echo.HandlerFunc {
	return NewCat(c echo.Context)
}

func patchCat(collection *mongo.Collection) echo.HandlerFunc {
	return PatchCat(c echo.Context) 
}

func deleteCat(collection *mongo.Collection) echo.HandlerFunc {
	return DeleteCat(c echo.Context) 
}

func newSponsor(collection *mongo.Collection) echo.HandlerFunc {
	return NewSponsor(c echo.Context)
}

func deleteSponsor(collection *mongo.Collection) echo.HandlerFunc {
	return DeleteSponsor(c echo.Context)
}
