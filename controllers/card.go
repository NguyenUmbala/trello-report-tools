package controllers

import (
	"TrelloReportTools/services"

	"github.com/gin-gonic/gin"
)

func GetAllCardReview(c *gin.Context) {
	idBoard := c.Param("id_board")

	cardsOnListReviewMe := services.GetCardsIsOpenOnWeek(idBoard, "review-me")
	cardsOnListDone := services.GetCardsIsOpenOnWeek(idBoard, "Done")

	c.JSON(200, gin.H{
		"List card on review-me": cardsOnListReviewMe,
		"List card on Done":      cardsOnListDone,
	})

}

func GetAllCardChangeDue(c *gin.Context) {
	dayNumber := 1 // Day number before check day
	cardsChangedDueDate := services.GetCardsChangedDueInTime(dayNumber)

	c.JSON(200, gin.H{
		"Cards changed due date": cardsChangedDueDate,
	})
}

func UpdateBoardRealTime() {
	idBoard := "iCBtQXmr"
	services.UpdateRealtime(idBoard)
}
