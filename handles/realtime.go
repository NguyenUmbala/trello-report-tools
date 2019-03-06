package handles

import (
	"TrelloReportTools/modules"
	"TrelloReportTools/trelloAPI"

	"github.com/adlio/trello"
)

func UpdateCardsInRealTime() {
	for {
		idBoard := "iCBtQXmr"
		cardsOnBoard := trelloAPI.GetCardsOnBoard(idBoard)
		cardsOnDB := modules.GetCards()

		var cardsChangedDueDate []*trello.Card
		var tmpCard modules.Card
		for i, _ := range cardsOnBoard {
			for j, _ := range cardsOnDB {
				if CompareTwoCard(cardsOnBoard[i], cardsOnDB[j]) {
					cardsChangedDueDate = append(cardsChangedDueDate, cardsOnBoard[i])
					modules.UpdateCard(tmpCard.NewCard(cardsOnBoard[i]))
				}
			}
		}
	}
}
