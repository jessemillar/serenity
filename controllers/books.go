package controllers

import (
	"net/http"

	"github.com/jessemillar/serenity/helpers"
	"github.com/labstack/echo"
)

func GetBooksV1(c echo.Context) error {
	db := helpers.InitDB("BookBuddy.backup")

	return c.JSON(http.StatusOK, helpers.ReadItem(db))
}
