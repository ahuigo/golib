package server

import (
	"errors"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"testing"
)

type Args struct {
	A, B int
}

type Quotient struct {
	Quo, Rem int
}

type Arith struct{}

func (t *Arith) Multiply(args *Args, reply *int) error {
	*reply = args.A * args.B
	return nil
}

func (t *Arith) Divide(args *Args, quo *Quotient) error {
	if args.B == 0 {
		return errors.New("divide by zero")
	}
	quo.Quo = args.A / args.B
	quo.Rem = args.A % args.B
	return nil
}

func TestServer(t *testing.T) {
    // 0. Register rpc 
    arith := new(Arith)
    if err := rpc.Register(arith);  err != nil {
        panic(err)
    }
    println("tcp", ":1234")
    l, e := net.Listen("tcp", ":1234")
    if e != nil {
        log.Fatal("listen error:", e)
    }
    c :=1
    if c==0{
        /**************************
        // client, err := rpc.DialHTTP("tcp", "localhost:1234")
        ***************/
        // 1. register http handler
        rpc.HandleHTTP()

        // 2. http listen server
        http.Serve(l, nil)

    }else{
        /**************************
        // client, err := rpc.Dial("tcp", "localhost:1234")
        ***************/
        // tcp server
        for {
            rpc.Accept(l)
        }
    }


}
