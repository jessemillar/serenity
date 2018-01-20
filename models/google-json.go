package models

import (
	"encoding/json"
	"time"
)

// The structs in this file represent a subset of the Google JSON Style Guide (https://google.github.io/styleguide/jsoncstyleguide.xml)
// We use pointers (`*`) for Data and Errors so that when Data and/or Errors is empty, they don't show up in the JSON payload
type Response struct {
	ApiVersion string `json:"apiVersion"`
	ID         int    `json:"id"`
	// According to the Google JSON Style Guide, a response should have either a Data or Error object, but not both
	Data  *Data     `json:"data,omitempty"`
	Error *Error    `json:"error,omitempty"`
	Links []Hateoas `json:"links,omitempty"`
}

type Data struct {
	Updated    time.Time         `json:"updated"`
	TotalItems int               `json:"totalItems,omitempty"`
	Items      []json.RawMessage `json:"items"`
}

type Error struct {
	Code    int         `json:"code"`
	Message string      `json:"string"`
	Errors  []ErrorItem `json:"errors,omitempty"`
}

type ErrorItem struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Domain  string `json:"domain,omitempty"`
	Reason  string `json:"reason,omitempty"`
}

func NewResponse(apiVersion string, id int, data *Data, err *Error, links *[]Hateoas) Response {
	response := Response{ApiVersion: apiVersion, ID: id}

	if err != nil {
		response.Error = err
	} else if data != nil {
		// Set the updated timestamp
		data.Updated = time.Now()
		// Populate the "totalItems" property for convenience
		data.TotalItems = len(data.Items)
		response.Data = data
	}
	if links != nil {
		response.Links = *links
	}
	return response
}

func NewError(code int, message string) *Error {
	return &Error{Code: code, Message: message}
}

func NewErrorItemFromError(err Error) *ErrorItem {
	return &ErrorItem{Code: err.Code, Message: err.Message}
}
