package main

import (
	dsa "algorithms/tower_of_hanoi/src"
)

func main() {
	// dsa.Hannoi(4, -1, 1)
	// dsa.Hanoi_v2(4, 0, -1)
	for i := 1; i < 4; i++ {
		dsa.Hanoi(i, 1, -1)
		dsa.Hanoi_v2(i, 1, -1)
	}
}
