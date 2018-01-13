package database

import "database/sql"

var Connection *sql.DB

func InitDB(filepath string) {
	db, err := sql.Open("sqlite3", filepath)
	if err != nil {
		panic(err)
	}

	if db == nil {
		panic("db nil")
	}

	Connection = db
}
