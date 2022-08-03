package main
import (
    "reflect"
    "fmt"
    "github.com/pkg/errors"
    "encoding/json"
)
func SetField(obj interface{}, name string, value interface{}) error {
    structValue := reflect.ValueOf(obj).Elem()
    structFieldValue := structValue.FieldByName(name)

    if !structFieldValue.IsValid() {
        //return fmt.Errorf("No such field: %s in obj", name)
        return nil
    }

    if !structFieldValue.CanSet() {
        return fmt.Errorf("Cannot set %s field value", name)
    }

    structFieldType := structFieldValue.Type()
    val := reflect.ValueOf(value)
    if structFieldType != val.Type() {
        return errors.Errorf("Provided value type(%v) didn't match obj field type(%s)", val.Type(), structFieldType)
    }

    structFieldValue.Set(val)
    return nil
}

type MyStruct struct {
    Name string
    Age  int64
    Extra  struct{
        Count int
    }
}

func (s *MyStruct) FillStruct(m map[string]interface{}) error {
    for k, v := range m {
        err := SetField(s, k, v)
        if err != nil {
            return err
        }
    }
    return nil
}

func case1() {
    myData := make(map[string]interface{})
    myData["Name"] = "Tony"
    myData["age"] = int64(23)
    /* not supported
    myData["Extra"] = map[string]int{
        "Count":11,
    }*/

    result := &MyStruct{}
    err := result.FillStruct(myData)
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(result)
}


// 第二种
func Interface2Struct(in interface{}, out interface{}) error {
	if bs, err := json.Marshal(in); err != nil {
		return err
	} else {
		err := json.Unmarshal(bs, out)
		return err
	}
}
func case2viajson(){
    in:=map[string]string{
        "name":"ahui",
    }
    out := struct{
        Name string
    }{}
    Interface2Struct(&in, &out)
    fmt.Println("Interface2Struct:out=", out)

}

func main() {
    case1()

    case2viajson()

    // case3
    // import "github.com/mitchellh/mapstructure"
}
