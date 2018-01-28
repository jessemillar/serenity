package helpers

import (
	"database/sql"
	"encoding/json"

	"github.com/jessemillar/serenity/models"
	_ "github.com/mattn/go-sqlite3"
)

// TODO Clean up panics and db passing

type Book struct {
	Title       string  `json:"title"`
	Subtitle    *string `json:"subtitle,omitempty"`
	Author      string  `json:"author"`
	Genre       *string `json:"genre,omitempty"`
	Synopsis    *string `json:"synopsis,omitempty"`
	LCC         *string `json:"lcc,omitempty"`
	ISBN        int     `json:"isbn"`
	Publisher   *string `json:"publisher,omitempty"`
	PublishYear *int    `json:"publishYear,omitempty"`
	PageCount   *int    `json:"pageCount,omitempty"`
	Image       *string `json:"image,omitempty"`
}

type Image struct {
	Blob *string
}

func ReadBookBuddyBooks(db *sql.DB, path string, query string) (*models.Data, *models.Error) {
	rows, err := db.Query(query)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	data := &models.Data{}

	for rows.Next() {
		book := Book{}
		err2 := rows.Scan(&book.Title, &book.Subtitle, &book.Author, &book.Genre, &book.Synopsis, &book.LCC, &book.ISBN, &book.Publisher, &book.PublishYear, &book.PageCount, &book.Image)
		if err2 != nil {
			panic(err2)
		}

		if len(*book.Image) > 0 {
			*book.Image = path + "/" + *book.Image + "/cover"
		}

		marshaledBook, err := json.Marshal(book)
		if err != nil {
			panic(err)
		}

		data.Items = append(data.Items, marshaledBook)
	}

	return data, nil
}

func ReadBookBuddyCatalogue(db *sql.DB, path string) (*models.Data, *models.Error) {
	query := `
	SELECT ZTITLE, ZSUBTITLE, ZDISPLAYNAME, ZGENRE, ZSYNOPSIS, ZLCC, ZISBN, ZPUBLISHER, ZPUBLISHYEAR, ZPAGECOUNT, ZBOOK.Z_PK
	FROM ZBOOK
	INNER JOIN ZAUTHOR ON ZBOOK.ZAUTHORINFO=ZAUTHOR.Z_PK
	WHERE ZWISHLIST=0;
`

	data, err := ReadBookBuddyBooks(db, path, query)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func ReadBookBuddyWishlist(db *sql.DB, path string) (*models.Data, *models.Error) {
	query := `
	SELECT ZTITLE, ZSUBTITLE, ZDISPLAYNAME, ZGENRE, ZSYNOPSIS, ZLCC, ZISBN, ZPUBLISHER, ZPUBLISHYEAR, ZPAGECOUNT, ZBOOK.Z_PK
	FROM ZBOOK
	INNER JOIN ZAUTHOR ON ZBOOK.ZAUTHORINFO=ZAUTHOR.Z_PK
	WHERE ZWISHLIST=1;
`

	data, err := ReadBookBuddyBooks(db, path, query)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func ConvertBookBuddyIdToIsbn(db *sql.DB, id string) (int, error) {
	query := `
	SELECT ZISBN
	FROM ZBOOK
	WHERE Z_PK=$1;
`

	row := db.QueryRow(query, id)

	book := Book{}
	err := row.Scan(&book.ISBN)
	if err != nil {
		panic(err)
	}

	return book.ISBN, nil
}

func ReadBookBuddyImage(db *sql.DB, id string) []byte {
	query := `
	SELECT ZIMAGE
	FROM ZIMAGE
	WHERE ZBOOK=$1;
`

	row := db.QueryRow(query, id)

	bookImage := Image{}
	err := row.Scan(&bookImage.Blob)
	if err != nil {
		panic(err)
	}

	return []byte(*bookImage.Blob)
}
