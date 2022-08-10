package process

import (
	"fmt"
	"os/exec"
	"syscall"
	"testing"
	"time"
)

// refer: https://medium.com/@felixge/killing-a-child-process-and-all-of-its-children-in-go-54079af94773
//通过pgid 会kill grandchild: ps ajx | ag sleep
func TestKillGroup(t *testing.T) {
	// cmd := exec.Command("sh", "-c", "sleep 52 ")
	cmd := exec.Command("sh", "-c", "for i in `seq 1 100`;do sleep 10;date >>date2.log; done")
	// 为避免父进程suuicide, 子进程生成新的pgid
	// call setpgid(2) between fork(2) and execve(2), to assign a new PGID identical to its PID.
	cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}

	start := time.Now()
	time.AfterFunc(100*time.Second, func() {
		// 子进程与孙子进程所在的pgid 是pid, 负数代表杀死组
		syscall.Kill(-cmd.Process.Pid, syscall.SIGKILL)
	})
	err := cmd.Run()
	fmt.Printf("pid=%d duration=%s err=%v\n", cmd.Process.Pid, time.Since(start), err)
}
