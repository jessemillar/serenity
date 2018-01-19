package helpers

import (
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/jessemillar/serenity/database"
)

func GetCover(id string) ([]byte, error) {
	isbn, err := ConvertBookBuddyIdToIsbn(database.Connection, id)
	if err != nil {
		panic(err)
	}

	url := "http://covers.openlibrary.org/b/isbn/" + strconv.Itoa(isbn) + "-L.jpg"

	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	// If we don't find a cover image online
	if len(body) < 2500 {
		// Return the BookBuddy cover image
		return ReadBookBuddyImage(database.Connection, id), nil
	}

	return body, nil
}
