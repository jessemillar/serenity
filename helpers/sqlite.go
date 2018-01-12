package helpers

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

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
}

func InitDB(filepath string) *sql.DB {
	db, err := sql.Open("sqlite3", filepath)
	if err != nil {
		panic(err)
	}
	if db == nil {
		panic("db nil")
	}
	return db
}

func ReadItem(db *sql.DB) []Book {
	sql_readall := `
	SELECT ZTITLE, ZSUBTITLE, ZDISPLAYNAME, ZGENRE, ZSYNOPSIS, ZLCC, ZISBN, ZPUBLISHER, ZPUBLISHYEAR, ZPAGECOUNT FROM ZBOOK
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
		err2 := rows.Scan(&book.Title, &book.Subtitle, &book.Author, &book.Genre, &book.Synopsis, &book.LCC, &book.ISBN, &book.Publisher, &book.PublishYear, &book.PageCount)
		if err2 != nil {
			panic(err2)
		}
		allBooks = append(allBooks, book)
	}

	return allBooks
}
