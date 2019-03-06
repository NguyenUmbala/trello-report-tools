package modules

import (
	"time"

	"github.com/adlio/trello"
)

type Card struct {
	ID                string     `json:"id"`
	Name              string     `json:"name"`
	TimeGuessForDone  int        `json:"timeGuessForDone"`
	TimeRealForDone   int        `json:"timeRealForDone"`
	Due               *time.Time `json:"due"`
	DateLastChangeDue *time.Time `json:"dateLastActivity"`
}

func (c Card) NewCard(card *trello.Card) Card {
	c.ID = card.ID
	c.Name = card.Name
	c.TimeGuessForDone = GetTimeGuessForDone(c.Name)
	c.TimeRealForDone = GetRealTimeOfDone(c.Name)
	c.Due = card.Due
	c.DateLastChangeDue = card.DateLastActivity
	return c
}
