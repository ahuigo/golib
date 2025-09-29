package main

import (
	"sync"

	"golang.org/x/sys/unix"
)

// 手动追踪脏页的结构体（只写脏页）
type DirtyTracker struct {
	pageSize int
	pages    []uint8 // 0/1 per page; 可改为位图以节省内存
	mu       sync.Mutex
	data     []byte // 映射的 []byte
}

func NewDirtyTracker(mapped []byte) *DirtyTracker {
	pageSize := unix.Getpagesize()
	pages := make([]uint8, (len(mapped)+pageSize-1)/pageSize)
	return &DirtyTracker{pageSize: pageSize, pages: pages, data: mapped}
}

// 在写入 mappedData 的封装函数里调用
func (t *DirtyTracker) MarkDirty(offset, length int) {
	if length <= 0 || offset < 0 || offset+length > len(t.data) {
		return
	}
	t.mu.Lock()
	defer t.mu.Unlock()
	start := offset / t.pageSize
	end := (offset + length + t.pageSize - 1) / t.pageSize
	for i := start; i < end; i++ {
		t.pages[i] = 1
	}
}

// 将脏页 flush 到磁盘，flags 可用 unix.MS_ASYNC 或 unix.MS_SYNC
func (t *DirtyTracker) Flush(flags int) error {
	t.mu.Lock()
	// 快速拷贝并清零标志，避免长时间持锁
	pagesCopy := make([]uint8, len(t.pages))
	copy(pagesCopy, t.pages)
	for i := range t.pages {
		t.pages[i] = 0
	}
	t.mu.Unlock()

	// 合并连续脏页区间并调用 msync
	for i := 0; i < len(pagesCopy); {
		if pagesCopy[i] == 0 {
			i++
			continue
		}
		j := i + 1
		for j < len(pagesCopy) && pagesCopy[j] == 1 {
			j++
		}
		// i..j-1 是连续的脏页
		start := i * t.pageSize
		end := j * t.pageSize
		if end > len(t.data) {
			end = len(t.data)
		}
		sub := t.data[start:end]
		if err := unix.Msync(sub, flags); err != nil {
			return err
		}
		i = j
	}
	return nil
}
