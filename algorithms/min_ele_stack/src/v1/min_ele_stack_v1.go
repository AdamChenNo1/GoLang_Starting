package min_ele_stack

import (
	dsa "algorithms/stack/src"
	"reflect"
)

type cmp interface {
	compare(v1 interface{}, v2 interface{}) int
}

type MinEleStack struct {
	data dsa.ArrayStack
	min  dsa.ArrayStack
}

func compare(v1 interface{}, v2 interface{}) int {
	type1 := reflect.TypeOf(v1).Kind()
	type2 := reflect.TypeOf(v2).Kind()
	if type1 != type2 {
		panic("different types")
	} else {
		value1 := v1.(int)
		value2 := v2.(int)
		if value1 < value2 {
			return -1
		} else if value1 > value2 {
			return 1
		} else {
			return 0
		}
	}
}

func (stack *MinEleStack) Push(v interface{}) {
	stack.data.Push(v)

	if stack.min.Size() == 0 || compare(v, stack.min.Top()) < 0 {
		stack.min.Push(v)
	}
}

func (stack *MinEleStack) Pop() interface{} {
	value := stack.data.Pop()
	if compare(value, stack.min.Top()) == 0 {
		stack.min.Pop()
	}
	return value
}

func (stack *MinEleStack) Min() interface{} {
	return stack.min.Top()
}
