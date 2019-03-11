package services

import (
	"TrelloReportTools/models"
	"TrelloReportTools/utils"
)

type ServiceUpdate struct{}

var UtilString utils.UtilString
var UtilTime utils.UtilTime

func (*ServiceUpdate) UpdateCards(idBoard string) {
	cardsOnBoard, _ := dbTrello.SelectManyTrello("board:" + idBoard)
	cardsOnDB, _ := db.SelectAll()
	var timeGuess, timeReal int

	// Compare due date of two cards
	for _, cardBoard := range cardsOnBoard {
		timeGuess = UtilString.GetTimeGuessForDone(cardBoard.Name)
		timeReal = UtilString.GetRealTimeOfDone(cardBoard.Name)

		for j, cardDB := range cardsOnDB {
			if cardBoard.ID == cardDB.ID {
				if UtilTime.CompareTime(cardBoard.Due, cardDB.Due) == false {
					db.InsertOrUpdate(models.NewCard(cardBoard, timeGuess, timeReal))
				}
				break
			}

			// If this card is not in database
			if j == len(cardsOnDB)-1 {
				db.InsertOrUpdate(models.NewCard(cardBoard, timeGuess, timeReal))
			}
		}
	}
}

func (*ServiceUpdate) SaveAllCards(idboard string) {
	cardsOnBoard, _ := dbTrello.SelectManyTrello("board:" + idboard)

	for _, v := range cardsOnBoard {
		timeGuess := UtilString.GetTimeGuessForDone(v.Name)
		timeReal := UtilString.GetRealTimeOfDone(v.Name)

		db.InsertOrUpdate(models.NewCard(v, timeGuess, timeReal))
	}
}
