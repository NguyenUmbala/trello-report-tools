package services

import (
	"TrelloReportTools/models"
	"TrelloReportTools/store"
	"time"

	"github.com/adlio/trello"
)

type Service struct{}

var db store.Database
var dbTrello store.TrelloDatabase

func init() {
	db = db.Start()
	dbTrello = dbTrello.Start()
}

func (s Service) GetCardsIsOpenOnWeek(idboard, is, sort, created, list string) []*trello.Card {
	var query = "board:" + idboard + " is:" + is + " sort:" + sort + " created:" + created + " list:" + list

	cards, err := dbTrello.SelectManyTrello(query)
	if err != nil {
		return nil
	}

	return cards
}

func (s Service) GetCardsChangedDueByTime(dayNumber int) []models.Card {
	cards, _ := db.SelectAll()
	dayBefore := time.Now().AddDate(0, 0, -dayNumber)

	var cardsChangedDueDate []models.Card
	for _, v := range cards {
		if dayBefore.Before(*v.DateLastChangeDue) {
			cardsChangedDueDate = append(cardsChangedDueDate, v)
		}
	}

	return cardsChangedDueDate
}
