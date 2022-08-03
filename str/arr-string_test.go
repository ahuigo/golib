package main
import(
    "strings"
    "fmt"
    "testing"
)

func Test_Array2str(t *testing.T){
    A := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
    delim:=","
    a:=fmt.Sprint(A)
    println(a)
    a=strings.Trim(strings.Join(strings.Fields(fmt.Sprint(A)), delim), "[]")
    println(a)
    a=strings.Trim(strings.Join(strings.Split(fmt.Sprint(A), " "), delim), "[]")
    println(a)
    a=strings.Trim(strings.Replace(fmt.Sprint(A), " ", delim, -1), "[]")
    println(a)
}
