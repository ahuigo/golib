package main

import (
	"fmt"
    _ "net/url"
	"net/http"
    "io/ioutil"
	_ "log"
	_ "reflect"
	_ "bytes"
    "strings"
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
    url := "http://127.0.0.1:5000/header.php"
    body:=`{"a":1}`
    req, _ := http.NewRequest("POST", url, strings.NewReader(body))
    req.Header.Set("Content-Type","application/json")
    resp, _ := client.Do(req)
    //ioutil.ReadAll(resp.Body)


    req, _ = http.NewRequest("POST", url, strings.NewReader(body))
    req.Header.Set("Content-Type","application/json")
    resp, _ = client.Do(req)
    for k, v := range resp.Header{
		fmt.Printf("k=%v, v=%v\n", k, v)
	}
    for i, cookie := range resp.Cookies() {
        fmt.Println("Found a cookie:",i, cookie.Name,"=", cookie.Value)
    }
    pn := fmt.Println
    bodyByte, _:= ioutil.ReadAll(resp.Body)
    pn(string(bodyByte))
    location, _ :=resp.Location()
    fmt.Println(location)
}
