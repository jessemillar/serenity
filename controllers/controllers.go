package controllers

import (
	"net/http"
	"sync/atomic"

	"github.com/jessemillar/serenity/models"
)

// Create a response counter
var responseId int64 = 1

func buildResponse(apiVersion string, data *models.Data, err *models.Error, links *[]models.Hateoas) (int, models.Response) {
	responseCode := http.StatusOK

	// Make the actual response code mirror the Error
	if err != nil {
		responseCode = err.Code
	}

	// Build the JSON response from the above return values
	responseBody := models.NewResponse(apiVersion, int(atomic.LoadInt64(&responseId)), data, err, links)

	// Increment the response ID in a thread safe manner
	atomic.AddInt64(&responseId, 1)

	return responseCode, responseBody
}
