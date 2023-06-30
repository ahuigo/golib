package main

import (
	"bytes"
	"fmt"
	"io"
	"reflect"
	"strings"
	"testing"
	"unicode"
	"unicode/utf8"

	"gopkg.in/yaml.v3"
)

func firstToLower(s string) string {
	r, size := utf8.DecodeRuneInString(s)
	if r == utf8.RuneError && size <= 1 {
		return s
	}
	lc := unicode.ToLower(r)
	if r == lc {
		return s
	}
	return string(lc) + s[size:]
}

func TestYamlMultiple(test *testing.T) {
	var data2 = `
a: Easy!
b:
  c: 2
  d: [3, 4]
`
	t2 := struct {
		A string
		B struct {
			RenamedC int   `yaml:"c"`
			D        []int `yaml:",flow"`
		}
	}{}
	err := UnmarshalYaml([]byte(data2), &t2)
	fmt.Printf("case1:%v, err:%v\n", t2, err)

	var data = `
a: Easy!
b:
  c: 2
  d: [3, 4]
---
e: Easy2!
f:
  g: 2
  h: [3, 4]
`
	var t interface{}
	err = UnmarshalYaml([]byte(data), &t)
	fmt.Printf("case2: %v, err:%v\n", t, err)
}
func convertValue(src, target reflect.Value) {
	switch src.Kind() {
	case reflect.Map:
		if target.Kind() == reflect.Struct {
			srcMap := src.Interface().(map[string]interface{})
			for i := 0; i < target.NumField(); i++ {
				field := target.Type().Field(i)
				// fieldName := field.Name
				fieldNameKey := firstToLower(field.Name)
				yamlTag := field.Tag.Get("yaml")
				if len(yamlTag) > 0 && yamlTag != "-" {
					tagParts := strings.SplitN(yamlTag, ",", 2)
					if tagParts[0] != "" {
						fieldNameKey = tagParts[0]
					}
				}
				if value, ok := srcMap[fieldNameKey]; ok {
					convertValue(reflect.ValueOf(value), target.Field(i))
				}
			}
		} else if target.Kind() == reflect.Map {
			target.Set(src)
		}
	case reflect.Slice:
		targetType := target.Type()
		targetElemType := target.Type().Elem()
		targetSlice := reflect.MakeSlice(targetType, src.Len(), src.Len())
		for i := 0; i < src.Len(); i++ {
			srcElem := src.Index(i)
			if srcElem.Type().Kind() == reflect.Interface {
				srcElem = srcElem.Elem()
			}
			if !srcElem.Type().ConvertibleTo(targetElemType) {
				panic("invalid type conversion")
			}
			targetSlice.Index(i).Set(srcElem.Convert(targetElemType))
		}
		// 将转换后的切片设置为目标值
		target.Set(targetSlice)

	default:
		target.Set(src)
	}
}

func UnmarshalYaml(data []byte, result interface{}) error {
	resultValue := reflect.ValueOf(result)
	if resultValue.Kind() != reflect.Ptr {
		return fmt.Errorf("result must be a pointer to the target value")
	}

	decoder := yaml.NewDecoder(bytes.NewReader(data))

	var documents []interface{}
	for {
		var t interface{}
		err := decoder.Decode(&t)
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}
		documents = append(documents, t)
	}

	targetValue := reflect.ValueOf(result).Elem()
	if len(documents) == 1 {
		convertValue(reflect.ValueOf(documents[0]), targetValue)
	} else {
		targetValue.Set(reflect.ValueOf(documents))
	}

	return nil
}

func isValidYaml(b []byte) bool {
	var t interface{}
	err := UnmarshalYaml([]byte(b), &t)
	fmt.Printf("%#v, err:%v\n", t, err)
	return err == nil
}

func TestIsValidYaml(t *testing.T) {
	var data = `
a: Easy!
b:
  c: 2
  d: [3, 4]
---
e: Easy2!
f:
  g: 2
  h: [3, 4]
`
	println(isValidYaml([]byte(data)))
	println(isValidYaml([]byte(`xxx: "sss`)))
}
