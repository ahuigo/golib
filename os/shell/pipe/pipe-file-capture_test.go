package demo

import (
	"os"
	"os/exec"
	"testing"
)

func TestPipeFileAsStdout(t *testing.T) {

    cmd := exec.Command("echo", "'WHAT THE HECK IS UP'")

    // open the out file for writing
    outfile, err := os.Create("./out.txt")
    if err != nil {
        panic(err)
    }
    defer outfile.Close()
    cmd.Stdout = outfile

    err = cmd.Start(); if err != nil {
        panic(err)
    }
    cmd.Wait()
}
