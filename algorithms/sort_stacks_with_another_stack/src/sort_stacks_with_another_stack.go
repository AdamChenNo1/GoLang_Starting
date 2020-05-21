package sort_stacks_with_another_stack

import dsa "algorithms/stack/src"

func Sort_stack(stack *dsa.ArrayStack) {
	if stack.IsEmpty() || stack.Size() == 1 {
		return
	}
	help := new(dsa.ArrayStack)
	for stack.Size() != 0 {
		if ele, ok := stack.Pop().(int); !ok {
			panic("non int stack")
		} else {
			help.Push(ele)
		}
	}
	total := help.Size()
	cur := help.Pop().(int)
	for stack.Size() != total && !help.IsEmpty() {
		if htop := help.Pop().(int); stack.IsEmpty() || cur <= htop {
			stack.Push(cur)
			cur = htop
		} else {
			for !stack.IsEmpty() {
				if stop := stack.Top().(int); stop <= htop {
					break
				} else {
					help.Push(stack.Pop().(int))
				}
			}
			stack.Push(htop)
		}
	}
	stack.Push(cur)
}
