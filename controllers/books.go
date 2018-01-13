package controllers

import (
	"net/http"

	"github.com/jessemillar/serenity/database"
	"github.com/jessemillar/serenity/helpers"
	"github.com/labstack/echo"
)

func GetBooksV1(c echo.Context) error {
	return c.JSON(http.StatusOK, helpers.ReadBooks(database.Connection, c.Request().URL.String()))
}
