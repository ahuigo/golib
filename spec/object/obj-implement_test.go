package demo

import (
	"fmt"
	"testing"
)

type InterfaceA interface {
	Play()
}

func play(o InterfaceA) {
	o.Play()
}

// Pet 要实现implement InterfaceA接口
type Pet struct {
	age int
}

func (p *Pet) Play() {
	fmt.Println("showc age:", p.age)
}

/*
*
本例实现了一个InterfaceA.Play接口。

最著名的就是 io.Read 和 ioutil.ReadAll 的玩法，其中 io.Read 是一个接口，你需要实现他的一个 Read(p []byte) (n int, err error) 接口方法，只要满足这个规模，就可以被 ioutil.ReadAll这个方法所使用。这就是面向对象编程方法的黄金法则——“Program to an interface not an implementation”
*/
func TestObjImplement(t *testing.T) {
	//pf := fmt.Printf

	play(&Pet{20}) //ok
	//pet.Play()

}
