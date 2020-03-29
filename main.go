package main

import (
	"github.com/gin-gonic/gin"
	"golango.cn/gin-performance/router"
	"log"
)

func main()  {

	gin.SetMode(gin.DebugMode)
	router := router.InitRouter()

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}

}
