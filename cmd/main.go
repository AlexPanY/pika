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

	routers.Load(g)

	if err := g.Run(":8080"); err != nil {
		fmt.Println(err)
	}
}
