package main
import (
    "time"
    "fmt"
    "sync"
)

// 交替打印字符
// go-lib/goroutine/string-atom-race.go
const (
  FIRST  = "WHAT THE"
  SECOND = "F*CK"
)

var mutex = &sync.Mutex{}

func main() {
  var s string
  go func() {
    i := 1
    for {
      i = 1 - i
       mutex.Lock()
      if i == 0 {
        s = FIRST
      } else {
        s = SECOND
      }
      mutex.Unlock()
      time.Sleep(10)
    }
  }()

  for {
       mutex.Lock()
    fmt.Println(s)
      mutex.Unlock()
    time.Sleep(10)
  }
}
