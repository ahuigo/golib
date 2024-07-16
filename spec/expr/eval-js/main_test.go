package m

import (
	"fmt"
	"testing"

	"github.com/dop251/goja"
)

func TestEvalExprs(t *testing.T) {
	vm := goja.New()
	vm.Set("headers", map[string]string{"token": ""})
	vm.Set("body", map[string]any{"version": 18})
	v, err := vm.RunString(` 
		const token = headers["token"];
		function matchResp(body){
			if (!token) {
				return false;
			}
			if(body.version == 18){
				return true;
			}else{
				return false;
			}
		}
		matchResp(body)
	`)
	if err != nil {
		panic(err)
	}
	if isMatched, ok := v.Export().(bool); !ok {
		t.Fatal(err)
	} else if isMatched {
		t.Log("Response is matched")
	} else {
		t.Log("Response is not matched")
	}
}
func TestEvalExpr(t *testing.T) {
	vm := goja.New()
	vm.Set("a", 2)
	vm.Set("b", 2)
	v, err := vm.RunString(` a + b `)
	if err != nil {
		panic(err)
	}
	if num := v.Export().(int64); num != 4 {
		panic(num)
	}
}

func TestEvalFunc(t *testing.T) {
	const SCRIPT = `
	function sum(a, b) {
		return +a + b;
	}
	`

	vm := goja.New()
	_, err := vm.RunString(SCRIPT)
	if err != nil {
		panic(err)
	}

	var sum func(int, int) int
	err = vm.ExportTo(vm.Get("sum"), &sum)
	if err != nil {
		panic(err)
	}

	fmt.Println(sum(40, 2)) // note, _this_ value in the function will be undefined.
}

func TestEvalSetStruct(t *testing.T) {
	vm := goja.New()
	vm.SetFieldNameMapper(goja.TagFieldNameMapper("json", true))
	type S struct {
		Field int `json:"field"`
	}
	vm.Set("s", S{Field: 42})
	res, _ := vm.RunString(`s.field`) // without the mapper it would have been s.Field
	fmt.Println(res.Export())
	// Output: 42
}
