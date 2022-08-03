package main
import (
    "fmt"
)

/**
     s:="1234567890"
    fmt.Println(Slice(s, 0, 100))
    fmt.Println(Slice(s, 0, 10))
    fmt.Println(Slice(s, 0, 9))
    fmt.Println(Slice(s, 0, -1))
    fmt.Println(Slice(s, 0, 0))
    fmt.Println(Slice(s, -1, 0))
    fmt.Println("empty:",Slice(s, -1, -1))
    fmt.Println("empty:",Slice(s, -1, -2))
    fmt.Println(Slice(s, 2,3))
*/
func Slice(s string, start int, end int) string {
	l := len(s)
	// negative
	if start < 0 {
		start = l + start
	}
	if end <= 0 {
		end = l + end
	}

	// handle overflow
	if start < 0 {
		return ""
	}
	if end > l {
		end = l
	}
    // start<=end
    if start>end{
        return ""
    }
	return string([]rune(s)[start:end])
}

func main(){
    fmt.Println(1)
    s:="1234567890"
    fmt.Println(Slice(s, 0, 100))
    fmt.Println(Slice(s, 0, 10))
    fmt.Println(Slice(s, 0, 9))
    fmt.Println(Slice(s, 0, -1))
    fmt.Println(Slice(s, 0, 0))
    fmt.Println(Slice(s, -1, 0))
    fmt.Println("empty:",Slice(s, -1, -1))
    fmt.Println("empty:",Slice(s, -1, -2))
    fmt.Println(Slice(s, 2,3))
}
