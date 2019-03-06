package main

import (
	"TrelloReportTools/handles"
	"TrelloReportTools/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.DebugMode)
	r := routers.Routers

	routers.SetupRouters()
	go handles.UpdateCardsInRealTime()
	r.Run(":3000")
}
