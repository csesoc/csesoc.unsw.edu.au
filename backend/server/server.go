package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
	_ "github.com/swaggo/echo-swagger/example/docs"

	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type (
	// H - interface for sending JSON
	H map[string]interface{}

	// CustomValidator - struct for simple validation
	CustomValidator struct {
		validator *validator.Validate
	}

	// Post - struct to contain post data
	Post struct {
		PostID           int
		PostTitle        string
		PostSubtitle     string
		PostType         string
		PostCategory     int
		CreatedOn        int64
		LastEditedOn     int64
		PostContent      string
		PostLinkGithub   string
		PostLinkFacebook string
		ShowInMenu       bool
	}

	// Category - struct to contain category data
	Category struct {
		CategoryID   int
		CategoryName string
		Index        int
	}

	// User - struct to contain user data
	User struct {
		UserID    string //sha256 the zid
		UserToken string
		Role      string
	}

	// Claims - struct to store jwt data
	Claims struct {
		HashedZID   [32]byte
		FirstName   string
		Permissions string
		jwt.StandardClaims
	}
)

// Validate - custom validator for user input.
func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

// @title CSESoc Website Swagger API
// @version 1.0
// @description Swagger API for the CSESoc Website Project.
// @termsOfService http://swagger.io/terms/

// @contact.name Tommy Truong
// @contact.email projects.website@csesoc.org.au

// @BasePath /api/v1
func main() {
	// Create new instance of echo
	e := echo.New()
	e.Debug = true
	// Validator for structs used
	e.Validator = &CustomValidator{validator: validator.New()}

	servePages(e)
	serveAPI(e)

	println("Web server is online :)")

	// Bind quit to listen to Interrupt signals
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)

	// Running server on a subroutine enables a gracefully shutdown
	// Reference: https://echo.labstack.com/cookbook/graceful-shutdown
	go func() {
		// Start echo instance on 1323 port
		if err := e.Start(":1323"); err != nil {
			e.Logger.Info("Error: shutting down the server")

			// Send interrupt signal to begin shutdown
			quit <- os.Signal(os.Interrupt)
		}
	}()

	///////////
	// SHUTDOWN
	///////////

	// Wait for interrupt signal
	<-quit
	// Dispatch bundles on shutdown
	DispatchEnquiryBundles()
	DispatchFeedbackBundle()
	// Gracefully shutdown the server with a timeout of 10 seconds
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}

func servePages(e *echo.Echo) {

	println("Serving pages...")

	// Setup our assetHandler and point it to our static build location
	assetHandler := http.FileServer(http.Dir("../../frontend/dist/"))

	// Setup a new echo route to load the build as our base path
	e.GET("/", echo.WrapHandler(assetHandler))

	// Serve our static assists under the /static/ endpoint
	e.GET("/js/*", echo.WrapHandler(assetHandler))
	e.GET("/css/*", echo.WrapHandler(assetHandler))

	echo.NotFoundHandler = func(c echo.Context) error {
		// TODO: Render your 404 page
		return c.String(http.StatusNotFound, "not found page")
	}
}

func serveAPI(e *echo.Echo) {

	////////////////
	// MongoDB Setup
	////////////////

	println("Initialising MongoDB...")

	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb://mongo:27017")
	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	// Check connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	////////////////
	// Swagger Setup
	////////////////

	println("Serving Swagger...")
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	///////////////////////
	// Binding API Handlers
	///////////////////////

	println("Serving API...")

	// Creating collections
	// postsCollection := client.Database("csesoc").Collection("posts")
	// catCollection := client.Database("csesoc").Collection("categories")
	// userCollection := client.Database("csesoc").Collection("users")

	// AUTHENTICATION
	// e.POST("/login/", login(userCollection))

	// POSTS
	// e.GET("/api/v1/posts/", getPosts(postsCollection))
	// e.POST("/api/v1/post/", newPosts(postsCollection))
	// e.PUT("/api/v1/post/", updatePosts(postsCollection))
	// e.DELETE("/api/v1/post/", deletePosts(postsCollection))

	// CATEGORIES
	// e.GET("/api/v1/category/:id/", getCats(catCollection))
	// e.POST("/api/v1/category/", newCats(catCollection))
	// e.PATCH("/api/v1/category/", patchCats(catCollection))
	// e.DELETE("/api/v1/category/", deleteCats(catCollection))

	v1 := e.Group("/api/v1")
	{
		// SPONSORS
		SponsorSetup(client)
		sponsors := v1.Group("/sponsors")
		{
			sponsors.GET("/:name", GetSponsor())
			sponsors.POST("", NewSponsor())
			sponsors.DELETE("/:name", DeleteSponsor())
			sponsors.GET("", GetSponsors())
		}

		// MAILING
		MailingSetup()
		mailing := v1.Group("/mailing")
		{
			mailing.POST("/general", HandleMessage(GeneralType))
			mailing.POST("/sponsorship", HandleMessage(SponsorshipType))
			mailing.POST("/feedback", HandleMessage(FeedbackType))
		}

		// FAQ
		faq := v1.Group("/faq")
		{
			faq.GET("", GetFaq())
		}
	}
}

// func login(collection *mongo.Collection) echo.HandlerFunc {
// 	return func(c echo.Context) error {
// 		zid := c.FormValue("zid")
// 		password := c.FormValue("password")
// 		permissions := c.FormValue("permissions")
// 		tokenString := Auth(collection, zid, password, permissions)
// 		return c.JSON(http.StatusOK, H{
// 			"token": tokenString,
// 		})
// 	}
// }

func getPosts(collection *mongo.Collection) echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.QueryParam("id")
		count, _ := strconv.Atoi(c.QueryParam("nPosts"))
		category := c.QueryParam("category")
		if id == "" {
			posts := GetAllPosts(collection, count, category)
			return c.JSON(http.StatusOK, H{
				"post": posts,
			})
		}

		idInt, _ := strconv.Atoi(id)
		result := GetPosts(collection, idInt, category)
		return c.JSON(http.StatusOK, H{
			"post": result,
		})
	}
}

func newPosts(collection *mongo.Collection) echo.HandlerFunc {
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
		NewPosts(collection, id, category, showInMenu, title, subtitle, postType, content, github, fb)
		return c.JSON(http.StatusOK, H{})
	}
}

func updatePosts(collection *mongo.Collection) echo.HandlerFunc {
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
		UpdatePosts(collection, id, category, showInMenu, title, subtitle, postType, content, github, fb)
		return c.JSON(http.StatusOK, H{})
	}
}

func deletePosts(collection *mongo.Collection) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, _ := strconv.Atoi(c.FormValue("id"))
		DeletePosts(collection, id)
		return c.JSON(http.StatusOK, H{})
	}
}

func getCats(collection *mongo.Collection) echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.FormValue("token")
		id, _ := strconv.Atoi(c.QueryParam("id"))
		result := GetCats(collection, id, token)
		return c.JSON(http.StatusOK, H{
			"category": result,
		})
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
