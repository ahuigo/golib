package main

import (
	"fmt"
	"os"
	"reflect"
	"strings"
	"testing"
	"unsafe"
)

type Config struct {
	Name string `json:"server-name"` // CONFIG_SERVER_NAME
	ip   string
	age  int
}

/*
// https://geektutu.com/post/hpg-reflect.html
BenchmarkSet-8                          1000000000               0.302 ns/op
BenchmarkReflect_FieldSet-8             33913672                34.5 ns/op
BenchmarkReflect_FieldByNameSet-8        3775234               316 ns/op // 查找下标O(n) 10倍
// 普通的赋值操作，每次耗时约为 0.3 ns，reflect通过下标找到对应的字段再赋值，每次耗时约为 30 ns，通过名称找到对应字段再赋值，每次耗时约为 300 ns。
*/
func TestStructSet(t *testing.T) {
	os.Setenv("CONFIG_SERVER_NAME", "10.0.0.1")
	config := Config{}
	typ := reflect.TypeOf(config)
	// value := reflect.Indirect(reflect.ValueOf(&config))
	value := reflect.ValueOf(&config).Elem()
	for i := 0; i < typ.NumField(); i++ {
		f := typ.Field(i)
		vf1 := value.Field(i)
		if v, ok := f.Tag.Lookup("json"); ok {
			key := fmt.Sprintf("CONFIG_%s", strings.ReplaceAll(strings.ToUpper(v), "-", "_"))
			if env, exist := os.LookupEnv(key); exist {
				// 4种方式设置值
				vf2 := value.FieldByName(f.Name) //O((n)
				vf3 := value.Field(i)            // O(1) Refer: https://geektutu.com/post/hpg-reflect.html
				vf1.SetString(env + ":1")
				vf2.Set(reflect.ValueOf(env))
				vf2.SetString(env + ":2")
				vf3.SetString(env + ":3")
			}
		}
	}
	fmt.Printf("%+v\n", config)
}

func TestStructSetPrivate(t *testing.T) {
	config := Config{ip: "init ip"}
	_ = config.ip
	_ = config.age
	typ := reflect.TypeOf(config)
	// value := reflect.Indirect(reflect.ValueOf(&config))
	value := reflect.ValueOf(&config).Elem()
	for i := 0; i < typ.NumField(); i++ {
		vf := value.Field(i)
		if vf.CanAddr() {
			fieldPtr := unsafe.Pointer(vf.UnsafeAddr())
			if vf.Kind() == reflect.String {
				realPtrToPrivateField := (*string)(fieldPtr)
				*realPtrToPrivateField = "setip"
			}
			if vf.Kind() == reflect.Int {
				realPtrToPrivateField := (*int)(fieldPtr)
				*realPtrToPrivateField = 123
			}
		}
	}
	fmt.Printf("%+v\n", config)
}
