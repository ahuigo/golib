package main

import (
	"encoding/json"
	"fmt"
)



type User struct{
    Name string `json:"name"`
    Age int  `json:"age"`
}

func vardump(name string, v interface{}){
    fmt.Printf("%s(%T):%#v\n\n",name, v, v)
}


func main() {
    // json string of user
    var u interface{}
    vardump("interface-u",u)
    u = User{Name:"hilo", Age:20}
    vardump("interface-u",u)
    userBytes, _ := json.Marshal(u)
    //[]byte(`{"name":"ahui","age":"20"}`

    // Umarshal interface
	var infObj interface{}
	fmt.Printf("example1:\n")
	json.Unmarshal(userBytes, &infObj)
	fmt.Printf("infObj:%+v\n", infObj)
	fmt.Printf("infObj[name]:%+v\n", infObj.(map[string]interface{})["name"])


    // Umarshal to interface
	fmt.Printf("\n2nd Umarshal: interface again\n")
    err:=json.Unmarshal([]byte(`{"name":"ahui2","age":"2"}`), &infObj)
	fmt.Printf("infObj:%+v, err=%v\n", infObj, err)


    // Umarshal to interface again
	fmt.Printf("\n3td Umarshal: interface again2\n")
    err=json.Unmarshal([]byte(`{"name":"ahui3"}`), &infObj)
	fmt.Printf("infObj:%+v, err=%v\n", infObj, err)

    // failed
    u2, ok:=infObj.(User)
    vardump("ok", ok)
    vardump("u2", u2)


    // Umarshal interface to bytes
	fmt.Printf("\n4td Umarshal to bytes: failed\n")
    b:=make([]byte, 10)
    err=json.Unmarshal([]byte(`{"name":"ahui3","age":3}`), &b)
	fmt.Printf("infObj:%+v, err=%v\n", string(b), err)


    // Umarshal interface to json.RawMessage
	fmt.Printf("\n4td Umarshal to json.RawMessage: succ\n")
    c:=make(json.RawMessage, 10)
    err=json.Unmarshal([]byte(`{"name":"ahui3","age":3}`), &c)
	fmt.Printf("infObj:%+v, err=%v\n", string(c), err)
}
