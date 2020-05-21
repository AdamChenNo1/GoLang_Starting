package main

import (
	dsa "algorithms/size_of_max_submartix/src"
	"fmt"
)

func main() {
	// dsa.Hannoi(4, -1, 1)
	// dsa.Hanoi_v2(4, 0, -1)
	// numbers := [][]int{{1}, {0}, {1}, {1}}
	numbers1 := [][]int{{1, 0, 0, 1}, {0, 1, 0, 1}, {1, 0, 1, 0}, {1, 1, 1, 1}}
	slice := dsa.Size_of_max_submatrix(numbers1)
	fmt.Println(slice)
}
