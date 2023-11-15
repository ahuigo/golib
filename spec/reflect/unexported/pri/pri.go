package pri

type Obj struct {
	val  int
	Name string
}

func New() *Obj {
	return &Obj{val: 1, Name: "hello"}
}
