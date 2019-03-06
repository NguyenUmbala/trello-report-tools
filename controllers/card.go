package controllers

import (
	"TrelloReportTools/modules"
	"TrelloReportTools/trelloAPI"
	"TrelloReportTools/utils"
	"time"

	"github.com/gin-gonic/gin"
)

var Card utils.Card

func init() {
	// Save cards
	idBoard := "iCBtQXmr"
	cardsOnBoard := trelloAPI.GetCardsOnBoard(idBoard)

	for _, v := range cardsOnBoard {
		modules.UpdateCard(Card.NewCard(v))
	}
}

func GetAllCardReview(c *gin.Context) {
	idBoard := c.Param("id_board")

	cardsOnListReviewMe := trelloAPI.GetCardsIsOpenOnWeek(idBoard, "review-me")
	cardsOnListDone := trelloAPI.GetCardsIsOpenOnWeek(idBoard, "Done")

	c.JSON(200, gin.H{
		"List card on review-me": cardsOnListReviewMe,
		"List card on Done":      cardsOnListDone,
	})

}

func GetAllCardChangeDue(c *gin.Context) {
	cardsOnDB := modules.GetCards()
	var cardsChangedDueDate []modules.Card

	lastDay := time.Now().AddDate(0, 0, -1)
	for _, v := range cardsOnDB {
		if lastDay.Before(*v.DateLastChangeDue) {
			cardsChangedDueDate = append(cardsChangedDueDate, v)
		}
	}

	c.JSON(200, gin.H{
		"Cards changed due date": cardsChangedDueDate,
	})
}
