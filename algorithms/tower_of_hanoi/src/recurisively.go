package tower_of_hanoi

import "fmt"

func Hanoi(n int, from int, to int) int {
	if n < 1 {
		panic("n must be greater than 0")
	} else if from > 1 || from < -1 {
		panic("from must be -1|0|1")
	} else if to > 1 || to < -1 {
		panic("from must be -1|0|1")
	} else {
		total := move(n, from, to)

		fmt.Printf("total steps: %d\n", total)
		return total
	}
}

//打印移动的步骤
func printStep(n int, from int, to int) {
	dic := map[int]string{-1: "left", 0: "mid", 1: "right"}
	fmt.Printf("Move %d "+"from %s "+"to %s \n", n, dic[from], dic[to])
}

//移动底层，返回需要的步数和上面n-1层的位置
func moveBottle(n int, from int, to int) (steps int, begin int) {
	total := 0
	var begins int
	if n == 1 {
		total += moveN(1, from, to)
		begins = 0
	} else {
		var other int
		if from == 0 {
			other = from - to
		} else if to == 0 {
			other = to - from
		} else {
			other = 0
		}
		if other != 0 {
			total += move(n-1, from, other)
			total += moveN(n, from, to)
			begins = other
		} else {
			total += move(n-1, from, to)
			total += moveN(n, from, other)
			total += move(n-1, to, from)
			total += moveN(n, other, to)
			begins = from
		}
	}
	return total, begins
}

//移动第N层，返回需要的步数
func moveN(n int, from int, to int) int {
	if from == 0 || to == 0 {
		printStep(n, from, to)
		return 1
	} else {
		printStep(n, from, 0)
		printStep(n, 0, to)
		return 2
	}
}

//移动从1层开始到n层的整体,返回总步数
func move(n int, from int, to int) int {
	total := 0
	for j := n; j > 0; j-- {
		steps, begins := moveBottle(j, from, to)
		total += steps
		from = begins
	}
	return total
}
