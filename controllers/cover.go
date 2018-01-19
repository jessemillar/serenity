package controllers

import (
	"net/http"

	"github.com/jessemillar/serenity/helpers"
	"github.com/labstack/echo"
)

func GetCoverV1(c echo.Context) error {
	blob, err := helpers.GetCover(c.Param("bookId"))
	if err != nil {
		panic(err)
	}

	return c.Blob(http.StatusOK, http.DetectContentType(blob), blob)
}
