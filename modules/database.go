package modules

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var db *gorm.DB
var err error

func init() {
	db, err = gorm.Open("sqlite3", "card.db")
	if err != nil {
		// Handle error
	}

	db.AutoMigrate(&Card{})
}

func GetCards() []Card {
	var cards []Card
	db.Find(&cards)

	return cards
}

func UpdateCard(card Card) {
	db.Save(&card)
}
