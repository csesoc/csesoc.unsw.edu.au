package main

import (
	"bufio"
	"context"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"

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

func main() {
	// Initialse SMTP mailing client
	InitMailClient()

	// Create new instance of echo
	e := echo.New()
	// Validator for structs used
	e.Validator = &CustomValidator{validator: validator.New()}

	servePages(e)
	serveAPI(e)

	// Start echo instance on 1323 port
	e.Debug = true
	e.Logger.Fatal(e.Start(":1323"))

	// Safely stop echo instance
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		if scanner.Text() == "quit" {
			e.Close()
			DispatchEnquiryBundles()
			DispatchFeedbackBundle()
		}
	}
}

func servePages(e *echo.Echo) {
	// Setup our assetHandler and point it to our static build location
	assetHandler := http.FileServer(http.Dir("../../frontend/dist/"))

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
	//Set client options - this line needs to change when moving from local deployment to docker containers
	clientOptions := options.Client().ApplyURI("mongodb://mongo:27017")
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

	println("Serving API...")

	// Creating collections
	postsCollection := client.Database("csesoc").Collection("posts")
	catCollection := client.Database("csesoc").Collection("categories")
	SponsorSetup(client)
	// userCollection := client.Database("csesoc").Collection("users")

	// Add more API routes here
	e.GET("/api/v1/test", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	// e.POST("/login/", login(userCollection))

	// Routes for posts
	e.GET("/api/posts/", getPosts(postsCollection))
	e.POST("/api/post/", newPosts(postsCollection))
	e.PUT("/api/post/", updatePosts(postsCollection))
	e.DELETE("/api/post/", deletePosts(postsCollection))

	// Routes for categories
	e.GET("/api/category/:id/", getCats(catCollection))
	e.POST("/api/category/", newCats(catCollection))
	e.PATCH("/api/category/", patchCats(catCollection))
	e.DELETE("/api/category/", deleteCats(catCollection))

	// Routes for sponsors
	e.GET("/api/sponsor/", GetSponsor())
	e.POST("/api/sponsor/", NewSponsor())
	e.DELETE("/api/sponsor/", DeleteSponsor())
	e.GET("/api/sponsors/", GetSponsors())

	// Routes for enquiries
	e.POST("/api/enquiry/sponsorship", HandleEnquiry("sponsorship@csesoc.org.au"))
	e.POST("/api/enquiry/info", HandleEnquiry("info@csesoc.org.au"))
	e.POST("/api/enquiry/feedback", HandleFeedback())

	// Routes for faq
	e.GET("/api/faq/", GetFaq())
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
