package main

import (
	dsa "algorithms/min_ele_stack/src/v2"
	"fmt"
	"reflect"
)

func main() {
	a := new(dsa.MinEleStack)
	for i := 1; i < 10; i++ {
		a.Push(i)
	}
	fmt.Println(a.Min())
	// printType(1)
}

func printType(v interface{}) {
	types := reflect.TypeOf(v).Kind()
	fmt.Println(types)
}
