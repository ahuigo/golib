// You can edit this code!
// Click here and start typing.
package demo

import (
	"fmt"
	"testing"

	"github.com/samber/lo"
)

func TestMapMerge(t *testing.T) {
	result := lo.Assign(
		map[string]int{"a": 1, "b": 2},
		map[string]int{"b": 3, "c": 4},
	)

	fmt.Printf("%v", result)
}
