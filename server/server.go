package main

import (
	"context"
	"log"
	"net/http"
	"regexp"
	"strconv"

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
type Category struct {
	CategoryID   int
	CategoryName string
	Index        int
}

// User - struct to contain user data
type User struct {
	UserID    string //sha256 the zid
	UserToken string
	Role      string
}

// Sponsor - struct to contain sponsor data
type Sponsor struct {
	SponsorID   uuid.UUID
	SponsorName string
	SponsorLogo string
	SponsorTier string
	Expiry      int64
}

// Claims - struct to store jwt data
type Claims struct {
	HashedZID   [32]byte
	FirstName   string
	Permissions string
	jwt.StandardClaims
}

func main() {
	// Initialse mailing client
	err := InitMailingClient()
	if err != nil {
		log.Fatal(err)
	}

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
	sponsorCollection := client.Database("csesoc").Collection("sponsors")
	// userCollection := client.Database("csesoc").Collection("users")

	// Add more API routes here
	e.GET("/api/v1/test", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	// e.POST("/login/", login(userCollection))

	// Routes for posts
	e.GET("/posts/", getPosts(postsCollection))
	e.POST("/post/", newPosts(postsCollection))
	e.PUT("/post/", updatePosts(postsCollection))
	e.DELETE("/post/", deletePosts(postsCollection))

	// Routes for categories
	e.GET("/category/:id/", getCats(catCollection))
	e.POST("/category/", newCats(catCollection))
	e.PATCH("/category/", patchCats(catCollection))
	e.DELETE("/category/", deleteCats(catCollection))

	// Routes for sponsors
	e.POST("/sponsor/", newSponsors(sponsorCollection))
	e.DELETE("/sponsor/", deleteSponsors(sponsorCollection))

	// Routes for enquiries
	e.POST("/enquiry/sponsorship", handleEnquiry("sponsorship@csesoc.org.au"))
	e.POST("/enquiry/info", handleEnquiry("info@csesoc.org.au"))
}

// Handle enquiry by forwarding it to a specified email
func handleEnquiry(targetEmail string) echo.HandlerFunc {
	return func(c echo.Context) error {
		name := c.FormValue("name")
		email := c.FormValue("email")
		message := c.FormValue("message")

		// Email validation
		emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)

		var status int
		if emailRegex.MatchString(email) && len(name) > 0 && len(message) > 0 {
			err := SendEmail(name, email, message, targetEmail)
			if err != nil {
				status = http.StatusServiceUnavailable
			} else {
				status = http.StatusOK
			}
		} else {
			status = http.StatusBadRequest
		}
		return c.JSON(status, H{})
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

func newSponsors(collection *mongo.Collection) echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.FormValue("token")
		expiryStr := c.FormValue("expiry")
		name := c.FormValue("name")
		logo := c.FormValue("logo")
		tier := c.FormValue("tier")
		NewSponsors(collection, expiryStr, name, logo, tier, token)
		return c.JSON(http.StatusOK, H{})
	}
}

func deleteSponsors(collection *mongo.Collection) echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.FormValue("token")
		id := c.FormValue("id")
		DeleteSponsors(collection, id, token)
		return c.JSON(http.StatusOK, H{})
	}
}
