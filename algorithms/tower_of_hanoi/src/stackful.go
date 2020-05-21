package tower_of_hanoi

import (
	dsa "algorithms/stack/src"
	"fmt"
)

func Hanoi_v2(n int, from int, to int) int {
	if n < 1 {
		panic("n must be greater than 0")
	} else if from > 1 || from < -1 {
		panic("from must be -1|0|1")
	} else if to > 1 || to < -1 {
		panic("from must be -1|0|1")
	} else {
		var maps map[int]string
		if from == -1 || to == 1 {
			maps = map[int]string{-1: "left", 0: "mid", 1: "right"}
		} else {
			maps = map[int]string{-1: "right", 0: "mid", 1: "left"}
		}
		if to < 0 {
			to = -to
		}

		total := move_v2(maps, n, from, to)
		fmt.Printf("total steps: %d\n", total)
		return total
	}
}

func printStepWithMap(maps map[int]string, n int, from int, to int) {
	fmt.Printf("Move %d "+"from %s "+"to %s \n", n, maps[from], maps[to])
}

func move_v2(maps map[int]string, n int, from int, to int) int {
	stacks := InitStack(n, from)
	total := 0
	LS := stacks[-1]
	MS := stacks[0]
	RS := stacks[1]
	if n == 1 {
		if from == 0 || to == 0 {
			printStepWithMap(maps, 1, from, to)
			total = 1
		} else {
			printStepWithMap(maps, 1, from, 0)
			printStepWithMap(maps, 1, 0, to)
			total = 2
		}
	} else {
		for stacks[to].Size() != n {
			if LS.IsEmpty() {
				if MS.Size()%2 == 0 {
					value := MS.Pop()
					LS.Push(value)
					printStepWithMap(maps, value.(int), 0, -1)
					total++
				} else if !RS.IsEmpty() && MS.Top().(int) > RS.Top().(int) {
					value := RS.Pop()
					MS.Push(value)
					printStepWithMap(maps, value.(int), 1, 0)
					total++
				} else {
					value := MS.Pop()
					RS.Push(value)
					printStepWithMap(maps, value.(int), 0, 1)
					total++
				}
			} else if MS.IsEmpty() {
				if RS.IsEmpty() || LS.Top().(int) < RS.Top().(int) {
					value := LS.Pop()
					RS.Push(value)
					printStepWithMap(maps, value.(int), -1, 0)
					printStepWithMap(maps, value.(int), 0, 1)
					total += 2
				} else {
					value := LS.Pop()
					MS.Push(value)
					printStepWithMap(maps, value.(int), -1, 0)
					total++
				}
			} else if RS.IsEmpty() {
				if MS.Size()%2 == 0 {
					if LS.Top().(int) > MS.Top().(int) {
						value := MS.Pop()
						LS.Push(value)
						printStepWithMap(maps, value.(int), 0, -1)
						total++
					} else {
						value := LS.Pop()
						MS.Push(value)
						printStepWithMap(maps, value.(int), -1, 0)
						total++
					}
				} else {
					value := MS.Pop()
					RS.Push(value)
					printStepWithMap(maps, value.(int), 0, 1)
					total++
				}
			} else {
				if MS.Size()%2 == 1 {
					if MS.Top().(int) < RS.Top().(int) {
						value := MS.Pop()
						RS.Push(value)
						printStepWithMap(maps, value.(int), 0, 1)
						total++
					} else {
						value := RS.Pop()
						MS.Push(value)
						printStepWithMap(maps, value.(int), 1, 0)
						total++
					}
				} else {
					if LS.Top().(int) < MS.Top().(int) {
						value := LS.Pop()
						MS.Push(value)
						printStepWithMap(maps, value.(int), -1, 0)
						total++
					} else {
						value := MS.Pop()
						LS.Push(value)
						printStepWithMap(maps, value.(int), 0, -1)
						total++
					}
				}
			}
		}
	}

	return total
}

func InitStack(n int, from int) map[int]*dsa.ArrayStack {
	stack := new(dsa.ArrayStack)
	for i := n; i > 0; i-- {
		stack.Push(i)
	}
	LS := new(dsa.ArrayStack)
	MS := new(dsa.ArrayStack)
	RS := new(dsa.ArrayStack)
	if from == 0 {
		MS = stack
	} else {
		LS = stack
	}
	return map[int]*dsa.ArrayStack{-1: LS, 0: MS, 1: RS}
}
