package main

import (
	"log"
	"net/http"

	"github.com/jessemillar/health"
	"github.com/jessemillar/serenity/controllers"
	"github.com/jessemillar/serenity/database"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	database.DownloadDatabase("BookBuddy.backup")
	database.InitDB("BookBuddy.backup")

	log.Println("Configuring server")

	e := echo.New()
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.CORS())

	e.GET("/health", echo.WrapHandler(http.HandlerFunc(health.Check)))
	e.Static("/*", "frontend/dist")
	e.GET("/library/v1/books", controllers.GetBooksV1)
	e.GET("/library/v1/books/:bookId/cover", controllers.GetCoverV1)
	e.GET("/library/v1/wishlist", controllers.GetWishlistV1)
	e.GET("/library/v1/wishlist/:bookId/cover", controllers.GetCoverV1)

	e.Logger.Fatal(e.Start(":8000"))
}
