package main

import (
	"net/http"

	"github.com/jessemillar/health"
	"github.com/jessemillar/serenity/controllers"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.CORS())

	e.GET("/health", echo.WrapHandler(http.HandlerFunc(health.Check)))
	e.GET("/library/v1/books", controllers.GetBooksV1)

	e.Logger.Fatal(e.Start(":8000"))
}
