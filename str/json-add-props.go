package main

import (
	"encoding/json"
	"fmt"
    "errors"
)



type User struct{
    Name string `json:"name"`
    Age int  `json:"age"`
}

func vardump(name string, v interface{}){
    fmt.Printf("%s(%T):%#v\n\n",name, v, v)
}


func addProp2Json(buf []byte, key, value string) ([]byte, error){
	var obji interface{}
	json.Unmarshal(buf, &obji)
    if objm, ok := obji.(map[string]interface{}); !ok{
        return nil, errors.New("could not parse input")
    }else{
        objm[key]=value
        buf, _ := json.Marshal(objm)
        return buf,nil
    }
    
}

func main() {
    var u interface{}
    u = User{Name:"hilo", Age:20}
    vardump("interface-u",u)
    buf, _ := json.Marshal(u)

    buf, err:= addProp2Json(buf, "key1", "key2")
    if err!=nil{
        println(err.Error())
    }else{
        vardump("output:",string(buf))
    }
}
