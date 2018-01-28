package database

import (
	"database/sql"
	"io"
	"io/ioutil"
	"log"
	"os"

	"github.com/dropbox/dropbox-sdk-go-unofficial/dropbox"
	"github.com/dropbox/dropbox-sdk-go-unofficial/dropbox/files"
)

var Connection *sql.DB

func InitDB(filepath string) {
	log.Println("Initializing database")

	db, err := sql.Open("sqlite3", filepath)
	if err != nil {
		panic(err)
	}

	if db == nil {
		panic("Database is empty")
	}

	Connection = db
}

func DownloadDatabase(filepath string) {
	log.Println("Configuring database download")

	config := dropbox.Config{
		Token: os.Getenv("SERENITY_LIBRARY_DROPBOX"),
	}

	log.Println("Downloading database")

	dbf := files.New(config)
	_, content, err := dbf.Download(files.NewDownloadArg("/Kimico/BookBuddy.backup"))
	if err != nil {
		body, _ := ioutil.ReadAll(content)
		log.Println(string(body))
		panic(err)
	}

	log.Println("Saving database")

	outFile, err := os.Create("BookBuddy.backup")
	if err != nil {
		panic(err)
	}

	defer outFile.Close()

	_, err = io.Copy(outFile, content)
	if err != nil {
		panic(err)
	}
}
