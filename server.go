package main

import (
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"net/http"

	"os"

	"github.com/dropbox/dropbox-sdk-go-unofficial/dropbox"
	"github.com/dropbox/dropbox-sdk-go-unofficial/dropbox/files"
	"github.com/jessemillar/health"
	"github.com/jessemillar/serenity/controllers"
	"github.com/jessemillar/serenity/database"
	"github.com/jessemillar/serenity/helpers"
	"github.com/jessemillar/serenity/views"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	config := dropbox.Config{
		Token: os.Getenv("SERENITY_LIBRARY_DROPBOX"),
	}

	dbf := files.New(config)
	_, content, err := dbf.Download(files.NewDownloadArg("/Kimico/BookBuddy.backup"))
	if err != nil {
		body, _ := ioutil.ReadAll(content)
		fmt.Println(string(body))
		panic(err)
	}

	outFile, err := os.Create("BookBuddy.backup")
	if err != nil {
		panic(err)
	}

	defer outFile.Close()

	_, err = io.Copy(outFile, content)
	if err != nil {
		panic(err)
	}

	database.InitDB("BookBuddy.backup")

	e := echo.New()
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.CORS())

	templateEngine := &helpers.Template{
		Templates: template.Must(template.ParseGlob("public/*/*.html")),
	}

	e.Renderer = templateEngine

	e.GET("/health", echo.WrapHandler(http.HandlerFunc(health.Check)))
	e.Static("/library/*", "public")
	e.GET("/library", views.Main)
	e.GET("/library/v1/books", controllers.GetBooksV1)
	e.GET("/library/v1/books/:bookId/cover", controllers.GetCoverV1)

	e.Logger.Fatal(e.Start(":8000"))
}
