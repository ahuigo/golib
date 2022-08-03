package main

import (
    "bytes"
    "fmt"
    "os/exec"
    "errors"
)

func runCommand(name string, args ...string) (out string, err error) {
	var stderr bytes.Buffer
	var stdout bytes.Buffer
	cmd := exec.Command(name, args...)
	cmd.Stderr = &stderr
	cmd.Stdout = &stdout
	if err := cmd.Run(); err != nil {
		out = stdout.String()
		if stderr.Len() > 0 {
			return out, errors.New(stderr.String())
		}
		return out, err
	}
	return stdout.String(), nil
}

func main() {
    out, err := runCommand("ls", "-l")
    fmt.Println("--- out ---\n", out)
    fmt.Println("--- err ---\n", err)
}
