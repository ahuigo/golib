package main
import "time"

func main(){
    stopCh :=make(chan struct{})
    go func(){
        select {
		case <-stopCh:
            time.Sleep(1 * time.Second)
            println("stop")
			return
		}
    }()
    close(stopCh)
    println("quit break goroutinue")
    //time.Sleep(2 * time.Second)
}
