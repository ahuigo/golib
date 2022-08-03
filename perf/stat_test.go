package demo

import (
	"fmt"
	"os"
	"testing"

	"github.com/mackerelio/go-osstat/memory"
)

func TestStatMemory(t *testing.T) {
	memory, err := memory.Get()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		return
	}
	fmt.Printf("memory total: %d bytes\n", memory.Total)
	fmt.Printf("memory used: %d bytes\n", memory.Used)
	fmt.Printf("memory cached: %d bytes\n", memory.Cached)
	fmt.Printf("memory free: %d bytes\n", memory.Free)
}
