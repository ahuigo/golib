package main

/*
基于无缓冲的channel的main和 goroutine的同步
*/

import (
        "log"
        "time"
)

func main() {
        time.Sleep(0*time.Second)
        done := make(chan string) // 默认0，no　buffer

        go func() {
                time.Sleep(2*time.Second)
                done <- "I am done"
                time.Sleep(2*time.Second)
                <-done
                log.Println("slave: get ch2 done")
        }()
        log.Println("ch1 is blocked...2s")
        <-done
        log.Println("ch1: done!")

        log.Println("main: send ch2...2s")
        done<-"main: send ch2"
        log.Println("main: send ch2...done")

        time.Sleep(1*time.Second)
        log.Println("main: done")
}
