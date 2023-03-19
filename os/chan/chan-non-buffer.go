package main

/*
基于无缓冲的channel的main和 goroutine的同步
*/

import (
        "io"
        "log"
        "net"
        "time"
        "os"
)

func main() {
        time.Sleep(0*time.Second)
        // nc -l 9125
        conn, err := net.Dial("tcp", "127.0.0.1:9125")
        if err != nil {
                log.Fatal(err)
        }

        done := make(chan string) // 默认0，no　buffer

        go func() {
                io.Copy(os.Stdout, conn)
                log.Println("wait groutine..")
                done <- "I am done"
                time.Sleep(1*time.Second)
                log.Println("groutine done") // 不会打印, main 提前结束
        }()

        //从客户端输入,将客户端标输入的数据发给客户端套接字
        log.Println("stdin(ctrl+d to close stdin):")
        io.Copy(conn, os.Stdin)
        log.Println("stdin closed...")

        conn.Close() //此时main要主动关闭conn, 否则goroutine里面的io.Copy()会一直阻塞等待conn

        log.Println("main sleep...")
        time.Sleep(3*time.Second)
        <-done
        log.Println("main: done!")

        //这样我们就保证了 "main::done!"打印之前 一定先打印"wait groutine...!"
}
