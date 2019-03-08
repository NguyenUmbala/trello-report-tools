package services

import (
	"TrelloReportTools/models"
)

func UpdateRealtime(idBoard string) {
	for {
		cardsOnBoard := GetCardsOnBoard(idBoard)
		cardsOnDB := models.GetAllCard()
		var timeGuess, timeReal int

		// Compare due date of two cards
		for i, _ := range cardsOnBoard {
			timeGuess = GetTimeGuessForDone(cardsOnBoard[i].Name)
			timeReal = GetRealTimeOfDone(cardsOnBoard[i].Name)

			for j, _ := range cardsOnDB {
				if cardsOnBoard[i].ID == cardsOnDB[j].ID {
					if CompareTwoCard(cardsOnBoard[i], cardsOnDB[j]) {
						models.UpdateCard(cardsOnDB[j].NewCard(cardsOnBoard[i], timeGuess, timeReal))
					}
					break
				}

				// If this card is not in database
				if j == len(cardsOnDB)-1 {
					var tmpCard models.Card
					models.UpdateCard(tmpCard.NewCard(cardsOnBoard[i], timeGuess, timeReal))
				}
			}
		}

	}
}
