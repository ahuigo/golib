package main

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
)

// 按照tag打印结构体
func TestStructTag(t *testing.T) {
	type Person struct {
		Name        string `label:"Person Name: " uppercase:"true"`
		Age         int    `label:"Age is: "`
		Sex         string `label:"Sex is: "`
		Description string
	}

	person := Person{
		Name:        "Tom",
		Age:         29,
		Sex:         "Male",
		Description: "Cool",
	}
	var ptr any = &person

	// 获取入参的类型
	rv := reflect.TypeOf(ptr)

	// 入参类型校验
	if rv.Kind() != reflect.Ptr || rv.Elem().Kind() != reflect.Struct {
		panic("参数应该为结构体指针")
	}

	// 取指针指向的结构体变量
	v := reflect.ValueOf(ptr).Elem()

	// 解析字段
	for i := 0; i < v.NumField(); i++ {

		// 取tag
		fieldInfo := v.Type().Field(i)
		tag := fieldInfo.Tag

		// 解析label tag
		label := tag.Get("label")
		if label == "" {
			label = fieldInfo.Name + ": "
		}

		// 解析uppercase tag
		value := fmt.Sprintf("%v", v.Field(i))
		if fieldInfo.Type.Kind() == reflect.String {
			uppercase := tag.Get("uppercase")
			if uppercase == "true" {
				value = strings.ToUpper(value)
			} else {
				value = strings.ToLower(value)
			}
		}

		fmt.Println(label + value)
	}

}
