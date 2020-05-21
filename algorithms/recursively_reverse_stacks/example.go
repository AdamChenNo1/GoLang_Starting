package main

import (
	algo "algorithms/recursively_reverse_stacks/src"
	dsa "algorithms/stack/src"
	"fmt"
)

func main() {
	stack := new(dsa.ArrayStack)
	for i := 0; i < 5; i++ {
		stack.Push(i)
	}

	algo.Reverse(stack)
	var res [5]int
	for i := 0; i < 5; i++ {
		res[i] = stack.Pop().(int)
	}

	if res == [5]int{0, 1, 2, 3, 4} {
		fmt.Println("succeed")
	} else {
		fmt.Println(res)
		fmt.Println("failed")
	}
}
