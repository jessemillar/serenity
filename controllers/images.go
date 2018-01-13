package controllers

import (
	"net/http"

	"github.com/jessemillar/serenity/database"
	"github.com/jessemillar/serenity/helpers"
	"github.com/labstack/echo"
)

func GetImagesV1(c echo.Context) error {
	blob := helpers.ReadImage(database.Connection, c.Param("bookId"))

	return c.Blob(http.StatusOK, http.DetectContentType(blob), blob)
}
