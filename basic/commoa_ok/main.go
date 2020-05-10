package commoa_ok

import (
	"strconv"
)

type Element interface{}
type List [] Element

type Person struct {
	Name string
	Age int
}

//打印
func (p Person) String() string {
	return "(name: " + p.Name + " - age: "+strconv.Itoa(p.Age)+ " years)"
}

