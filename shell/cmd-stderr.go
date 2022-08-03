package main

import (
    "bytes"
    "fmt"
    "log"
    "os/exec"
)

const ShellToUse = "bash"

func Shellout() (error, string, string) {
    var stdout bytes.Buffer
    var stderr bytes.Buffer
    //cmd := exec.Command("sh", "-c", "echo 'sleep 3' && sleep 3 && echo 'sleep end';cmd-not-existed||echo ok")
    args := []string{"sh","-c","exit 3"}
    cmd := exec.Command(args[0], args[1:]...)
    cmd.Stdout = &stdout
    cmd.Stderr = &stderr
    err := cmd.Run() // cmd.Run() 内含cmd.Wait()

    return err, stdout.String(), stderr.String()
}

func main() {
    err, out, errout := Shellout()
    if err != nil {
        if exitError, ok := err.(*exec.ExitError); ok {
            fmt.Printf("shell ExitCode: %d\n", exitError.ExitCode())
        }else{
            log.Printf("\ncmd.Run err=%v\n", err)
        }
    }
    fmt.Println("--- stdout ---")
    fmt.Println(out)
    fmt.Println("--- stderr ---")
    fmt.Println(errout)
}
