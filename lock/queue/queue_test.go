package queue

import (
	"testing"
)

func BenchmarkPushTailCAS(b *testing.B) {
	b.StopTimer()
	q := New()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		q.PushTailCAS(Message{id: i})
	}
}

func BenchmarkPushTailMutex(b *testing.B) {
	b.StopTimer()
	q := New()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		q.PushTailMutex(Message{id: i})
	}
}
