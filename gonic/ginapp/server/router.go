package server

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	mid "ginapp/middleware"

	"github.com/DeanThompson/ginpprof"
	"github.com/gin-gonic/gin"
)

func Register(r *gin.Engine, staticFS bool) {
	if staticFS {
		staticFsHandler(r)
		return
	}
	// automatically add routers for net/http/pprof
	ginpprof.Wrap(r)
	// swagger
	RegisterSwagger(r)
	// midddleware
	r.Use(mid.LogTime, mid.CORS, mid.Error)
	// staticFS
	r.Static("/js", "./js")
	// template
	setTemplate(r)

	// router
	r.GET("/gorm/insert", insertHandler)
	r.GET("/f", fileReadHandler)
	r.GET("/api/v1/*path", fileReadHandler)
	r.GET("/api/panic", panicApi)
	r.GET("/dump/*anypath", DumpServer)
	r.POST("/dump/*anypath", DumpServer)
	r.GET("/redirect/:code", RedirectServer)
	r.POST("/redirect/form", RedirectServer)
	r.GET("/echo/:size", EchoServer)
	r.GET("/bind/*anypath", BindServer)
	r.POST("/bind/*anypath", BindServer)
	r.GET("/sleep/:second", sleepFunc)
	r.GET("/cpu/:second", cpuFunc)
	r.GET("/json/map", jsonMapFunc)
	r.GET("/stream", streamApi)
	// r.Any("/bind/*anypath", BindServer)
}

func setTemplate(e *gin.Engine) {
	if _, err := os.Stat("./tpl"); os.IsNotExist(err) {
		return
	}
	e.Delims("{{", "}}")
	e.LoadHTMLGlob("./tpl/*")
	e.SetFuncMap(template.FuncMap{
		"formatAsDate": func(t time.Time) string {
			year, month, day := t.Date()
			return fmt.Sprintf("%d%02d/%02d", year, month, day)
		},
		"jsonEncode": func(obj interface{}) string {
			if d, err := json.Marshal(obj); err == nil {
				return string(d)
			} else {
				return err.Error()
			}
		},
	})
}

func jsonMapFunc(c *gin.Context) {
	m := map[string][]byte{
		"status": []byte("running!中国"),
	}
	c.JSON(http.StatusOK, m)
}

func EchoServer(c *gin.Context) {
	size, _ := strconv.Atoi(c.Param("size"))
	if size > 1e5 {
		size = 1e5
	}
	msg := strings.Repeat("中", size)
	c.String(http.StatusOK, msg)
}

func cpuFunc(c *gin.Context) {
	seconds, _ := strconv.Atoi(c.Param("second"))
	n := longRun(seconds)
	msg := fmt.Sprintf("sleep second: %v s, n=%d\n", seconds, n)
	c.JSON(http.StatusOK, msg)
}

func longRun(seconds int) int {
	now := time.Now()
	end_time := now.Add(time.Duration(seconds) * time.Second)
	n := 0
	for ; end_time.After(now); now = time.Now() {
		for i := 0; i < 1e8; i += 1 {
			n += 1
		}
	}
	return n
}

func sleepFunc(c *gin.Context) {
	seconds, _ := strconv.Atoi(c.Param("second"))
	time.Sleep(time.Duration(seconds) * time.Second)
	fmt.Printf("%vs passed!\n", seconds)
	c.JSON(http.StatusOK, "sleep second: "+c.Param("second"))
}
