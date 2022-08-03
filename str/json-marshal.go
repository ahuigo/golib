package main

import (
	"fmt"
    "encoding/json"
)

type ErrorHttp struct {
	Method string `json:"method"`
	Url    string `json:"url"`
	Err    error  `json:"err"`
}


type Stu struct {
	Name string `json:"name"`
	Age  int
	height  int //json ignore
	ErrorType string
    Goods []int `json:"goods,omitempty"` //json ignore if len(Goods)=0
    Url    string `json:"url"`
    NilPtr *ErrorHttp `json:"nilPtr,omitempty"` // json ignore if omitempty
    ErrorHttp
}

type Dict map[string]interface{}

func main() {
     m,_:= json.Marshal(Dict{
        "a":1,
        "bytes":[]byte("中国"),
        "k2":"b",
        "k3":false,
        "time": Dict{
            "a" : 222222,
        },
    })
    fmt.Println("Marshal map with lowercase:",string(m))

    stu := Stu{Name:"ahui", Age:20, height:100, ErrorType:"error_type", Goods:[]int{}, Url:"sub_url", ErrorHttp:ErrorHttp{Url:"parent_url"}}
    m,_= json.Marshal(stu)
    fmt.Println("Warn: Marshal struct will ignore lowercase height and goods:", string(m))
    // json interface
    var stuI interface{}=stu
    m,_= json.Marshal(stuI)
    fmt.Println("json(interface):")
    fmt.Println(string(m))
}
