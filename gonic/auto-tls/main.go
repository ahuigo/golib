// refer1: https://stackoverflow.com/questions/67625752/how-to-use-a-certificate-from-a-certificate-store-and-run-tls-in-gin-framework-i
// refer2: creat cert via requests/conf/readme.md
package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	g := gin.Default()
	g.GET("/hello/:name", func(c *gin.Context) {
		c.String(200, "Hello %s", c.Param("name"))
	})
	g.RunTLS(":3000", "./certs/server.crt", "./certs/server.key")
}
