package services

import (
	"TrelloReportTools/models"
	"TrelloReportTools/store"
	"time"

	"github.com/adlio/trello"
)

type ServiceGet struct{}

var db store.Database
var dbTrello store.TrelloDatabase

func init() {
	db = *db.Start()
	dbTrello = *dbTrello.Start()
}

func (*ServiceGet) GetCardsIsOpenOnWeek(idboard, is, sort, created, list string) []*trello.Card {
	var query = "board:" + idboard + " is:" + is + " sort:" + sort + " created:" + created + " list:" + list
	cards, err := dbTrello.SelectManyTrello(query)
	if err != nil {
		return nil
	}

	return cards
}

func (*ServiceGet) GetCardsChangedDueByTime(dayNumber int) []models.Card {
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

func (*ServiceGet) GetCardsOnBoard(idBoard string) []*trello.Card {
	query := "board:" + idBoard
	cards, err := dbTrello.SelectManyTrello(query)
	if err != nil {
		return nil
	}

	return cards
}

func (*ServiceGet) GetCardsChangedDueDateOnBoard(idBoard string) (cards []*trello.Card) {
	cardsOnBoard, _ := dbTrello.SelectManyTrello("board:" + idBoard)
	cardsOnDB, _ := db.SelectAll()

	// Compare due date of two cards
	for _, cardBoard := range cardsOnBoard {
		for _, cardDB := range cardsOnDB {
			if cardBoard.ID == cardDB.ID {
				if UtilTime.CompareTime(cardBoard.Due, cardDB.Due) == false {
					cards = append(cards, cardBoard)
				}
				break
			}
		}
	}

	return
}
