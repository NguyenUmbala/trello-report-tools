package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var db *gorm.DB

func init() {
	var err error
	db, err = gorm.Open("sqlite3", "card.db")
	if err != nil {
		fmt.Println(err)
	}

	db.AutoMigrate(&Card{})
}
