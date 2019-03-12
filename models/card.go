package models

import (
	"TrelloReportTools/utils"
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

var Utils utils.Utils

func (*DBCard) NewCard(card *trello.Card) Card {
	return Card{
		ID:                card.ID,
		Name:              card.Name,
		IDList:            card.IDList,
		TimeGuessForDone:  Utils.GetTimeGuessForDone(card.Name),
		TimeRealForDone:   Utils.GetTimeRealForDone(card.Name),
		Due:               card.Due,
		DateLastChangeDue: card.DateLastActivity,
	}
}

func (*DBCard) InsertOrUpdate(card Card) (err error) {
	db.Save(&card)
	return nil
}

func (*DBCard) GetAll() (cards []Card, err error) {
	err = db.Find(&cards).Error
	return
}

func (dbCard *DBCard) SaveTrelloCards(trelloCards []*trello.Card) {
	for _, trelloCard := range trelloCards {
		card := dbCard.NewCard(trelloCard)
		db.Save(&card)
	}
}

func (dbCard *DBCard) GetCardsChangedDueByTime(dayNumber int) []Card {
	dayBefore := time.Now().AddDate(0, 0, -dayNumber)
	cards, err := dbCard.GetAll()
	if err != nil {
		return nil
	}

	var cardsChangedDueDate []Card
	for _, v := range cards {
		if dayBefore.Before(*v.DateLastChangeDue) {
			cardsChangedDueDate = append(cardsChangedDueDate, v)
		}
	}

	return cardsChangedDueDate
}

func (dbCard *DBCard) UpdateCardsChangedDueDate(cardsOnBoard []*trello.Card, cardsOnDB []Card) {
	for _, cardBoard := range cardsOnBoard {
		for _, cardDB := range cardsOnDB {

			if cardBoard.ID == cardDB.ID {
				if Utils.CompareTime(cardBoard.Due, cardDB.Due) == false {
					cardDB.Due = cardBoard.Due
					cardDB.DateLastChangeDue = cardBoard.DateLastActivity
					dbCard.InsertOrUpdate(cardDB)
				}
				break
			}
		}
	}
}
