package main
import (
    "github.com/go-resty/resty/v2"
    "fmt"
)

func main(){
    fmt.Println(1)
    client := resty.New()
    var res interface{}
    resp, err := client.R().
		SetResult(&res).
		SetQueryParam("version", "2").
		SetQueryParam("domain", "d1").
        Get("http://m:7900/get1?");
    if err!=nil{
        panic(err)
    }
    fmt.Printf("res:%+v\n", res)
    fmt.Printf("hasError:%+v\n", resp.IsError())
    fmt.Printf("body:%+v\n", resp.String()[:3])
    fmt.Printf("body:%+v\n", "中国人名"[:3])
    fmt.Printf("body:%+v\n", len("中国人名"))
    fmt.Printf("body:%+v\n", string([]byte{'1'}))
}
