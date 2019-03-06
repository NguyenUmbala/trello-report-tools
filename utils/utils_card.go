package utils

import (
	"TrelloReportTools/modules"

	"github.com/adlio/trello"
)

type Card struct{}

func (Card) NewCard(card *trello.Card) modules.Card {
	var c modules.Card
	c.ID = card.ID
	c.Name = card.Name
	c.TimeGuessForDone = GetTimeGuessForDone(card.Name)
	c.TimeRealForDone = GetRealTimeOfDone(card.Name)
	c.Due = card.Due
	c.DateLastChangeDue = card.DateLastActivity
	return c
}

// if trello.Card.Due != modules.Card.Due => return false
func (Card) CompareTwoCard(cardBoard *trello.Card, cardDB modules.Card) bool {
	if cardBoard.ID != cardDB.ID {
		return false
	}
	if cardBoard.Due == nil && cardDB.Due == nil {
		return false
	}
	if cardBoard.Due == nil && cardDB.Due != nil {
		return true
	}
	if cardBoard.Due != nil && cardDB.Due == nil {
		return true
	}
	if cardBoard.Due.String() == cardDB.Due.String() {
		return false
	}
	return true
}
