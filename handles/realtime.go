package handles

import (
	"TrelloReportTools/modules"
	"TrelloReportTools/trelloAPI"
	"TrelloReportTools/utils"
)

var Card utils.Card

func UpdateCardsInRealTime() {
	for {
		idBoard := "iCBtQXmr"
		cardsOnBoard := trelloAPI.GetCardsOnBoard(idBoard)
		cardsOnDB := modules.GetCards()

		// Compare due date of two cards
		for i, _ := range cardsOnBoard {
			for j, _ := range cardsOnDB {
				if Card.CompareTwoCard(cardsOnBoard[i], cardsOnDB[j]) {
					modules.UpdateCard(Card.NewCard(cardsOnBoard[i]))
				}
			}
		}
	}
}
