package services

import (
	"TrelloReportTools/models"
	"fmt"
	"time"

	"github.com/adlio/trello"
)

func init() {
	SaveAllCards()
}

func SaveAllCards() {
	idBoard := "iCBtQXmr"
	cardsOnBoard := GetCardsOnBoard(idBoard)
	var timeGuess, timeReal int
	var tmpCard models.Card

	for _, v := range cardsOnBoard {
		timeGuess = GetTimeGuessForDone(v.Name)
		timeReal = GetRealTimeOfDone(v.Name)

		models.UpdateCard(tmpCard.NewCard(v, timeGuess, timeReal))
	}
}

func GetCardsIsOpenOnWeek(idBoard, nameList string) []*trello.Card {
	var trelloAPI models.TrelloAPI
	var query = "board:" + idBoard + " is:open sort:created created:week list:" + nameList

	cardsOnList, err := trelloAPI.GetCards(query)
	if err != nil {
		fmt.Println(err)
	}

	return cardsOnList
}

func GetCardsOnBoard(idBoard string) []*trello.Card {
	var trelloAPI models.TrelloAPI
	var query = "board:" + idBoard

	cardsOnBoard, err := trelloAPI.GetCards(query)
	if err != nil {
		fmt.Println(err)
	}

	return cardsOnBoard
}

func GetCardsChangedDueInTime(dayNumber int) []models.Card {
	cardsOnDB := models.GetAllCard()
	dayBefore := time.Now().AddDate(0, 0, -dayNumber)

	var cardsChangedDueDate []models.Card
	for _, v := range cardsOnDB {
		if dayBefore.Before(*v.DateLastChangeDue) {
			cardsChangedDueDate = append(cardsChangedDueDate, v)
		}
	}

	return cardsChangedDueDate
}

// if trello.Card.Due != modules.Card.Due => return false
func CompareTwoCard(cardBoard *trello.Card, cardDB models.Card) bool {
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
