package print

import (
	"encoding/json"
	"fmt"
	"testing"
)

func OffTestPrintPointer(t *testing.T) {
	type User struct {
		Name string
		age  int
	}
	u := &User{Name: "Tom", age: 20}
	fmt.Printf("print:%#v\n", u)
	fmt.Printf("print:%v\n", u)
	bs, _ := json.Marshal(u)
	fmt.Printf("json:%s\n", string(bs)) // json不打印private
}


    type UserInfop struct {
		Age  int
		id int
	}
    func (ui *UserInfop) String() string {
        return fmt.Sprintf("%#v", ui)
    }

func OffTestPrintInnerPointer(t *testing.T) {
    type UserInfo struct {
		Age  int
		id int
	}

	type ExtUserInfo struct {
		u *UserInfo
        U *UserInfop

	}
    v := ExtUserInfo{u:&UserInfo{id: 2, Age:3}, U:&UserInfop{id: 3, Age:4}}
    a := fmt.Sprintf("print:%#v\n", v)        // print %#v 能打印private, 不能输出inner pointer原始值
    b := fmt.Sprintf("print:%v\n", v)        // print %v   能打印private, %v输出inner pointer原始值
	fmt.Println("print a:", a)       
	fmt.Println("print b:", b)        

    bs, _ := json.Marshal(v)
	fmt.Printf("json:%#v\n", string(bs))// json 不能打印private, 但能输出inner pointer原始值
}

// 似乎是有序的
func TestPrintMap(t *testing.T) {
    type Usermap = map[string]int
    v := Usermap{"b":2, "c":3,}
    v["a"] = 100
    a := fmt.Sprintf("print:%#v\n", v)     
    b := fmt.Sprintf("print:%v\n", v)     
	fmt.Println("print a:", a)       
	fmt.Println("print b:", b)        

    bs, _ := json.Marshal(v)
	fmt.Printf("json:%#v\n", string(bs))
}
