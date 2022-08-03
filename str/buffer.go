package main

import (
	"fmt"
//	"net/http"
//	"log"
//	"reflect"
    "os"
	"bytes"
)


func bufferNew(){
    buf := new(bytes.Buffer)
    buf.WriteString("new")
    println(buf.String())
}

func main() {
    bufferNew()

    //等价
   buf1 := bytes.NewBufferString("hello")
   buf2 := bytes.NewBuffer([]byte("hello"))
    var buf3 bytes.Buffer

   buf3.Write([]byte("hello ahui world"))
   buf3.WriteString(" world2")

   buf4 := bytes.NewBufferString("abc")
   buf5 := bytes.NewBuffer([]byte{})
   fmt.Println(buf1.String(),buf2.String(),buf3.String(),buf4,buf5,100)
   line,_ :=buf3.ReadString('w')
   fmt.Println(line)

   file, _ := os.Open("text.txt")
   buf := bytes.NewBufferString("bob ")
   buf.ReadFrom(file)
   fmt.Println(buf.String()) //bob hello world

   b := buf.Next(2)  //取前2个

    fmt.Printf("%#v\n", string(b))
//	buf := bytes.NewBuffer(make([]byte, 0, 512))

//	length, _ := buf.ReadFrom(resp.Body)
//
//	fmt.Println(len(buf.Bytes()))
//	fmt.Println(length)
//	fmt.Println(string(buf.Bytes()))
}

