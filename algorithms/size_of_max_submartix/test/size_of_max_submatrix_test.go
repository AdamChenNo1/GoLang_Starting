package MaxSubmatrixSize

import (
	dsa "algorithms/size_of_max_submartix/src"
	"testing"
)

func TestSize_of_max_submatrix(t *testing.T) {
	var arrays = []struct {
		numbers [][]int
		res     int
	}{
		{
			[][]int{{1, 0, 1, 1}},
			2,
		}, {
			[][]int{{1}, {0}, {1}, {1}},
			2,
		},
		{
			[][]int{{1, 0, 0, 1}, {0, 1, 0, 1}, {1, 0, 1, 0}, {1, 1, 1, 1}},
			4,
		},
	}

	for _, pair := range arrays {
		numbers := pair.numbers
		if dsa.Size_of_max_submatrix(numbers) != pair.res {
			t.Error("Size_of_max_submatrix failed")
		}
	}
	t.Log("Size_of_max_submatrix succeed")
}
