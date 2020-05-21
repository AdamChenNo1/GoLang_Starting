package main

import (
	dsa "algorithms/max_of_window/src"
	"fmt"
)

func main() {
	// dsa.Hannoi(4, -1, 1)
	// dsa.Hanoi_v2(4, 0, -1)
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 2, 4, 6, 8, 1}
	slice := dsa.Max_of_window(numbers, 3)
	fmt.Println(slice)
}
