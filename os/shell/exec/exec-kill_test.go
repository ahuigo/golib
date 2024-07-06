package demo

import (
	"fmt"
	"os/exec"
	"testing"
	"time"
)

func TestExec(t *testing.T) {
	cmd := exec.Command("sh", "-c", "for i in `seq 1 100`;do sleep 5;date >>date2.log; done")
	start := time.Now()
	time.AfterFunc(1*time.Second, func() { cmd.Process.Kill() })
	err := cmd.Run()
	fmt.Printf("pid=%d duration=%s err=%s\n", cmd.Process.Pid, time.Since(start), err)
}
