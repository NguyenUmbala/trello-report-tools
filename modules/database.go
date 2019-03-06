package modules

import (
	"fmt"

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

func SaveCard(card Card) {
	if err := db.Create(&card).Error; err != nil {
		// Handle error
		fmt.Println(err)
	}
}

func UpdateCard(card Card) {
	if db.NewRecord(card) == true { // Create new card
		db.Create(&card)
	} else { // Update old card
		db.Save(&card)
	}
}
