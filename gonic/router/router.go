package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.Any("/v1/*path", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
            "all_path":    c.Request.URL.Path,
            "sub_path": c.Param("path"),
		})
	})
    r.Run(":8082") // listen and serve on 0.0.0.0:8080
}
