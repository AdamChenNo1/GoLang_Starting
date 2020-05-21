package main

import (
	algo "algorithms/recursively_reverse_stacks/src"
	dsa "algorithms/stack/src"
	"testing"
)

func TestReverse(t *testing.T) {
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
		t.Log("stack reverse succeed")
	} else {
		t.Error("stack reverse failed")
	}
}
