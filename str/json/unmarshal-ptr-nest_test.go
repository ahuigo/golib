package t

import (
	"encoding/json"
	"fmt"
	"testing"
)

type A interface{}

type Task struct {
	Name string
}
type Data struct {
	Tasks *[]Task
	Age   int
}

func TestUnmarshalNested(t *testing.T) {
	var tasks []Task
	data := Data{
		Tasks: &tasks,
	}
	rawbody := []byte(`{"tasks":[{"name":"hilojack"}], "Age":2}`)
	err := json.Unmarshal(rawbody, &data)
	fmt.Printf("data:%+v, err:%v\n", data, err)
	fmt.Printf("tasks:%+v\n", data.Tasks)
	fmt.Printf("tasks(pointer):%+v\n", tasks) // 有被赋值
}
