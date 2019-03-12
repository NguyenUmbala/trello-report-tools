package models

import (
	"time"

	"github.com/adlio/trello"
)

type Card struct {
	ID                string     `json:"id"`
	Name              string     `json:"name"`
	IDList            string     `json:"idlist`
	TimeGuessForDone  int        `json:"timeGuessForDone"`
	TimeRealForDone   int        `json:"timeRealForDone"`
	Due               *time.Time `json:"due"`
	DateLastChangeDue *time.Time `json:"dateLastActivity"`
}

type DBCard struct{}

func (*DBCard) NewCard(card *trello.Card, timeguess, timereal int) Card {
	return Card{
		ID:                card.ID,
		Name:              card.Name,
		IDList:            card.IDList,
		TimeGuessForDone:  timeguess,
		TimeRealForDone:   timereal,
		Due:               card.Due,
		DateLastChangeDue: card.DateLastActivity,
	}
}

func (*DBCard) InsertOrUpdate(card Card) (err error) {
	db.Save(&card)
	return nil
}

func (*DBCard) SelectAll() (cards []Card, err error) {
	err = db.Find(&cards).Error
	return
}
