package routers

import (
	"TrelloReportTools/controllers"

	"github.com/gin-gonic/gin"
)

var Routers *gin.Engine

func init() {
	Routers = gin.Default()
}

func SetupRouters() {
	Routers.GET("/b/cards/review", controllers.GetAllCardReview)
	Routers.GET("/b/cards/changedue", controllers.GetAllCardChangeDue)
}

func RoutineUpdate() {
	controllers.UpdateCardRealTime()
}
