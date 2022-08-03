//https://medium.com/@deckarep/the-new-kid-in-town-gos-sync-map-de24a6bf7c2c
package main
import (
    "sync"
    "fmt"
)


type User struct{
    age int
}
type Key struct{
    name string
}

func test_update(){
   var sm sync.Map
   key := &Key{}
   m := map[string]int{
       "age":1,
   }
   mt:=m
   mt["age"] = 2
   fmt.Printf("ori m: %#v\n", m)
    sm.Store(key, m)

    // add 1 by reference
    v,_:= sm.Load(key)
    m2 := v.(map[string]int)
    m2["age"]+=100

    // read m
    v,_= sm.Load(key)
    m3 := v.(map[string]int)
    fmt.Printf("%#v\n", m3)
}

func test_range(){
   var sm sync.Map
    sm.Store("key1", 1)
    sm.Store("key2", "v2")
    sm.Store("key3", "v3")
    sm.Range(func(k, v interface{}) bool {
        fmt.Println("k:", k.(string))
        if k.(string)=="key2"{
            return false
        }
        return true
    })
}
func main(){
    /*
    println("test_update--------")
    test_update()
    */
    println("test_range--------")
    test_range()

}
