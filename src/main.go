package main

import (
	"html/template"
	"io"
	"net/http"

	"github.com/labstack/echo"
)

//Template ......
type Template struct {
	templates *template.Template
}

func main() {
	e := echo.New()
	e.Renderer = parseTemplates()
	e.GET("/", serveLanding)
	e.Start(":8000")
}

func serveLanding(c echo.Context) error {
	return c.Render(http.StatusOK, "landingPage.html", map[string]interface{}{
		"title":   "Home",
		"message": "Test: Welcome to the Website",
	})
}

//Render templates
func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func parseTemplates() *Template {
	t := &Template{
		templates: template.Must(template.ParseGlob("templates/*.html")),
	}
	return t
}
