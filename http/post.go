package main

import (
	"fmt"
    "net/url"
    "strings"
	"net/http"
    "io/ioutil"
	_ "log"
	_ "reflect"
	_ "bytes"
)

func main() {
    //resp, _ := http.PostForm("https://httpbin.org/cookies/set?ahui=mo",
    resp, _ := http.PostForm("https://httpbin.org/post", url.Values{"key": {"Value"}, "id": {"123"}})
    //resp, _ := http.PostForm("https://httpbin.org/post", strings.NewReader("username=admin&password=test"))
    resp, _ = http.Post("http://127.0.0.1:5000/header.php", "application/x-www-form-urlencoded", strings.NewReader("username=admin&password=pass"))
    pn := fmt.Println
    body, _:= ioutil.ReadAll(resp.Body)
    pn(string(body))
}
