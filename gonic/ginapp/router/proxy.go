package router

import (
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/gin-gonic/gin"
)

func ProxyServer(c *gin.Context) {
	proxy := &httputil.ReverseProxy{
		Director: func(r *http.Request) {
			r.Header.Set("X-Forwarded-Host", r.Header.Get("Host"))
			r.URL.Scheme = "https"
			r.URL.Host = "httpbin.org"
			r.Host = "httpbin.org"
		},
		ModifyResponse: func(r *http.Response) error {
			r.Header.Del("Access-Control-Allow-Origin")
			r.Header.Del("Access-Control-Allow-Credentials")
			r.Header.Del("Access-Control-Allow-Headers")
			r.Header.Del("Access-Control-Allow-Methods")
			return nil
		},
	}
	if false {
		proxy = httputil.NewSingleHostReverseProxy(&url.URL{Scheme: "https", Host: "httpbin.org"})
	}
	proxy.ServeHTTP(c.Writer, c.Request)
}
