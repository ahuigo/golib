package main

import (
    "container/list"
    "fmt"
    "sync"
    "time"
)

type Queue struct {
    data *list.List
}

func NewQueue() *Queue {
    q := new(Queue)
    q.data = list.New()
    return q
}

func (q *Queue) push(v interface{}) {
    defer lock.Unlock()
    lock.Lock()
    q.data.PushFront(v)
}

func (q *Queue) pop() interface{} {
    defer lock.Unlock()
    lock.Lock()
    iter := q.data.Back()
    v := iter.Value
    q.data.Remove(iter)
    return v
}

func (q *Queue) dump() {
    for iter := q.data.Back(); iter != nil; iter = iter.Prev() {
        fmt.Println("item:", iter.Value)
    }
}

var lock sync.Mutex

func main() {
    q := NewQueue()
    go func() {
        q.push("one")
    }()
    go func() {
        q.push("four")
    }()
    q.push("two")
    q.push("three")
    v := q.pop()
    fmt.Println("pop v:", v)
    fmt.Println("......")
    time.Sleep(1 * time.Second)
    q.dump()
}
