package main

/*

单方向channel
*/

import (
        "fmt"
        "time"
)

//Counter ---naturals---> Squarer ---squares---> Printer

func Counter(out chan<- int) {
        for x := 0; x < 5; x++ {
                out <- x
                time.Sleep(time.Second)
        }
        //只能关闭输出channel,语法限定
        close(out)
}

func Squarer(out chan<- int, in <-chan int) {
        for x := range in {
                out <- x * x
        }
        //只能关闭输出类型channel,语法限定,
        //因为只有调用输出channel的groutine才能决定数据是否传送完毕
        close(out)
}

func Printer(in <-chan int) {

        for value := range in {
                fmt.Println(value)
        }
}

func main() {
        naturals := make(chan int)
        squares := make(chan int)

        //Counter
        go Counter(naturals)

        //Squarer
        go Squarer(squares, naturals)

        //Printer 执行printer不能用go, 因为那样main groutine会直接退出,导致程序退出
        Printer(squares)

        fmt.Println("done!")
}
