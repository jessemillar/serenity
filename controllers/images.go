package controllers

import (
	"net/http"

	"github.com/jessemillar/serenity/helpers"
	"github.com/labstack/echo"
)

func GetImagesV1(c echo.Context) error {
	db := helpers.InitDB("BookBuddy.backup")

	blob := helpers.ReadImage(db, c.Param("bookId"))

	return c.Blob(http.StatusOK, http.DetectContentType(blob), blob)
}
