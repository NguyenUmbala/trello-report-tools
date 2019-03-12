package services

import (
	"TrelloReportTools/models"
	"TrelloReportTools/utils"
	"time"

	"github.com/adlio/trello"
)

type Card struct{}

var DB models.DBCard
var TrelloDB models.DBTrelloCard
var UtilString utils.UtilString
var UtilTime utils.UtilTime

func (*Card) GetCardsIsOpenOnWeek(idboard, list string) []*trello.Card {
	cards, err := TrelloDB.SelectMany("board:" + idboard + " is:open sort:created created:week list:" + list)
	if err != nil {
		return nil
	}

	return cards
}

func (*Card) GetCardsChangedDueByTime(dayNumber int) []models.Card {
	dayBefore := time.Now().AddDate(0, 0, -dayNumber)
	cards, err := DB.SelectAll()
	if err != nil {
		return nil
	}

	var cardsChangedDueDate []models.Card
	for _, v := range cards {
		if dayBefore.Before(*v.DateLastChangeDue) {
			cardsChangedDueDate = append(cardsChangedDueDate, v)
		}
	}

	return cardsChangedDueDate
}

func (*Card) GetCardsOnBoard(idBoard string) []*trello.Card {
	cards, err := TrelloDB.SelectMany("board:" + idBoard)
	if err != nil {
		return nil
	}

	return cards
}

func (*Card) GetCardsOnDB() []models.Card {
	cards, err := DB.SelectAll()
	if err != nil {
		return nil
	}

	return cards
}

func (*Card) UpdateCardsChangedDueDate(cardsOnBoard []*trello.Card, cardsOnDB []models.Card) {
	for _, cardBoard := range cardsOnBoard {
		for _, cardDB := range cardsOnDB {
			if cardBoard.ID == cardDB.ID {
				if UtilTime.CompareTime(cardBoard.Due, cardDB.Due) == false {
					cardDB.Due = cardBoard.Due
					DB.InsertOrUpdate(cardDB)
				}
				break
			}
		}
	}
}

func (*Card) ConvertCardsTrelloDBToDB(trelloCards []*trello.Card) (dbCards []models.Card) {
	for _, trelloCard := range trelloCards {
		timeGuess := UtilString.GetTimeGuessForDone(trelloCard.Name)
		timeReal := UtilString.GetRealTimeOfDone(trelloCard.Name)

		card := DB.NewCard(trelloCard, timeGuess, timeReal)
		dbCards = append(dbCards, card)
	}

	return
}

func (*Card) SaveCards(cards []models.Card) bool {
	for _, card := range cards {
		err := DB.InsertOrUpdate(card)
		if err != nil {
			return false
		}
	}

	return true
}
