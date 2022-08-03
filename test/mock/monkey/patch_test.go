package main

import (
	"testing"
    "github.com/undefinedlabs/go-mpatch"
)


//go:noinline
func methodA() int { return 1 }

//go:noinline
func methodB() int { return 2 }

func TestPatcher(t *testing.T) {
	patch, err := mpatch.PatchMethod(methodA, methodB)
	if err != nil {
		t.Fatal(err)
	}
	if methodA() != 2 {
		t.Fatal("The patch did not work")
	}

	err = patch.Unpatch()
	if err != nil {
		t.Fatal(err)
	}
	if methodA() != 1 {
		t.Fatal("The unpatch did not work")
	}
}
