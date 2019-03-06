package handles

import (
	"TrelloReportTools/database"
	"TrelloReportTools/modules"
	"TrelloReportTools/trelloAPI"

	"github.com/adlio/trello"
)

func UpdateCardsInRealTime() {
	for {
		idBoard := "iCBtQXmr"
		cardsOnBoard, err := trelloAPI.GetCardsOnBoard(idBoard)
		if err != nil {
			// Handle error
		}

		cardsOnDB := database.GetCards()
		var cardsChangedDueDate []*trello.Card
		var tmpCard modules.Card
		for i, _ := range cardsOnBoard {
			for j, _ := range cardsOnDB {
				if CompareTwoCard(cardsOnBoard[i], cardsOnDB[j]) {
					cardsChangedDueDate = append(cardsChangedDueDate, cardsOnBoard[i])
					database.UpdateCard(tmpCard.NewCard(cardsOnBoard[i]))
				}
			}
		}
	}
}
