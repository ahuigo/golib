package main

import (
    "bytes"
    "os"
    "os/exec"
)

func main() {
    var stderr bytes.Buffer

    cmd := exec.Command("cat")
    cmd.Stdin = bytes.NewBuffer([]byte("hi"))
    cmd.Stdout = os.Stdout
	cmd.Stderr = &stderr

    cmd.Run()

}
