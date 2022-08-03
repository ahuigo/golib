package main

import (
	"fmt"
    _ "net/url"
	"net/http"
    "io/ioutil"
	_ "log"
	_ "reflect"
	"bytes"
    "net/http/cookiejar"
)
func main() {
    cookieJar, _ := cookiejar.New(nil)
    client := &http.Client{
        Jar: cookieJar,
        CheckRedirect: func(req *http.Request, via []*http.Request) error {
            return http.ErrUseLastResponse
        },
    }
    data := []byte(`{"hello": "world"}`)
    url := "https://httpbin.org/put?a=1"
	//work req, _ := http.NewRequest("PUT", url, bytes.NewReader(data))
	req, _ := http.NewRequest("PUT", url, bytes.NewBuffer(data))
    fmt.Println(bytes.NewBuffer, data)
    //req, err := http.NewRequest("PUT", "http://httpbin.org/put", nil) //body io.Reader=nil
	req.Header.Set("Content-Type", "application/json") //覆盖
    resp, _:= client.Do(req)

    pn := fmt.Println
    body, _:= ioutil.ReadAll(resp.Body)
    resp.Body.Close()
    pn(string(body))
}
