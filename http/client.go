package main

import (
	"fmt"
    _ "net/url"
	"net/http"
    "io/ioutil"
	_ "log"
	_ "reflect"
	"bytes"
    _"strings"
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
    //data := []byte(`{"hello": "world"}`)
    url := "https://httpbin.org/put?a=1"
	//req, _ := http.NewRequest("PUT", url, bytes.NewBuffer(data))
	req, _ := http.NewRequest("PUT", url, bytes.NewBuffer([]byte(`a=1&b=2`)))
	//req, _ := http.NewRequest("PUT", url, strings.NewReader("username=admin&password=pass"))
    fmt.Println(bytes.NewBuffer)
    //req, err := http.NewRequest("PUT", "http://httpbin.org/put", nil) //body io.Reader=nil
	req.Header.Set("Content-Type", "application/json") //覆盖
    resp, _:= client.Do(req)

    pn := fmt.Println
    body, _:= ioutil.ReadAll(resp.Body)
    resp.Body.Close()
    pn(string(body))
}
