package main

import (
	"sync"

	"trello-report-tools/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.DebugMode)
	r := routers.SetupRouters()

	var wg sync.WaitGroup
	wg.Add(1)

	go routers.RoutineUpdate()
	r.Run(":8081")
	wg.Wait()
}
