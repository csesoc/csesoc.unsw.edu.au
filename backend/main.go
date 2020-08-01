package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	// Import utilities
	. "csesoc.unsw.edu.au/m/v2/server"

	"csesoc.unsw.edu.au/m/v2/server/events"
	"csesoc.unsw.edu.au/m/v2/server/faq"
	"csesoc.unsw.edu.au/m/v2/server/login"
	"csesoc.unsw.edu.au/m/v2/server/mailing"
	"csesoc.unsw.edu.au/m/v2/server/resources"
	"csesoc.unsw.edu.au/m/v2/server/social"
	"csesoc.unsw.edu.au/m/v2/server/sponsor"

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
	println("Web server is online :)")

	// Disable the fetch timer until we get access to the actual CSESoc page
	if !DEVELOPMENT {
		go events.FetchTimer()
	}

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
		mailing.DispatchEnquiryBundles()
		mailing.DispatchFeedbackBundle()
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

	// AUTHENTICATION
	e.POST("/login", login.TempLogin)

	v1 := e.Group("/api/v1")
	{
		// SPONSORS
		sponsor.Setup(client)
		sponsorsAPI := v1.Group("/sponsors")
		{
			sponsorsAPI.GET("/:name", sponsor.HandleGetSingle)
			sponsorsAPI.POST("", sponsor.HandleNew, middleware.JWT(JWT_SECRET))
			sponsorsAPI.DELETE("/:name", sponsor.HandleDelete, middleware.JWT(JWT_SECRET))
			sponsorsAPI.GET("", sponsor.HandleGetMultiple)
		}

		// MAILING
		mailing.Setup()
		mailingAPI := v1.Group("/mailing")
		{
			mailingAPI.POST("/general", mailing.HandleGeneralMessage)
			mailingAPI.POST("/sponsorship", mailing.HandleSponsorshipMessage)
			mailingAPI.POST("/feedback", mailing.HandleFeedbackMessage)
		}

		// FAQ
		faqsAPI := v1.Group("/faq")
		{
			faqsAPI.GET("", faq.HandleGet)
		}

		// SOCIAL
		socialAPI := v1.Group("/social")
		{
			socialAPI.GET("", social.HandleGet)
		}

		// EVENTS
		eventsAPI := v1.Group("/events")
		{
			eventsAPI.GET("", events.HandleGet)
		}

		// RESOURCES
		resources.Setup(client)
		resourcesAPI := v1.Group("/resources")
		{
			resourcesAPI.GET("/preview", resources.HandleGetPreview)
		}
	}
}
