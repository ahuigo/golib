package gotest

import (
	"testing"
)

/*
*
$ go test -v parallel_test.go
=== RUN   TestParallel
=== RUN   TestParallel/test_1
=== PAUSE TestParallel/test_1
=== RUN   TestParallel/test_2
=== PAUSE TestParallel/test_2
=== RUN   TestParallel/test_3
=== PAUSE TestParallel/test_3
=== RUN   TestParallel/test_4
=== PAUSE TestParallel/test_4
=== CONT  TestParallel/test_1

	parallel_test.go:23: 4

=== CONT  TestParallel/test_2
=== CONT  TestParallel/test_4
=== CONT  TestParallel/test_2

	parallel_test.go:23: 4

=== CONT  TestParallel/test_3

	parallel_test.go:23: 4

=== CONT  TestParallel/test_4

	parallel_test.go:23: 4

--- PASS: TestParallel (0.00s)

	--- PASS: TestParallel/test_1 (0.00s)
	--- PASS: TestParallel/test_2 (0.00s)
	--- PASS: TestParallel/test_3 (0.00s)
	--- PASS: TestParallel/test_4 (0.00s)
*/
func TestParallel(t *testing.T) {
	tests := []struct {
		name  string
		value int
	}{
		{name: "test 1", value: 1},
		{name: "test 2", value: 2},
		{name: "test 3", value: 3},
		{name: "test 4", value: 4},
	}
	for _, tc := range tests {
		v := tc.value
		// fix: tc:=tc
		// tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()    // pause and test other case(否则就是串行)
			t.Log(tc.value) // error: tc是闭包，永远为4
			if tc.value != v {
				t.Error("invalid v:", tc.value)
			}
		})
	}
}
