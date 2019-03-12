package controllers

import (
	"TrelloReportTools/config"
	"TrelloReportTools/services"

	"github.com/gin-gonic/gin"
)

type Card struct{}

var conf = config.Setup()
var ServiceCard services.Card

func init() {
	cardsOnTrelloDB := ServiceCard.GetCardsOnBoard(conf.IDBoard)
	cardsOnDB := ServiceCard.ConvertCardsTrelloDBToDB(cardsOnTrelloDB)
	ServiceCard.SaveCards(cardsOnDB)
}

// Review all card is open and created on week
func (*Card) GetAllCardReview(c *gin.Context) {
	idBoard := conf.IDBoard

	c.JSON(200, gin.H{
		"List card on review-me": ServiceCard.GetCardsIsOpenOnWeek(idBoard, "review-me"),
		"List card on Done":      ServiceCard.GetCardsIsOpenOnWeek(idBoard, "Done"),
	})
}

// Review all card changed due date by last day
func (*Card) GetAllCardChangeDue(c *gin.Context) {
	dayNumber := 1 // Day number before check day

	c.JSON(200, gin.H{
		"Cards changed due date": ServiceCard.GetCardsChangedDueByTime(dayNumber),
	})
}
