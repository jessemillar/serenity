package helpers

import (
"database/sql"
_"github.com/mattn/go-sqlite3"
)

type Book struct {
	Title string
	//Subtitle string
	Author string
	//Synopsis string
	LCC string
	//Publisher string
	//ISBN int
	//Genre string
	//PageCount int
	//PublishYear int
}

func InitDB(filepath string) *sql.DB {
	db, err := sql.Open("sqlite3", filepath)
	if err != nil { panic(err) }
	if db == nil { panic("db nil") }
	return db
}

func ReadItem(db *sql.DB) []Book {
	sql_readall := `SELECT ZTITLE FROM ZBOOK`

	rows, err := db.Query(sql_readall)
	if err != nil { panic(err) }
	defer rows.Close()

	var result []Book
	for rows.Next() {
		item := Book{}
		err2 := rows.Scan(&item.Title)
		if err2 != nil { panic(err2) }
		result = append(result, item)
	}

	return result
}