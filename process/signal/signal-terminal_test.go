package test

// Refer: https://bigkevmcd.github.io/go/pgrp/context/2019/02/19/terminating-processes-in-go.html
import (
	"bytes"
	"context"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
	"testing"
)

func TestSignalTerminal(t *testing.T) {
    log.Println("start test...")
	defer func() {
		log.Println("uploading output")
	}()
	ctx, cancel := context.WithCancel(context.Background())
	cancelOnTermination(cancel)
	cmd := exec.CommandContext(ctx, "bash", "-c", "sleep 21 &")

	var outb, errb bytes.Buffer
	log.Printf("about to run - process has PID %d\n", os.Getpid())
	cmd.Stdout = &outb
	cmd.Stderr = &errb
	if err := cmd.Run(); err != nil {
		log.Printf("err was not nil: %v", err)
	}
	log.Println("Process terminated normally")
}

func cancelOnTermination(cancel context.CancelFunc) {
	log.Println("setting up a signal handler")
	s := make(chan os.Signal, 1)
	signal.Notify(s, syscall.SIGTERM)
	go func() {
		log.Printf("received SIGTERM %v\n", <-s)
		cancel()
	}()
}
