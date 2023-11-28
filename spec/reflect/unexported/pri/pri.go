package pri

type Stu struct {
	Name string
	age  int
}
type obj struct {
	val  int
	stu  *Stu
	Name string
}

func New() *obj {
	return &obj{
		val:  1,
		Name: "obj1",
		stu:  &Stu{Name: "Alex", age: 18},
	}
}
