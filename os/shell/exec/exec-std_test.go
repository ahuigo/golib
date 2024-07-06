package demo

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"testing"
)

func TestExecEnvs(t *testing.T) {
	out, msg, errno := ExecCommandEnvs([]string{"A=1"}, "sh", "-c", "echo A is $A; no-existed-cmd")
	t.Log(out, msg, errno)
}

func ExecCommandEnvs(envs []string, command string, args ...string) (output string, errmsg string, errno int) {
	var stderr bytes.Buffer
	fmt.Printf("ExecCommand: %s %s \n", command, strings.Join(args, " "))
	cmd := exec.Command(command, args...)
	cmd.Env = append(os.Environ(), envs...)
	cmd.Stderr = &stderr
	// err := cmd.Run() // err: exit status 1
	outputBytes, err := cmd.Output()
	output = string(outputBytes)
	errmsg = stderr.String()
	if err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok {
			if status, ok := exitErr.Sys().(interface{ ExitStatus() int }); ok {
				errno = status.ExitStatus()
			}
		} else if execErr, ok := err.(*exec.Error); ok {
			errmsg = execErr.Error() + " " + errmsg
			errno = 127
		}
	}
	return
}

func ExecCommandPipe(cmd string, stdin []byte, args ...string) (output []byte, errmsg string, errno int) {
	var stdout, stderr bytes.Buffer
	var err error
	command := exec.Command(cmd, args...)
	if pipe, err := command.StdinPipe(); err != nil {
		return nil, err.Error(), 1
	} else {
		if _, err = pipe.Write(stdin); err != nil {
			return nil, err.Error(), 1
		}
	}
	command.Stdout = &stdout
	command.Stderr = &stderr
	err = command.Run() // err: exit status 1
	output = bytes.TrimSpace(stdout.Bytes())
	errmsg = strings.TrimSpace(stderr.String())
	if err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok {
			if status, ok := exitErr.Sys().(interface{ ExitStatus() int }); ok {
				errno = status.ExitStatus()
			}
		}
	}
	return
}
