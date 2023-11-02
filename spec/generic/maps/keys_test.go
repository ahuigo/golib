// https://stackoverflow.com/questions/71905177/why-does-maps-keys-in-go-specify-the-map-type-as-m
package maps

import (
	"fmt"
	"testing"
)

func Keys[M ~map[K]V, K comparable, V any](m M) []K {
	r := make([]K, 0, len(m))
	for k := range m {
		r = append(r, k)
	}
	return r
}

/*
func Keys[K comparable, V any](m map[K]V) []K {
    r := make([]K, 0, len(m))
    for k := range m {
        r = append(r, k)
    }
    return r
}
*/

func TestMap(t *testing.T) {
	type Dictionary map[string]int
	m := Dictionary{"foo": 1, "bar": 2}
	k := Keys[Dictionary](m)
	fmt.Println(k) // it just works
	k2 := Keys(m)
	fmt.Println(k2) // it just works
}
