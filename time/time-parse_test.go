package main

import (
	"fmt"
	"os"
	"testing"
	"time"
)

func TestParse(t *testing.T) {
	println("chang zone to Africa...")
	os.Setenv("TZ", "Europe/Berlin")

	t1, _ := time.Parse("2006-01-02 15:04:05Z07", "2012-02-03 09:09:41+08")
	//t1, _ := time.Parse("2006-01-02 15:04:05Z07:00", "2012-02-03 09:09:41+08:00")
	//t1, _ := time.Parse(time.RFC3339, "2012-02-03T9:09:41+01:00")
	fmt.Println("time:", t1)
	fmt.Println("RFC3339:", time.RFC3339)
}
