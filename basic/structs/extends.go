package structs

import "fmt"

type Human struct {
	Name string
	Age int
	Phone string
}

type Student struct {
	Human //匿名字段
	School string
}

type Employee struct {
	Human //匿名字段
	Company string
}

//在human上面定义了一个method
func (h *Human) SayHi() {
	fmt.Printf("Hi, I am %s you can call me on %s\n", h.Name, h.Phone)
}

