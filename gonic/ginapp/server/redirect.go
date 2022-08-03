package server

import (
	"io/ioutil"
	"net/http"
	URL "net/url"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/render"
)

// EchoHandler _
func RedirectServer(ctx *gin.Context) {
	redirectUri := ctx.Query("redirect_uri")
	switch ctx.Param("code") {
	case "302":
		ctx.Redirect(http.StatusFound, redirectUri)
	case "307":
		ctx.Redirect(http.StatusTemporaryRedirect, redirectUri)
	case "308":
		ctx.Redirect(http.StatusPermanentRedirect, redirectUri)
	case "refresh":
		redirectRefresh(ctx)
	case "form":
		fallthrough
	default:
		redirectForm(ctx)
	}
}

// RedirectHandler
func redirectRefresh(ctx *gin.Context) {
	url := encodeURI(ctx.Query("redirect_uri"))
	html := `<head>
	<meta http-equiv="refresh"  content="1; url=` + url + `">
	</head>`
	ctx.Render(http.StatusOK, render.Data{
		ContentType: "text/html",
		Data:        []byte(html),
	})
}

// requestForm('post','http://m:4500/redirect/form?redirect_uri=https://s/dump/ab/c/c2',{id_token:"xx",timeout:11,redirect_uri:"https://s/dump/ab/c/c"})
func redirectForm(ctx *gin.Context) {
	url := encodeURI(ctx.Query("redirect_uri"))
	if url == "" {
		url = "https://s/dump/ab/c/c2"
	}
	body, _ := ioutil.ReadAll(ctx.Request.Body)
	query, _ := URL.ParseQuery(string(body))
	data := map[string]string{}
	for k, v := range query {
		data[k] = v[0]
	}
	rdata := map[string]interface{}{
		"url":  url,
		"body": data,
	}

	ctx.HTML(http.StatusOK, "redirect.tmpl", rdata)

}

func encodeURI(url string) string {
	u, _ := URL.Parse(url)
	u.RawQuery = u.Query().Encode()
	return u.String()
}
