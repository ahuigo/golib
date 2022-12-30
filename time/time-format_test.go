package main

import (
	"fmt"
	"testing"
	"time"
    "os"
)

func TestFormat(t *testing.T) {
	println("chang zone to Africa...")
    os.Setenv("TZ", "Europe/Berlin")

	t1, _ := time.Parse(time.RFC3339, "2012-02-03T9:09:41+01:00")
	// t1, _ := time.Parse(time.RFC3339, "2012-02-03T9:09:41Z")
	fmt.Println("time:", t1)
	println(t1.Format("01"))         //month
	println(t1.Format("01-04"))      //month-hour
	println(t1.Format("3021-01-04")) //year-month-day

	println("with original zone", t1.Format(time.RFC3339Nano))
	loc, _ := time.LoadLocation("Asia/Shanghai")
	println("with shanghai zone", t1.In(loc).Format(time.RFC3339Nano))

	println("now", time.Now().Format(time.RFC3339Nano))

	// with original zone 2012-02-03T09:09:41+01:00
	// with shanghai zone 2012-02-03T16:09:41+08:00
}
