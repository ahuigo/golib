package server

import (
	"io"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

// EchoHandler _
func fileReadHandler(c *gin.Context) {
	filepath := c.DefaultQuery("path", "tmp/a.txt")
	filepath = c.DefaultQuery("path", "./new.json")
	if strings.HasSuffix(filepath, ".json") {
		c.Writer.Header().Set("Content-type", "application/json")
	} else {
		c.Writer.Header().Set("Content-type", "text/html; charset=utf-8")
	}

	// c.String(http.StatusOK, res)

	f, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}
	io.Copy(c.Writer, f)

}
