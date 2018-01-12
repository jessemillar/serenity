package helpers

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type Book struct {
	Title string `json:"title"`
	//Subtitle string
	Author string `json:"author"`
	//Synopsis string
	LCC *string `json:"lcc,omitempty"`
	//Publisher string
	//ISBN int
	//Genre string
	//PageCount int
	//PublishYear int
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
	SELECT ZTITLE, ZDISPLAYNAME, ZLCC FROM ZBOOK
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
		err2 := rows.Scan(&book.Title, &book.Author, &book.LCC)
		if err2 != nil {
			panic(err2)
		}
		allBooks = append(allBooks, book)
	}

	return allBooks
}
