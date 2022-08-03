package main

import (
	"fmt"
	"html/template"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func formatAsDate(t time.Time) string {
	year, month, day := t.Date()
	return fmt.Sprintf("%d%02d/%02d", year, month, day)
}

func main() {
	router := gin.Default()
	router.Delims("{{", "}}")
	router.SetFuncMap(template.FuncMap{
		"formatAsDate": formatAsDate,
	})
    router.LoadHTMLGlob("./tpl/*")
	//router.LoadHTMLFiles("templates/template1.html", "templates/template2.html")
	//router.LoadHTMLFiles("./testdata/raw.tmpl")

	router.GET("/raw", func(c *gin.Context) {
		c.HTML(http.StatusOK, "raw.tmpl", map[string]interface{}{
			"now": time.Date(2017, 07, 01, 0, 0, 0, 0, time.UTC),
			"now1": time.Date(2017, 07, 01, 0, 0, 0, 0, time.UTC),
			"name": "ahuigo",
			"count": 1,
            "map":map[string]string{
                "key1":"v1",
            },
		})
	})

	router.Run(":8080")
}
