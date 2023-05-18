package tpl

import (
	"embed"
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

//go:embed login/*
var tplFS embed.FS

func GetTplFS() embed.FS {
	return tplFS
}

func SetTemplate(e *gin.Engine) {
	funcMap := template.FuncMap{
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
	}
	if _, err := os.Stat("./tpl/login/redirect.tmpl"); os.IsNotExist(err) {
		panic("redirect.tmpl is not exist")
		return
	}
	// 2. config
	e.Delims("{{", "}}")

	/** drop
	e.SetFuncMap(funcMap)
	e.LoadHTMLGlob("./tpl/login/*")
	**/

	// 3. template with funcs
	tmpl := template.Must(template.New("").Funcs(funcMap).ParseFS(GetTplFS(), "login/*.tmpl"))
	e.SetHTMLTemplate(tmpl)

	// 3. curl -D- m:4500/tpl-raw/login/redirect.tmpl
	e.StaticFS("tpl-raw", http.FS(GetTplFS()))

}
