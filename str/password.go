package main

import "fmt"
import "time"
import "math/rand"
func GeneratePassword()(string){
    rand.Seed(time.Now().UnixNano())
    digits := "0123456789"
    specials := "~=+%^*/()[]{}/!@#$?|"
    all := "ABCDEFGHIJKLMNOPQRSTUVWXYZ" +
        "abcdefghijklmnopqrstuvwxyz" +
        digits + specials
    length := 8
    buf := make([]byte, length)
    buf[0] = digits[rand.Intn(len(digits))]
    buf[1] = specials[rand.Intn(len(specials))]
    for i := 2; i < length; i++ {
        buf[i] = all[rand.Intn(len(all))]
    }
    rand.Shuffle(len(buf), func(i, j int) {
        buf[i], buf[j] = buf[j], buf[i]
    })
    return string(buf) // E.g. "3i[g0|)z"
}
func main(){
    s:=GeneratePassword()
    fmt.Println(s)
}
