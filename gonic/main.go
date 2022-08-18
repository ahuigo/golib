package main

import (
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
        var query struct{
        	Debug bool `json:"debug" form:"debug" gorm:"-"`
        }
        c.ShouldBindQuery(query)
        time.Sleep(3000 * time.Millisecond)
        println("hello")

		c.JSON(200, gin.H{
			"message": "pong",
            "debug":query.Debug,
		})
	})
    r.Run(":8082") // listen and serve on 0.0.0.0:8080
}
