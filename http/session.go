package main

import (
	"fmt"
    _ "net/url"
	"net/http"
    "io/ioutil"
	_ "log"
	_ "reflect"
	_ "bytes"
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

    resp, _ := client.Get("http://127.0.0.1:5000/header.php")
    resp, _ = client.Get("http://bing.com")
    //ioutil.ReadAll(resp.Body)
    resp, _ = client.Get("http://127.0.0.1:5000/header.php")
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
