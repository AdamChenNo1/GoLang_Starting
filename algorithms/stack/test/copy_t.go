package main

import (
	dsa "algorithms/stack/src"
	"fmt"
)

func main() {
	// a := []string{"jacl", "mary", "safas"}
	// b := a[0 : len(a)-1]
	// a = []string{}
	// copy(a, b)
	// fmt.Println(a, cap(a), b, cap(b))
	a := new(dsa.ArrayStack)
	fmt.Println(a.Size())
	a.Push(1)
	a.Push(2)
	a.Pop()
	fmt.Println(a.Top())
}
