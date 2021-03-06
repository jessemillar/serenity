package controllers

import (
	"github.com/jessemillar/serenity/database"
	"github.com/jessemillar/serenity/helpers"
	"github.com/labstack/echo"
)

func GetBooksV1(c echo.Context) error {
	const apiVersion = "1.0.0"

	data, err := helpers.ReadBookBuddyCatalogue(database.Connection, c.Request().URL.String())

	responseStatus, responseBody := buildResponse(apiVersion, data, err, nil)

	return c.JSON(responseStatus, responseBody)
}
