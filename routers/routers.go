package routers

import (
	"trello-report-tools/controllers"

	"github.com/gin-gonic/gin"
)

func init() {
}

func SetupRouters() *gin.Engine {
	var Routers *gin.Engine
	Routers = gin.Default()

	Routers.GET("/b/cards/review", controllers.GetAllCardReview)
	Routers.GET("/b/cards/changedue", controllers.GetAllCardChangeDue)

	return Routers
}

func RoutineUpdate() {
	controllers.UpdateCardRealTime()
}
