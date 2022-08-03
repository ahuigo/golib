package demo

import (
	"flag"
	"net/http"
	"testing"

	"github.com/elazarl/goproxy"
)

func TestHttpsProxy(t *testing.T) {
	proxy := goproxy.NewProxyHttpServer()
	proxy.Verbose = true
	var port = flag.String("port", "8080", "The port.")
	flag.Parse()
	t.Logf("proxy listen port:%s", *port)
	http.ListenAndServe(":"+*port, proxy)
	t.Log("proxy done")
}
