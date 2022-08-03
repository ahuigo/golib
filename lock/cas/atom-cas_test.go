package queue

import (
    "testing"
    "fmt"
    _ "time"
)

func BenchmarkPushTailCAS(b *testing.B) {
    fmt.Println("size0")
    b.StopTimer()
    q := New()
    b.StartTimer()
    m:=0
    for i := 0; i < b.N; i++ {
        m=i
        q.PushTailCAS(Message{id: i})
    }
    fmt.Println("size1:", q.Len(), m)
}

func BenchmarkPushTailMutex(b *testing.B) {
    b.StopTimer()
    q := New()
    b.StartTimer()
    m:=0
    for i := 0; i < b.N; i++ {
        m=i
        q.PushTailMutex(Message{id: i})
    }
    fmt.Println("size2:", q.Len(), m)
}

func BenchmarkPushTailCASFixed(b *testing.B) {
    b.StopTimer()
    q := New()
    b.StartTimer()
    m:=0
    for i := 0; i < b.N; i++ {
        m=i
        q.PushTailCASFixed(Message{id: i})
    }
    fmt.Println("size3:", q.Len(), m)
}

func BenchmarkPushTailMutexFixed(b *testing.B) {
    b.StopTimer()
    q := New()
    b.StartTimer()
    m:=0
    for i := 0; i < b.N; i++ {
        m=i
        q.PushTailMutexFixed(Message{id: i})
    }
    fmt.Println("size4:", q.Len(), m)
}
