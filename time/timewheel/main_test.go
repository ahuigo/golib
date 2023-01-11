package timer

import (
	"fmt"
	"testing"
	"time"
)

const CheckInterval = 1 * time.Second

func task(name string) {
	fmt.Println(time.Now(), "run task:", name)
}
func TestMain(m *testing.T) {
	tw := NewTimer()
	tw.CronFunc("*/2 * * * *", func() {
		task("task2")
	})

	/*
		timer := tw.ScheduleFunc(CheckInterval*2, func() {
			task("task1")
		})
		tw.ScheduleFunc(CheckInterval*3, func() {
			task("task2")
		})
	*/

	// // delete topic
	// s.runningAPIMap.Delete(topic)

	// run timer in StartMonitorSrv
	go tw.Run()

	time.Sleep(time.Minute * 500)
	// timer.Stop()
	time.Sleep(time.Second * 50)
}
