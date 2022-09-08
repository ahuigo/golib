package main

import (
	"testing"
	"time"
)

func TestRaceSlice(t *testing.T) {
	type Obj struct {
		count int
	}

	s := []*Obj{
		{1},
		{2},
	}
	// writer
	go func() {
		i := 1
		for {
			i += 1
			s = []*Obj{
				{1},
				{2},
			}
			time.Sleep(10 * time.Nanosecond)
			//time.Sleep(10*time.Second)
		}
	}()

	// read
	for {
		s1 := s
		if s1[1].count != 2 {
			panic("error")
		}
		time.Sleep(10 * time.Nanosecond)
		// time.Sleep(10 * time.Millisecond)
	}
}
