package main

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
func (m Monkey) DoWork() { println("No! I don't!") }
func (m Monkey) Sing() { println("Screech!") }
func (m Monkey) Name() string { return m.name }


type Person struct {
	name string
	Animal
}
func (p Person) DoWork() { println("Yes, Sir!") }
func (p Person) Sing() { println("LaLaLa!") }
func (p Person) Name() string { return p.name }

type Employee struct {
	Animal
	Worker
}

func main() {
	//e := Employee{Person{name:"mattn"}}
	p := Person{name:"mattn"}
	e := Employee{Animal:p, Worker:p}
	print("I am ", e.Name(), ": ")
	e.DoWork()
	e.Sing()

	m := Monkey{name:"george"}
	e = Employee{Animal:m, Worker:m}
	print("I am ", e.Name(), ": ")
	e.DoWork()
	e.Sing()
}

