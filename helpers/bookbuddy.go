package helpers

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

// TODO Clean up panics and db passing

type Book struct {
	Title       string  `json:"title"`
	Subtitle    *string `json:"subtitle,omitempty"`
	Author      string  `json:"author"`
	Genre       string  `json:"genre"`
	Synopsis    *string `json:"synopsis,omitempty"`
	LCC         *string `json:"lcc,omitempty"`
	ISBN        int     `json:"isbn"`
	Publisher   *string `json:"publisher,omitempty"`
	PublishYear int     `json:"publishYear"`
	PageCount   *int    `json:"pageCount,omitempty"`
	Image       *string `json:"image,omitempty"`
}

type Image struct {
	Blob *string
}

func ReadBooks(db *sql.DB, path string) []Book {
	sql_readall := `
	SELECT ZTITLE, ZSUBTITLE, ZDISPLAYNAME, ZGENRE, ZSYNOPSIS, ZLCC, ZISBN, ZPUBLISHER, ZPUBLISHYEAR, ZPAGECOUNT, ZBOOK.Z_PK
	FROM ZBOOK
	INNER JOIN ZAUTHOR ON ZBOOK.ZAUTHORINFO=ZAUTHOR.Z_PK;
`

	rows, err := db.Query(sql_readall)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var allBooks []Book
	for rows.Next() {
		book := Book{}
		err2 := rows.Scan(&book.Title, &book.Subtitle, &book.Author, &book.Genre, &book.Synopsis, &book.LCC, &book.ISBN, &book.Publisher, &book.PublishYear, &book.PageCount, &book.Image)
		if err2 != nil {
			panic(err2)
		}

		if len(*book.Image) > 0 {
			*book.Image = path + "/" + *book.Image + "/cover.jpg"
		}

		allBooks = append(allBooks, book)
	}

	return allBooks
}

func ReadImage(db *sql.DB, id string) []byte {
	sql_readall := `
	SELECT ZIMAGE
	FROM ZIMAGE
	WHERE Z_PK=$1;
`

	row := db.QueryRow(sql_readall, id)

	bookImage := Image{}
	err := row.Scan(&bookImage.Blob)
	if err != nil {
		panic(err)
	}

	return []byte(*bookImage.Blob)

}
