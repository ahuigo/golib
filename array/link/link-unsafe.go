    package main
    import (
        "container/list"
        "fmt"
    )
    
    func main() {
        // 生成队列
        l := list.New()
        // 入队, 压栈
        l.PushBack(1)
        l.PushBack(2)
        // 出队
        v1 := l.Front()
        l.Remove(v1)
        fmt.Printf("%d\n", v1.Value)
        // 出栈
        a1 := l.Back()
        l.Remove(a1)
        fmt.Printf("%d\n", a1.Value)
    }
