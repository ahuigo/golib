// You can edit this code!
// Click here and start typing.
package demo

import (
	"fmt"
	"testing"

	"github.com/samber/lo"
	"golang.org/x/exp/maps"
)

func TestMapMerge(t *testing.T) {
	result := lo.Assign(
		map[string]int{"a": 1, "b": 2},
		map[string]int{"b": 3, "c": 4},
	)

	fmt.Printf("%v", result)
}

func TestMapMerge2(t *testing.T) {
	src := map[string]int{
		"one": 1,
		"two": 2,
	}
	dst := map[string]int{
		"two":   2,
		"three": 3,
	}
	maps.Copy(dst, src)
	fmt.Println("src:", src)
	fmt.Println("dst:", dst)
}
