package main

import (
	"TrelloReportTools/handles"
	"TrelloReportTools/routers"
	"sync"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.DebugMode)
	r := routers.Routers
	routers.SetupRouters()

	var wg sync.WaitGroup
	wg.Add(1)
	go handles.UpdateCardsInRealTime()
	r.Run(":3000")
	wg.Wait()
}
