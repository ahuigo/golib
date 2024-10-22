package gotest

import (
	"log"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	log.Println("Do tearUp BEFORE the tests!")
	exitVal := m.Run()
	log.Println("Do tearDown AFTER the tests!")

	os.Exit(exitVal)
}

func TestA(t *testing.T) {
	log.Println("TestA running")
}

func TestB(t *testing.T) {
	log.Println("TestB running")
}
