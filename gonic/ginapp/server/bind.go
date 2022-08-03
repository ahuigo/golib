package server

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

// curl -H 'Content-Type: application/json' 'm:4500/bind/f' -d '{"extra":{"name":"ahui", "age":1}}'
func BindServer(c *gin.Context) {
	//backup
	buf, _ := ioutil.ReadAll(c.Request.Body)
	// revert main
	c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(buf)) // important!!

	// bind
	user := User{}
	if err := c.ShouldBind(&user); err != nil {
		fmt.Println("bind error:", err)
	}
	fmt.Printf("user:%#v \n", user)

	// c.String(http.StatusOK, res)
	c.JSON(http.StatusOK, user)

}
