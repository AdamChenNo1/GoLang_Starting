package SumOfSubarray

import (
	dsa "algorithms/width_of_subarray_less_than_num/src"
	"testing"
)

func TestSize_of_max_submatrix(t *testing.T) {
	var arrays = []struct {
		numbers []int
		num     int
		res     int
	}{
		{
			[]int{1, 0, 1, 1},
			1,
			16,
		}, {
			[]int{1, 2, 4, 5},
			3,
			12,
		},
		{
			[]int{1, 2, 4, 6},
			3,
			10,
		},
	}

	for _, pair := range arrays {
		numbers := pair.numbers
		num := pair.num
		if dsa.Sum_of_subarray_less_than_num(numbers, num) != pair.res {
			t.Error("Sum_of_subarray_less_than_num failed")
		}
	}
	t.Log("Sum_of_subarray_less_than_num succeed")
}
