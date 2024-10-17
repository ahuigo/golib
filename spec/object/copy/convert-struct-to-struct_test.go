package demo

import (
	"fmt"
	"testing"

	"github.com/jinzhu/copier"
)

func TestCopyStructToStruct(t *testing.T) {
	type User struct {
		Name string
		Role string
		Age  int32
        flag int // ignore when copy
	}

	type Employee struct {
		Name      *string
		Age       int32
		DoubleAge int32
		SuperRole string
        flag int
	}
    user := User{Name: "Jinzhu", Age: 18, Role: "Admin", flag:1}
	employee := Employee{}

	copier.Copy(&employee, &user)
	fmt.Printf("%#v\n", employee)
	if employee.Name != nil {
		fmt.Printf("name:%s\n", *employee.Name)
	}
	// Output: Employee{Name:"Jinzhu", Age:18, DoubleAge:36, SuperRole:"Super Admin"}
}
