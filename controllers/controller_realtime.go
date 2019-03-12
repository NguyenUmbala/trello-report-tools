package controllers

import "fmt"

func UpdateCardRealTime() {
	idBoard := conf.IDBoard

	for {
		cardsOnTrelloDB, err := DBTrelloCard.GetMany("board:" + idBoard)
		if err != nil {
			fmt.Println(err)
		}

		cardsOnDB, err := DBCard.GetAll()
		if err != nil {
			fmt.Println(err)
		}

		DBCard.UpdateCardsChangedDueDate(cardsOnTrelloDB, cardsOnDB)
	}
}
