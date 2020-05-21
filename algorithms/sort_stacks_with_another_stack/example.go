package main

import (
	algo "algorithms/sort_stacks_with_another_stack/src"
	dsa "algorithms/stack/src"
	"fmt"
	"math/rand"
)

func main() {
	stack := new(dsa.ArrayStack)
	for i := 0; i < 10; i++ {
		stack.Push(rand.Intn(100))
	}

	algo.Sort_stack(stack)
	fmt.Println(stack.Size())
	for !stack.IsEmpty() {
		j := stack.Pop().(int)
		fmt.Println(j)
	}
}
