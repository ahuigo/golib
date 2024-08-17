package demo

import (
	"io"
	"testing"
)

type Animal interface {
	Name() string
	Sing()
}

type Worker interface {
	DoWork()
}

type Monkey struct {
	name string
	Animal
}

func (m Monkey) DoWork()      { println("No! I don't!") }
func (m Monkey) Sing()        { println("Screech!") }
func (m Monkey) Name() string { return m.name }

type Person struct {
	name string
	// embed interface1: Animal
	Animal

	// releaseOnce sync.Once
	// reqBuf      *bytes.Buffer
	// embed interface2
	io.ReadCloser
}

func (p Person) DoWork()      { println("Yes, Sir!") }
func (p Person) Sing()        { println("LaLaLa!") }
func (p Person) Name() string { return p.name }

type Employee struct {
	Animal
	Worker
}

func TestEmbedStruct(t *testing.T) {
	// Employee 和 interface(Animal/Worker) 没有实现任何方法
	// 1. Employee embeds Person:
	p := Person{name: "Alex"}
	e := Employee{Animal: p, Worker: p}
	print("I am ", e.Name(), ": ")
	e.DoWork()
	e.Sing()

	// 2. Employee embeds Monkey
	m := Monkey{name: "monkey"}
	e = Employee{Animal: m, Worker: m}
	print("I am ", e.Name(), ": ")
	e.DoWork()
	e.Sing()
}
