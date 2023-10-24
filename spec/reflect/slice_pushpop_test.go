package main

// refer: https://coolshell.cn/articles/21179.html
import (
	"fmt"
	"reflect"
)

/**
可以通过reflect.Slice 实现
	.Slice(0, len-1) //pop last element
	.Slice(1, len-1) //shift first element
	.Slice(1, len-1) //prepend element
	reflect.Append(c.s, reflect.ValueOf(val))
*/
type Container struct {
	s reflect.Value
}

func NewContainer(t reflect.Type, size int) *Container {
	if size <= 0 {
		size = 64
	}
	return &Container{
		s: reflect.MakeSlice(reflect.SliceOf(t), 0, size),
	}
}
func (c *Container) Append(val interface{}) error {
	if reflect.ValueOf(val).Type() != c.s.Type().Elem() {
		return fmt.Errorf(`Put: cannot put a %T into a slice of %s`, val, c.s.Type().Elem())
	}
	c.s = reflect.Append(c.s, reflect.ValueOf(val))
	return nil
}

func (c *Container) Shift(refval interface{}) error {
	if reflect.ValueOf(refval).Kind() != reflect.Ptr || reflect.ValueOf(refval).Elem().Type() != c.s.Type().Elem() {
		return fmt.Errorf("Get: needs *%s but got %T", c.s.Type().Elem(), refval)
	}
	reflect.ValueOf(refval).Elem().Set(c.s.Index(0))
	c.s = c.s.Slice(1, c.s.Len())
	return nil
}

// 实现类似javascript中array的unshift功能
func Unshift(slice, v interface{}) interface{} {
	var typ = reflect.TypeOf(slice)
	if typ.Kind() == reflect.Slice {
		var vv = reflect.ValueOf(slice)
		var tmp = reflect.MakeSlice(typ, vv.Len()+1, vv.Cap()+1)
		tmp.Index(0).Set(reflect.ValueOf(v))
		var dst = tmp.Slice(1, tmp.Len())
		reflect.Copy(dst, vv)
		return tmp.Interface()
	}
	panic("not a slice")
}
