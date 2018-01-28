package helpers

import (
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/jessemillar/serenity/database"
	"github.com/jessemillar/serenity/models"
)

func GetCover(id string) ([]byte, *models.Error) {
	isbn, err := ConvertBookBuddyIdToIsbn(database.Connection, id)
	if err != nil {
		return []byte{}, models.NewError(http.StatusInternalServerError, err.Error())
	}

	url := "http://covers.openlibrary.org/b/isbn/" + strconv.Itoa(isbn) + "-L.jpg"

	resp, err := http.Get(url)
	if err != nil {
		return []byte{}, models.NewError(http.StatusInternalServerError, err.Error())
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, models.NewError(http.StatusInternalServerError, err.Error())
	}

	// If we don't find a cover image online
	if len(body) < 2500 {
		// Return the BookBuddy cover image
		blob, err := ReadBookBuddyImage(database.Connection, id)
		if err != nil {
			return []byte{}, models.NewError(http.StatusInternalServerError, err.Error())
		}

		return blob, nil
	}

	return body, nil
}
