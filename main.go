package main

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/labstack/echo/v4"
)

/*====================TEMPLATE CONFIGURATION===============================*/
type TemplateRenderer struct {
	templates *template.Template
}

func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	e := echo.New()
	//setup database
	err := initDB()
	if err != nil {
		log.Fatal(err)
	}
	defer closeDB()
	//setup template
	views := &TemplateRenderer{
		templates: template.Must(template.ParseGlob("templates/*.html")),
	}
	e.Renderer = views
	//routes
	e.Static("/static", "static")
	e.GET("/", welcomePage)
	e.GET("/submit", submitForm)
	e.GET("/short/:myurl", shortenUrlPage)
	e.GET("/:url", redirectToLongURL)
	e.GET("/:url/watch", watchShortedUrlPage)
	e.Logger.Fatal(e.Start(":7000"))
}

func welcomePage(c echo.Context) error {
	data := map[string]interface{}{
		"MainTitle":   "byteurl Main page",
		"MainContent": "Byteurl, short url with no limitation",
	}
	return c.Render(http.StatusOK, "index", data)
}

func submitForm(c echo.Context) error {
	data_input := c.QueryParam("original_url")
	main_url := url.QueryEscape(data_input)
	return c.Redirect(http.StatusSeeOther, fmt.Sprintf("/short/%s", main_url))
}

func shortenUrlPage(c echo.Context) error {
	input_param := c.Param("myurl")
	longUrl, err := url.QueryUnescape(input_param)
	if err != nil {
		log.Fatal(err)
	}
	short := generateAutomaticShortLink()
	currentTime, err := insertURLToDatabase(longUrl, short)
	if err != nil {
		return err
	} else {
		data := map[string]interface{}{
			"long_url": longUrl,
			"short": short,
			"created_at" : currentTime,
		}
		return c.Render(http.StatusOK, "shorten", data)
	}
}

func redirectToLongURL(c echo.Context) error {
	shorted_url := c.Param("url")
	original_url, err := getLongURLFromDatabase(shorted_url)
	if err != nil {
		return c.Render(http.StatusForbidden, "cannot find long url in database", nil)
	}

	if strings.HasPrefix(original_url, "w") {
		return c.Redirect(http.StatusSeeOther, fmt.Sprintf("https://%s", original_url))
	} else {
		return c.Redirect(http.StatusSeeOther, original_url)
	}
}

func watchShortedUrlPage(c echo.Context) error {
	shortUrl := c.Param("url")
	url_data, err := getURLInfoFromDatabase(shortUrl)
	if err != nil {
		log.Fatal(err)
	}
	data := map[string]interface{}{
		"id": url_data.id,
		"short_url": url_data.short_url,
		"long_url": url_data.long_url,
		"timestamp": url_data.timestamp,
	}
	return c.Render(http.StatusOK, "watch", data)
}
