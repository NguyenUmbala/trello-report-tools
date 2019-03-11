package routers

import (
	"TrelloReportTools/controllers"

	"github.com/gin-gonic/gin"
)

var Routers *gin.Engine

func init() {
	Routers = gin.Default()
}

var Card controllers.Card
var UpdateCard controllers.UpdateCard

func SetupRouters() {
	Routers.GET("/b/cards/review", Card.GetAllCardReview)
	Routers.GET("/b/cards/changedue", Card.GetAllCardChangeDue)
}

func RoutineUpdate() {
	UpdateCard.UpdateCardRealTime()
}
