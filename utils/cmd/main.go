package main

import (
	"pika/driver"
	"pika/pkg/logger"
	"pika/routers"

	"github.com/gin-gonic/gin"
)

func main() {

	driver.InitDB()
	driver.InitRedis()

	logger.InitLogger()

	gin.SetMode(gin.ReleaseMode)

	g := gin.New()

	g = routers.Load(g)
	// ginpprof.Wrap(g)

	if err := g.Run(":8082"); err != nil {
		logger.F.Error("service start fail: " + err.Error())
	}
}
