package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var db *gorm.DB

func init() {
	db, _ = gorm.Open("sqlite3", "card.db")
	db.AutoMigrate(&Card{})
}
