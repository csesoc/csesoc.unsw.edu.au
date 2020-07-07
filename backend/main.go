package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	// Doing a local import
	//. "csesoc.unsw.edu.au/m/v2/server/category"
	//. "csesoc.unsw.edu.au/m/v2/server/post"
	. "csesoc.unsw.edu.au/m/v2/server"
	. "csesoc.unsw.edu.au/m/v2/server/faq"
	. "csesoc.unsw.edu.au/m/v2/server/login"
	. "csesoc.unsw.edu.au/m/v2/server/mailing"
	. "csesoc.unsw.edu.au/m/v2/server/social"
	. "csesoc.unsw.edu.au/m/v2/server/sponsor"
	// events "csesoc.unsw.edu.au/m/v2/server/events"

	_ "csesoc.unsw.edu.au/m/v2/docs"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"

	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type (
	// CustomValidator - struct for simple validation
	CustomValidator struct {
		validator *validator.Validate
	}
)

// Validate - custom validator for user input.
func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

// @title CSESoc Website Swagger API
// @version 1.0
// @description Swagger API for the CSESoc Website project.
// @termsOfService http://swagger.io/terms/

// @contact.name Project Lead
// @contact.email projects.website@csesoc.org.au

// @BasePath /api/v1

// @securityDefinitions.bearerAuth BearerAuthKey
// @in header
// @name Authorization
func main() {
	// Create new instance of echo
	e := echo.New()
	e.Debug = true
	// Validator for structs used
	e.Validator = &CustomValidator{validator: validator.New()}

	servePages(e)
	serveAPI(e)
	eventFetchTimer()
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
	if !DEVELOPMENT {
		DispatchEnquiryBundles()
		DispatchFeedbackBundle()
	}
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
	// SwaggerUI can be accessed from: http://localhost:1323/swagger/index.html
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
	e.POST("/login", TempLogin)

	// POSTS
	// e.GET("/api/posts/", getPosts(postsCollection))
	// e.POST("/api/post/", newPosts(postsCollection), middleware.JWT(JWT_SECRET))
	// e.PUT("/api/post/", updatePosts(postsCollection), middleware.JWT(JWT_SECRET))
	// e.DELETE("/api/post/", deletePosts(postsCollection), middleware.JWT(JWT_SECRET))

	// CATEGORIES
	// e.GET("/api/category/:id/", getCats(catCollection))
	// e.POST("/api/category/", newCats(catCollection), middleware.JWT(JWT_SECRET))
	// e.PATCH("/api/category/", patchCats(catCollection), middleware.JWT(JWT_SECRET))
	// e.DELETE("/api/category/", deleteCats(catCollection), middleware.JWT(JWT_SECRET))

	v1 := e.Group("/api/v1")
	{
		// SPONSORS
		SponsorSetup(client)
		sponsors := v1.Group("/sponsors")
		{
			sponsors.GET("/:name", GetSponsor)
			sponsors.POST("", NewSponsor, middleware.JWT(JWT_SECRET))
			sponsors.DELETE("/:name", DeleteSponsor, middleware.JWT(JWT_SECRET))
			sponsors.GET("", GetSponsors)
		}

		// MAILING
		MailingSetup()
		mailing := v1.Group("/mailing")
		{
			mailing.POST("/general", HandleGeneralMessage)
			mailing.POST("/sponsorship", HandleSponsorshipMessage)
			mailing.POST("/feedback", HandleFeedbackMessage)
		}

		// FAQ
		faq := v1.Group("/faq")
		{
			faq.GET("", GetFaq)
		}

		// SOCIAL
		social := v1.Group("/social")
		{
			social.GET("", GetSocial)
		}
	}
}
