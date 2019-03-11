package controllers

import (
	"TrelloReportTools/config"
	"TrelloReportTools/services"

	"github.com/gin-gonic/gin"
)

type GetCard struct{}

var conf = config.Setup()
var ServiceGet services.ServiceGet
var ServiceUpdate services.ServiceUpdate

func init() {
	cardsOnBoard := ServiceGet.GetCardsOnBoard(conf.IDBoard)
	for _, value := range cardsOnBoard {
		ServiceUpdate.UpdateCard(value)
	}
}

func (*GetCard) GetAllCardReview(c *gin.Context) {
	idBoard := conf.IDBoard

	cardsOnListReviewMe := ServiceGet.GetCardsIsOpenOnWeek(idBoard, "open", "created", "week", "review-me")
	cardsOnListDone := ServiceGet.GetCardsIsOpenOnWeek(idBoard, "open", "created", "week", "Done")

	c.JSON(200, gin.H{
		"List card on review-me": cardsOnListReviewMe,
		"List card on Done":      cardsOnListDone,
	})
}

func (*GetCard) GetAllCardChangeDue(c *gin.Context) {
	dayNumber := 1 // Day number before check day
	cardsChangedDueDate := ServiceGet.GetCardsChangedDueByTime(dayNumber)

	c.JSON(200, gin.H{
		"Cards changed due date": cardsChangedDueDate,
	})
}
