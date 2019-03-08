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
	Routers.GET("/b/cards/review/:id_board", controllers.GetAllCardReview)
	Routers.GET("/b/cards/changedue/:id_board", controllers.GetAllCardChangeDue)
}

func RoutineUpdate() {
	controllers.UpdateBoardRealTime()
}
