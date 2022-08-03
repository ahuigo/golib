package main

import (

    "fmt"

    errors081 "github.com/pkg/errors/081"

    errors091 "github.com/pkg/errors/091"

)




func main() {

    err := errors081.New("New error for v0.8.1")

    fmt.Println(err)

    err = errors091.New("New error for v0.9.1")

    fmt.Println(err)

}
1
2
3
4
5
6
7
8
9
10
11
12
13
14
15
16
17
18
19
20
21
22
23
24
25
26
package main
 
import (
 
    "fmt"
 
    errors081 "github.com/pkg/errors/081"
 
    errors091 "github.com/pkg/errors/091"
 
)
 
 
 
 
func main() {
 
    err := errors081.New("New error for v0.8.1")
 
    fmt.Println(err)
 
    err = errors091.New("New error for v0.9.1")
 
    fmt.Println(err)
 
