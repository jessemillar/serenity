package main

import (
	"github.com/jessemillar/serenity/helpers"
	"fmt"
)

func main() {
	db:= helpers.InitDB("BookBuddy.backup")
	fmt.Println(helpers.ReadItem(db))
}
