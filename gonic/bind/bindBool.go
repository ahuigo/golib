package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
        var query struct{
            // true 的情况只能是1,true,True
        	Debug bool `json:"debug" form:"debug" gorm:"-"`
        }
        c.ShouldBindQuery(&query)

		c.JSON(200, gin.H{
			"message": "pong",
            "debug":query.Debug,
		})
	})
    r.Run(":8082") // listen and serve on 0.0.0.0:8080
}
