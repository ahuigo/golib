package main

import (
	"fmt"
    _ "net/url"
	"net/http"
    "io/ioutil"
	_ "log"
	_ "reflect"
	_ "bytes"
)

func main() {
    //resp, _ := http.PostForm("https://httpbin.org/cookies/set?ahui=mo",
    client := &http.Client{
        CheckRedirect: func(req *http.Request, via []*http.Request) error {
            return http.ErrUseLastResponse
        },
    }
    resp, _ := client.Get("https://httpbin.org/cookies/set?ahui=mo")
    for k, v := range resp.Header{
		fmt.Printf("k=%v, v=%v\n", k, v)
	}
    for i, cookie := range resp.Cookies() {
        fmt.Println("Found a cookie:",i, cookie.Name,"=", cookie.Value)
    }
    pn := fmt.Println
    body, _:= ioutil.ReadAll(resp.Body)
    pn(string(body))
    location, _ :=resp.Location()
    fmt.Println(location)
}
