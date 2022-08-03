package main

import (
	"encoding/json"
	"fmt"
    "reflect"
)

func main() {
    testMemberCase()
}

// testMemberCase
func testMemberCase(){
	output := struct{
		Result interface{} `json:"result"`
	    Has *bool `json:"has"`
	}{}
    json.Unmarshal([]byte(`{"has":true}`), &output)
	fmt.Printf("output:%#v\n", output)
	fmt.Printf("output:%+v\n", output)
	fmt.Printf("outputType:%#v\n", isNil(output.Result))
}


func isNil(c interface{}) bool{
    //At least it detects (*T)(nil) cases.
    return c == nil || (reflect.ValueOf(c).Kind() == reflect.Ptr && reflect.ValueOf(c).IsNil())
}
