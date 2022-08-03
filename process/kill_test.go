package test

import (
	"fmt"
	"os/exec"
	"testing"
	"time"
)

func TestKill(t *testing.T) {
	cmd := exec.Command("sh", "-c", "watch date > date.log")
	start := time.Now()
	time.AfterFunc(1*time.Second, func() { cmd.Process.Kill() })
	err := cmd.Run()
	fmt.Printf("pid=%d duration=%s err=%s\n", cmd.Process.Pid, time.Since(start), err)
}

//不会kill grandchild: ps aux|grep watch; tail -f date.log
