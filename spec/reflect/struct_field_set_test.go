package main

import (
	"fmt"
	"os"
	"reflect"
	"strings"
	"testing"
)

type Config struct {
	Name    string `json:"server-name"` // CONFIG_SERVER_NAME
	IP      string `json:"server-ip"`   // CONFIG_SERVER_IP
}

func readConfig() *Config {
	// read from xxx.json，省略
	config := Config{}
	typ := reflect.TypeOf(config)
	value := reflect.Indirect(reflect.ValueOf(&config))
	value = reflect.ValueOf(&config).Elem()
	for i := 0; i < typ.NumField(); i++ {
		f := typ.Field(i)
		if v, ok := f.Tag.Lookup("json"); ok {
			key := fmt.Sprintf("CONFIG_%s", strings.ReplaceAll(strings.ToUpper(v), "-", "_"))
			if env, exist := os.LookupEnv(key); exist {
				value.FieldByName(f.Name).Set(reflect.ValueOf(env))  //O((n)
				//value.FieldByName(f.Name).SetString(env)  
                value.Field(0).SetString(env)                      // O(1) Refer: https://geektutu.com/post/hpg-reflect.html
			}
		}
	}
	return &config
}

func TestStructSet(t *testing.T) {
	os.Setenv("CONFIG_SERVER_IP", "10.0.0.1")
	c := readConfig()
	fmt.Printf("%+v", c)
}
