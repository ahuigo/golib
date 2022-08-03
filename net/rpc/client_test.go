package server

import (
	"fmt"
	"log"
	"net/rpc"
	"testing"
)

func TestClient(t *testing.T) {
	type Args struct {
		A, B int
	}

	type Quotient struct {
		Quo, Rem int
	}

	//client, err := rpc.DialHTTP("tcp", "localhost:1234")
    client, err := rpc.Dial("tcp", "localhost:1234")

	if err != nil {
		log.Fatal("dialing:", err)
	}

	// Synchronous call
	args := &Args{11, 8}
	var reply int
	err = client.Call("Arith.Multiply", args, &reply)
	if err != nil {
		log.Fatal("arith error:", err)
	}
	fmt.Printf("Arith: %d*%d=%d\n", args.A, args.B, reply)

	// Asynchronous call
	quotient := new(Quotient)
	divCall := client.Go("Arith.Divide", args, quotient, nil)
	replyCall := <-divCall.Done // will be equal to divCall
	fmt.Printf("replyCall:%#v\n", replyCall)
	fmt.Printf("replyCall.Reply:%#v\n", replyCall.Reply)
	fmt.Printf("quotient:%#v\n", quotient)
	// check errors, print, etc.
}
