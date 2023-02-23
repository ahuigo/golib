package main

import (
	"fmt"
	"path/filepath"
)
func main() {
    m,err:=filepath.Match("/home/catch/**", "/home/catch/foo/bar") ///不支持** // false
	fmt.Println("dir match", m, err) 
    m, err = filepath.Match("[^0-9]ata[0-9]*", "data123.csv") // true
	fmt.Println("dir match", m, err) 
    m, err = filepath.Match("*-h.ahui.cn", "h.ahui.cn") // false
	fmt.Println("dir match", m, err) 
    m, err = filepath.Match("ahui.cn", "h.ahui.cn") // false
	fmt.Println("dir match", m, err) 
    m, err = filepath.Match("production-management:**","production-management:/api/v1/worker/register") // false
	fmt.Println("dir match", m, err) 
}
