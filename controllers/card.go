package controllers

import (
	"TrelloReportTools/database"
	"TrelloReportTools/modules"
	"TrelloReportTools/trelloAPI"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func init() {
	// Save cards
	var tmpCard modules.Card
	idBoard := "iCBtQXmr"
	cardsOnBoard, err := trelloAPI.GetCardsOnBoard(idBoard)
	if err != nil {
		fmt.Println(err)
	}

	for _, v := range cardsOnBoard {
		database.SaveCard(tmpCard.NewCard(v))
	}
}

func GetAllCardReview(c *gin.Context) {
	idBoard := c.Param("id_board")

	cardsOnListReviewMe, err := trelloAPI.GetCardsIsOpenOnWeek(idBoard, "review-me")
	if err != nil {
		fmt.Println(err)
	}
	cardsOnListDone, err := trelloAPI.GetCardsIsOpenOnWeek(idBoard, "Done")
	if err != nil {
		fmt.Println(err)
	}

	c.JSON(200, gin.H{
		"List card on review-me": cardsOnListReviewMe,
		"List card on Done":      cardsOnListDone,
	})

}

func GetAllCardChangeDue(c *gin.Context) {
	lastDay := time.Now().AddDate(0, 0, -1)
	cardsOnDB := database.GetCards()
	var cardsChangedDueDate []modules.Card

	for _, v := range cardsOnDB {
		date := *v.DateLastChangeDue
		if date.After(lastDay) {
			cardsChangedDueDate = append(cardsChangedDueDate, v)
		}
	}

	c.JSON(200, gin.H{
		"Cards changed due date": cardsChangedDueDate,
	})
}
