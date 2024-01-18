package t

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestUmarshalPtr(t *testing.T) {
	testSlice()
	testPointer()
}

// testSlice
func testSlice() {
	fmt.Println("21. unmarshal slice failed:")
	var s []int
	json.Unmarshal([]byte(`[1,2,3]`), s)
	fmt.Printf("parse slice failed:%#v\n", s)
	json.Unmarshal([]byte(`[1,2,3]`), &s)
	fmt.Printf("parse &slice ok:%#v\n", s)
	fmt.Println("")
}

// testPointer
func testPointer() {
	type A struct {
		Name string `json:"name"`
		Age  string `json:"age"`
	}

	type Task struct {
		Name string `json:"name"`
		Age  string `json:"age"`
	}

	type Family struct {
		// work
		Friend2  string `json:"friend"`
		Other    string `json:other` //json:other 依然有效
		Son      string //对于大写成员，默认为`json:"son"`
		Daughter string `json:"DAU"` //json大小写不敏感
	}

	var stu *A
	// 1. 不能直接访问nil pointer： stu.Name = "ahui"
	// 2. &stu 是指向指针的指针
	fmt.Println("1.unmarshal could pass partial value:")
	err := json.Unmarshal([]byte(`{"name":"ahui","age":"20", "other":1}`), &stu)
	fmt.Println("err:", err, ",stu:", stu)
	err = json.Unmarshal([]byte(`{"age":"21"}`), &stu)
	fmt.Println("err:", err, ",stu:", stu)

	fmt.Println("\n2.if we pass value: obj/nil(no pointer &obj), we could not get value:")
	var obj *map[string]interface{}
	err = json.Unmarshal([]byte(`{"name":"ahui","age":"20"}`), obj)
	fmt.Println("err:", err, ",*obj->obj:", obj)

	fmt.Println("\n3.We should pass pointer: pointer &obj")
	var obj2 *map[string]interface{}
	err = json.Unmarshal([]byte(`{"name":"ahui","age":"20"}`), &obj2)
	fmt.Println("err:", err, ",*obj2->&obj:", obj2)

	var obj4 map[string]interface{}
	err = json.Unmarshal([]byte(`{"name":"ahui","age":"20"}`), &obj4)
	fmt.Println("err:", err, ",obj4->&obj:", obj4)

	fmt.Println("\n4.We should pass pointer: &interface{}")
	var infObj interface{}
	json.Unmarshal([]byte(`{"name":"ahui","age":20}`), &infObj)
	fmt.Printf("infObj:%+v\n", infObj)
	fmt.Printf("infObj[age]:%+v\n", infObj.(map[string]interface{})["age"])
	fmt.Printf("infObj[age]:%T\n", infObj.(map[string]interface{})["age"])

	fmt.Println("\n5.We should pass pointer: &([]*Task)")
	tasks := []*Task{}
	json.Unmarshal([]byte(`[{"name":"ahui","age":"20"}]`), &tasks)
	fmt.Printf("tasks:%+v\n", *(tasks[0]))
}
