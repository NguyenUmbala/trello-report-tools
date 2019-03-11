package controllers

import (
	"TrelloReportTools/config"
	"TrelloReportTools/services"

	"github.com/gin-gonic/gin"
)

type Card struct{}

var conf = config.Setup()
var Service services.Service
var ServiceUpdate services.ServiceUpdate

func (*Card) GetAllCardReview(c *gin.Context) {
	idBoard := conf.IDBoard

	cardsOnListReviewMe := Service.GetCardsIsOpenOnWeek(idBoard, "open", "created", "week", "review-me")
	cardsOnListDone := Service.GetCardsIsOpenOnWeek(idBoard, "open", "created", "week", "Done")

	c.JSON(200, gin.H{
		"List card on review-me": cardsOnListReviewMe,
		"List card on Done":      cardsOnListDone,
	})
}

func (*Card) GetAllCardChangeDue(c *gin.Context) {
	dayNumber := 1 // Day number before check day
	cardsChangedDueDate := Service.GetCardsChangedDueByTime(dayNumber)

	c.JSON(200, gin.H{
		"Cards changed due date": cardsChangedDueDate,
	})
}
