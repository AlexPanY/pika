package main

import (
	"fmt"
	"pika/driver"
	"pika/routers"

	"github.com/gin-gonic/gin"
)

func main() {

	driver.InitDB()

	g := gin.Default()

	g = routers.Load(g)
	// ginpprof.Wrap(g)

	if err := g.Run(":8082"); err != nil {
		fmt.Println(err)
	}
}
