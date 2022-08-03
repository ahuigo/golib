package main

import (
	"fmt"
    "strconv"
//	"net/http"
//	"log"
//	"reflect"
    //"os"
	//"bytes"
)


func main() {
    msg := `I have ` + strconv.Itoa(10) + `个error\n`
    msg += fmt.Sprintf("%05d:", 5)
    fmt.Println(msg)
    i, err := strconv.Atoi("077")
    fmt.Println("atoi:", i, err)
}

