package database

import (
	"TrelloReportTools/modules"

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

	db.AutoMigrate(&modules.Card{})
}

func GetCards() []modules.Card {
	var cards []modules.Card
	db.Find(&cards)

	return cards
}

func SaveCard(card modules.Card) {
	db.Create(&card)
}

func UpdateCard(card modules.Card) {
	if db.NewRecord(card) == true { // Create new card
		db.Create(&card)
	} else { // Update old card
		db.Save(&card)
	}
}
