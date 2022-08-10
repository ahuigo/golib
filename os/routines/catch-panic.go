package main

import (
	"fmt"
	"time"
    "github.com/pkg/errors"
    "math/rand"
)

// 相当于getFlowChart3
func task(t int, seq int) (int, error ){
    time.Sleep(time.Duration(randIntn(100))* time.Millisecond)
    if t<=0{
        return 0,nil
    } else if t==1{
        return 1, errors.Errorf("some error(task:%d)", seq)
    } else{
        s:=fmt.Sprintf("(task:%d)", seq)
        panic("panic fatal error"+s)
    }
}

type Resp struct{
    n int
    err error
}

func catchPanic(seq int){
    fmt.Printf("catch panic:task%d:%v\n",seq, recover())
}

// Server Request
func handler() (n int, err error){
    ch1 := make(chan Resp)
    ch2 := make(chan Resp)

    // start task
    t := randIntn(3) 
    t=1
    fmt.Println("randt:", t)
    go func() {
       // defer catchPanic(1)
        fmt.Println("task1 start...")
        n,err:=task(t,1)
        if err==nil{
            ch1<-Resp{n,err}
        }
        fmt.Println("task1 end")
    }()
    go func() {
        //defer catchPanic(2)
        fmt.Println("task2 start...")
        n,err:=task(t,2)
        if err==nil{
            ch2<-Resp{n,err}
        }
        fmt.Println("task2 end")
    }()

    // get response
    var resp Resp
    select{
    case resp=<-ch1:
        fmt.Println("ch1:", resp)
    case resp=<-ch2:
        fmt.Println("ch2:", resp)
    }
    fmt.Println("return:",resp,"\n")
    return resp.n,resp.err
}


func main(){
    handler()
    time.Sleep(time.Hour)
}


func randIntn(n int) int {
    s := rand.NewSource(time.Now().UnixNano())
    return rand.New(s).Intn(n)
}
