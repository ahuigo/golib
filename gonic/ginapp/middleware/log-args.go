package middleware

import (
	"bytes"
	_ "fmt"
	"io/ioutil"
	"log"
	"net/http"
	_ "regexp"
	"time"

	"github.com/gin-gonic/gin"
)

func LogTime(c *gin.Context) {
	// begin time
	t := time.Now()
	// Set example variable
	c.Set("example", "12345")

	if false {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "hey", "status": http.StatusOK})
		c.Abort() //暂停冒泡
		return    //没有return 的话，gonic 会继续执行after request
	}

	// next  middleware
	c.Next()

	// after request
	latency := time.Since(t)

	// access the statusCode we are sending
	statusCode := c.Writer.Status()
	loginName := getLoginUserFromBody(c)
	log.Println(c.Request.RequestURI, statusCode, latency, loginName, "----------")
}

func getLoginUserFromBody(c *gin.Context) (name string) {
	// bind
	query := struct {
		Username string `json:"username"`
	}{}
	bindPre(c, &query)
	name = query.Username
	return
}

func bindPre(c *gin.Context, req interface{}) error {
	buf, _ := ioutil.ReadAll(c.Request.Body)
	defer func() {
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(buf))
	}()
	// revert
	c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(buf))
	return c.ShouldBind(&req)
}
