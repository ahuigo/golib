package main

import (
	ahui2 "ahui1"
	"ahui1/hello"

	anyname "github.com/ahuigo/go-hello"
)

func main() {

	println("-------------start anyname.Test----------")
	anyname.Test()
	println("-------------start anyname.test2----------")
	anyname.Test2()
	println("-------------start ahui.Test3----------")
	ahui2.Test3()
	// ahui2 is package name
	println(ahui2.Add(1, 5))
	hello.Hello()
}
