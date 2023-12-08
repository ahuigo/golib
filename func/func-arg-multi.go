package main
import "fmt"

type ClaimPayload interface{}
type UserType string
type UserType2 string

// golang/js的variadic arguments
func f(args ...interface{}) {
    claims := map[string]interface{}{}
    for _, arg := range args {
        switch arg := arg.(type) {
        case UserType:
            claims["type"] = arg
        case UserType2:
            claims["type2"] = arg
        case ClaimPayload:
            claims["payload"] = arg
        }
    }
    fmt.Println(claims)
}

func main(){
    var type1 UserType
    var type2 UserType2
    type2 = "t1"
    type1 = "t2"
    f(type1, type2)
    f(type2, type1)
}
