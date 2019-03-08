package models

import (
	"time"

	"github.com/adlio/trello"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var db *gorm.DB

func init() {
	db, _ = gorm.Open("sqlite3", "card.db")
	db.AutoMigrate(&Card{})
}

type Card struct {
	ID                string     `json:"id"`
	Name              string     `json:"name"`
	TimeGuessForDone  int        `json:"timeGuessForDone"`
	TimeRealForDone   int        `json:"timeRealForDone"`
	Due               *time.Time `json:"due"`
	DateLastChangeDue *time.Time `json:"dateLastActivity"`
}

func (c Card) NewCard(card *trello.Card, timeGuess, timeReal int) Card {
	c.ID = card.ID
	c.Name = card.Name
	c.TimeGuessForDone = timeGuess
	c.TimeRealForDone = timeReal
	c.Due = card.Due
	c.DateLastChangeDue = card.DateLastActivity
	return c
}

func GetAllCard() (cards []Card) {
	db.Find(&cards)

	return cards
}

func UpdateCard(card Card) {
	db.Save(card)
}
