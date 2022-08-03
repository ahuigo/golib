package main
import "time"

func main(){
    stopCh :=make(chan struct{})
    go func(){
        select {
		case <-stopCh:
            println("stop")
			return
		}
    }()
    println("sleep 5s")
    time.Sleep(5 * time.Second)
    close(stopCh)
    //time.Sleep(100000000000 * time.Millisecond)
}
