package main
import "fmt"

func f(ids *[]byte){
    *ids = (*ids)[1:]
}

func main(){
    ids := []byte{1,2,3,4,5,6,7,8,9,0}
    fmt.Printf("ids:%v, len(ids)=%d\n",ids, len(ids))
    f(&ids)
    fmt.Printf("ids:%v, len(ids)=%d\n",ids, len(ids))

}
