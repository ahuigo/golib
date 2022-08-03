package main

import (
	"bytes"
	"fmt"

	"github.com/gin-gonic/gin"
)

/********** responseBodyWriter ******************/
type responseBodyWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer //cache
}

func (r responseBodyWriter) Write(b []byte) (int, error) {
	r.body.Write(b)
	return r.ResponseWriter.Write(b)
}

func (r responseBodyWriter) WriteString(s string) (n int, err error)  {
	r.body.WriteString(s)
	return r.ResponseWriter.WriteString(s)
}



/********** replace responseBodyWriter ******************/
func logResponseBody(c *gin.Context) {
	w := &responseBodyWriter{body: &bytes.Buffer{}, ResponseWriter: c.Writer}
	c.Writer = w
    // rewrite header
    w.Header().Set("Server", "ahui server")

    // rewrite code
    w.ResponseWriter.WriteHeader(205)
    w.WriteHeaderNow() // call before c.Next

    // do...
	c.Next()
    body := w.body.String()

    //rewrite body
    w.WriteString("\n")
    header := fmt.Sprintf("dumpheader\n:%#v\n",w.Header())
    w.WriteString(header)
    w.WriteString("dumpbody:\n"+body+"\n")
	fmt.Println("Response body: " + w.body.String())

}

func sayHello(c *gin.Context) {
	c.JSON(200, gin.H{
		"hello": "privationel",
	})
}

func main() {
	r := gin.Default()
	r.Use(logResponseBody)
	r.GET("/", sayHello)
	r.Run()
}
