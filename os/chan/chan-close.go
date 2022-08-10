package main

import (
    "sync"
    "time"
)

/**
eventually read the "true" value from it (v <- c)
read the "true" value and 'not closed' indicator (v, ok <- c)
read a zero value and the 'closed' indicator (v, ok <- c)
will block in the channel read forever (v <- c)
**/
func closeCh() {
    var wg sync.WaitGroup
    wg.Add(1)

    ch := make(chan int, 0) //默认是0, <-ch 第一个就block
    go func() {
        println("start routine")
        for {
            foo, ok := <- ch
            if !ok {
                println("done", foo)
                wg.Done()
                return
            }
            println(foo)
            time.Sleep(100*time.Millisecond)
        }
    }()
    ch <- 1
    ch <- 2
    ch <- 3
    close(ch)

    wg.Wait()
}



func main(){
    closeCh()
}
