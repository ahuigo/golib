package httpx

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestHttpFileServer(t *testing.T) {

	URL, _ := url.Parse("/index.html")
	fs := http.Dir("./")
	respRecorder := httptest.NewRecorder()
	req := &http.Request{
		Method:     "GET",
		Proto:      "HTTP/1.1",
		ProtoMajor: 1,
		ProtoMinor: 1,
		URL:        URL,
	}

	fileServer := http.FileServer(fs)
	fileServer.ServeHTTP(respRecorder, req)
	if respRecorder.Code == 301 {
		t.Fatal("Expect code 200, result code: ", respRecorder.Code, "reason:", respRecorder.Body)
	}

}
