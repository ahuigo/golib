package main

import (
	"encoding/json"
	"fmt"
)

type A interface{}


type Task struct{
    name string
}
type Data struct{
    Tasks *[]Task
    Age int
}


func main() {
    var tasks []Task
    data := Data{
        Tasks:&tasks,
    }
    rawbody := []byte(`{"tasks":[{"name":"hilojack"}], "Age":2}`)
    err:=json.Unmarshal(rawbody, &data)
    fmt.Printf("data:%+v, err:%v\n", data, err)
	fmt.Printf("tasks:%+v\n", data.Tasks)
	fmt.Printf("tasks(pointer):%+v\n", tasks)
}
