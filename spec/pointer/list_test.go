package demo

import (
	"testing"
	"unsafe"

	"github.com/antlabs/stl/list"
)

func TestListForEachPtr2(t *testing.T) {
	type timeNode struct {
		expire uint64
		list.Head
	}
	type Time struct {
		timeNode
	}
	head := &Time{}
	{
		head.Init()
		node1 := &timeNode{expire: 1}
		node2 := &timeNode{expire: 4}
		head.AddTail(&node1.Head)
		head.AddTail(&node2.Head)
	}

	offset := unsafe.Offsetof(head.Head)
	head.ForEachSafe(func(pos *list.Head) {
		val := (*timeNode)(pos.Entry(offset))
		t.Logf("expire:%d", val.expire)
	})
}
