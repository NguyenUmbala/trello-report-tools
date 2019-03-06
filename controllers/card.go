package controllers

import (
	"TrelloReportTools/database"
	"TrelloReportTools/modules"
	"TrelloReportTools/trelloAPI"
	"time"

	"github.com/adlio/trello"
	"github.com/gin-gonic/gin"
)

func init() {
	// Save cards
	var tmpCard modules.Card
	idBoard := "iCBtQXmr"
	cardsOnBoard, err := trelloAPI.GetCardsOnBoard(idBoard)
	if err != nil {
		// Handle error
	}

	for _, v := range cardsOnBoard {
		database.SaveCard(tmpCard.NewCard(v))
	}
}

func GetAllCardReview(c *gin.Context) {
	idBoard := c.Param("id_board")

	cardsOnListReviewMe, err := trelloAPI.GetCardsIsOpenOnWeek(idBoard, "review-me")
	if err != nil {
		// Handle error
	}
	cardsOnListDone, err := trelloAPI.GetCardsIsOpenOnWeek(idBoard, "Done")
	if err != nil {
		// Handle error
	}

	c.JSON(200, gin.H{
		"List card on review-me": cardsOnListReviewMe,
		"List card on Done":      cardsOnListDone,
	})

}

func GetAllCardChangeDue(c *gin.Context) {
	cardsOnDB := database.GetCards()
	var cardsChangedDueDate []modules.Card
	lastDay := time.Now().AddDate(0, 0, -1)
	for _, v := range cardsOnDB {
		date := *v.DateLastActivity
		if date.After(lastDay) {
			cardsChangedDueDate = append(cardsChangedDueDate, v)
		}
	}

	c.JSON(200, gin.H{
		"Cards changed due date": cardsChangedDueDate,
	})
}

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

// if trello.Card != modules.Card => return true
func CompareTwoCard(cardsBoard *trello.Card, cardsDB modules.Card) bool {
	if cardsBoard.ID != cardsDB.ID {
		return false
	}
	if cardsBoard.Due == nil && cardsDB.Due == nil {
		return false
	}
	if cardsBoard.Due == nil && cardsDB.Due != nil {
		return true
	}
	if cardsBoard.Due != nil && cardsDB.Due == nil {
		return true
	}
	if cardsBoard.Due.String() == cardsDB.Due.String() {
		return false
	}
	return true
}
