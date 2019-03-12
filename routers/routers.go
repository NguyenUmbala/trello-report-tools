package routers

import (
	"TrelloReportTools/controllers"

	"github.com/gin-gonic/gin"
)

var Routers *gin.Engine

func init() {
	Routers = gin.Default()
}

var ControllerCard controllers.Card
var UpdateCard controllers.UpdateCard

func SetupRouters() {
	Routers.GET("/b/cards/review", ControllerCard.GetAllCardReview)
	Routers.GET("/b/cards/changedue", ControllerCard.GetAllCardChangeDue)
}

func RoutineUpdate() {
	UpdateCard.UpdateCardRealTime()
}
