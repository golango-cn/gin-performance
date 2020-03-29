package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"golango.cn/gin-performance/core"
	"log"
	"sync/atomic"
)

var count int32

func InitRouter() *gin.Engine {

	core.Init()

	r := gin.New()

	r.GET("/", func(c *gin.Context) {

		employees, err := core.GetEmployees()
		if err != nil {
			log.Println(err)
		} else {
			atomic.AddInt32(&count, 1)
			fmt.Println("count", count)
		}
		c.JSON(200, employees)
	})

	return r

}
