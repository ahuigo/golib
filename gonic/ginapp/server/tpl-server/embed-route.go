package tpls

import (
	"encoding/json"
	"fmt"
	"ginapp/tpl"
	"html/template"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

func TplRouter(e *gin.Engine) {
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
	}
	// 2. config
	e.Delims("{{", "}}")

	/** drop
	e.SetFuncMap(funcMap)
	e.LoadHTMLGlob("./tpl/login/*")
	**/

	// 3. template with funcs
	tmplWithFuncs := template.Must(template.New("").Funcs(funcMap).ParseFS(tpl.GetLoginFS(), "login/*.tmpl"))
	e.SetHTMLTemplate(tmplWithFuncs)

	// 3. curl -D- m:4500/tpl-raw/login/redirect.tmpl
	e.StaticFS("tpl-raw", http.FS(tpl.GetLoginFS()))
	// 4. redirect
	e.POST("tpl/redirect", TplRedirectPage)
}
