package main

import (
	"fmt"
	"testing"
	"unsafe"

	"github.com/antlabs/stl/list"
)

type worker struct {
	list.Head //内嵌到业务结构体里面
	ID        int
}

func TestMain(t *testing.T) {
	workerHead := worker{} //声明
	workerHead.Init()      //初始化

	n1 := worker{ID: 1}
	n2 := worker{ID: 2}

	workerHead.AddTail(&n1.Head) //添加到尾部
	workerHead.AddTail(&n2.Head) //添加到尾部

	offset := unsafe.Offsetof(workerHead.Head)
	//遍历
	workerHead.ForEach(func(pos *list.Head) {

		worker := (*worker)(pos.Entry(offset))

		fmt.Printf("worker id:%d\n", worker.ID)
	})
}
