package main

import (
	"fmt"
	"sort"
)

func main() {
	values := []int32{2, 54, 8, 42, 6, 25, 47, 12}

	sort.Slice(values, func(i, j int) bool {
		return values[i] < values[j]
	})

	fmt.Println(values)
}
