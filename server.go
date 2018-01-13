package main

import (
	"net/http"

	"os"

	"fmt"

	"io"

	"github.com/dropbox/dropbox-sdk-go-unofficial/dropbox"
	"github.com/dropbox/dropbox-sdk-go-unofficial/dropbox/files"
	"github.com/jessemillar/health"
	"github.com/jessemillar/serenity/controllers"
	"github.com/jessemillar/serenity/database"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	config := dropbox.Config{
		Token:    os.Getenv("SERENITY_LIBRARY_DROPBOX"),
		LogLevel: dropbox.LogInfo, // if needed, set the desired logging level. Default is off
	}

	dbf := files.New(config)
	meta, content, err := dbf.Download(files.NewDownloadArg("/BookBuddy.backup"))
	fmt.Println(meta, content, err)

	outFile, err := os.Create("BookBuddy.backup")
	// handle err
	defer outFile.Close()
	_, err = io.Copy(outFile, content)
	if err != nil {
		panic(err)
	}

	database.InitDB("BookBuddy.backup")

	e := echo.New()
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.CORS())

	e.GET("/health", echo.WrapHandler(http.HandlerFunc(health.Check)))
	e.GET("/library/v1/books", controllers.GetBooksV1)
	e.GET("/library/v1/books/:bookId/cover.jpg", controllers.GetImagesV1)

	e.Logger.Fatal(e.Start(":8000"))
}
