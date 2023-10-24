package demo

import (
	"fmt"
	"testing"
	"unsafe"
)

func TestPtrOffset(t *testing.T) {
	a := [4]int{0, 1, 2, 3}
	p1 := unsafe.Pointer(&a[1])
	p3 := unsafe.Pointer(uintptr(p1) + 2*unsafe.Sizeof(a[0]))
	*(*int)(p3) = 6
	fmt.Println("a =", a) // a = [0 1 2 6]

	type Person struct {
		name   string
		age    int
		gender bool
	}

	who := Person{"John", 30, true}
	pp := unsafe.Pointer(&who)
	pname := (*string)(unsafe.Pointer(uintptr(pp) + unsafe.Offsetof(who.name)))
	page := (*int)(unsafe.Pointer(uintptr(pp) + unsafe.Offsetof(who.age)))
	pgender := (*bool)(unsafe.Pointer(uintptr(pp) + unsafe.Offsetof(who.gender)))
	*pname = "Alice"
	*page = 28
	*pgender = false
	fmt.Println(who) // {Alice 28 false}
}

func TestFindNodeByMemberNonPtr(t *testing.T) {
	type Head struct {
		Next *Head
	}
	type Node struct {
		name string //16
		age  int64  //8
		Head
	}

	node1 := &Node{age: 3}
	head := &Node{
		name: "headNonPtr",
		Head: Head{
			// Next: &Head{}, // 不能这样
			Next: &node1.Head,
		},
	}

	offset := unsafe.Offsetof(head.Head) // 16+8=24
	// 获取 node1.Head 的地址
	pos := head.Next // 如果用 &head.Next，会得到head.Head的地址
	node2 := (*Node)(unsafe.Pointer(uintptr(unsafe.Pointer(pos)) - offset))
	fmt.Println("node.age:", node2.age)
}

func TestFindNodeByMemberPtr(t *testing.T) {
	type Node struct {
		name string //16
		age  int64  //8
		next *Node
	}

	head := &Node{
		name: "head",
		next: &Node{
			name: "next",
		},
	}

	offset := unsafe.Offsetof(head.next)      // 16+8=24
	propPointer := unsafe.Pointer(&head.next) // 注意，不是next指针，而是head.next指针的地址
	t.Log("offset=", offset, "pos=", uintptr(propPointer))
	// t.Fatal("stop")
	node := (*Node)(unsafe.Pointer(uintptr(propPointer) - offset))
	fmt.Println("node.name=", node.name)
}
