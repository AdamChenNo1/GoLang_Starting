package structs

import "fmt"

//Employee的method重写Human的method
func (e *Employee) SayHi() {
	fmt.Printf("Hi, I am %s, I work at %s. Call me on %s\n", e.Name,
		e.Company, e.Phone) //Yes you can split into 2 lines here.
}

