package routers

import (
	"TrelloReportTools/controllers"

	"github.com/gin-gonic/gin"
)

var Routers *gin.Engine

func init() {
	Routers = gin.Default()
}

var GetCard controllers.GetCard
var UpdateCard controllers.UpdateCard

func SetupRouters() {
	Routers.GET("/b/cards/review", GetCard.GetAllCardReview)
	Routers.GET("/b/cards/changedue", GetCard.GetAllCardChangeDue)
}

func RoutineUpdate() {
	UpdateCard.UpdateCardRealTime()
}
