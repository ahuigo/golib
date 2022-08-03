package eg

import (
	"fmt"
)

func Example_assert_output_ok() {
	fmt.Println("hello world")
	// Output:
	// hello world
}
func Example_assert_output_fail() {
	fmt.Println("hello world!")
	// Output:
	// hello world
}
