package main

import (
	dsa "algorithms/width_of_subarray_less_than_num/src"
	"fmt"
)

func main() {
	// dsa.Hannoi(4, -1, 1)
	// dsa.Hanoi_v2(4, 0, -1)
	// numbers := [][]int{{1}, {0}, {1}, {1}}
	numbers1 := []int{1, 2, 4, 5}
	// numbers2 := []int{1, 0, 1, 1}
	slice := dsa.Sum_of_subarray_less_than_num(numbers1, 3)
	// slice := dsa.Sum_of_subarray_less_than_num(numbers2, 1)
	fmt.Println(slice)
}
