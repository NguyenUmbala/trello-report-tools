package controllers

import (
	"trello-report-tools/config"
	"trello-report-tools/models"

	"github.com/gin-gonic/gin"
)

var conf = config.Setup()
var DBCard models.DBCard
var DBTrelloCard models.DBTrelloCard

// Review all card is open and created on week
func GetAllCardReview(c *gin.Context) {
	idBoard := conf.IDBoard

	listCardOnReviewme, _ := DBTrelloCard.GetMany("board:" + idBoard + " is:open sort:created created:week list:review-me")
	listCardOnDone, _ := DBTrelloCard.GetMany("board:" + idBoard + " is:open sort:created created:week list:Done")

	c.JSON(200, gin.H{
		"List card on review-me": listCardOnReviewme,
		"List card on Done":      listCardOnDone,
	})
}

// Review all card changed due date by last day
func GetAllCardChangeDue(c *gin.Context) {
	dayNumber := 1 // Day number before check day

	c.JSON(200, gin.H{
		"Cards changed due date": DBCard.GetCardsChangedDueByTime(dayNumber),
	})
}
