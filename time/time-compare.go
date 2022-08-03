package main

import "time"

func main() {
	t1, _ := time.Parse(time.RFC3339, "2012-11-01T22:08:41+00:00")
	t2, _ := time.Parse(time.RFC3339, "2012-11-01T22:08:41+00:00")
	println(t1 == t2)     //false
	println(t1.Equal(t2)) //true
}
