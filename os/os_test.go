package o

import (
	"fmt"
	"os"
    "testing"
)

func TestOS(t *testing.T) {
	name, err := os.Hostname()
	if err != nil {
		panic(err)
	}

	fmt.Println("hostname:", name)
}
