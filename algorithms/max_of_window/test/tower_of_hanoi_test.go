package maxOfWindow

import (
	dsa "algorithms/max_of_window/src"
	"reflect"
	"testing"
)

func TestMax_of_window(t *testing.T) {
	var arrays = []struct {
		numbers []int
		maxes   []int
	}{
		{
			[]int{1, 3, 5, 7, 9, 4, 3, 4, 6, 7},
			[]int{5, 7, 9, 9, 9, 4, 6, 7},
		}, {
			[]int{4, 3, 5, 4, 3, 3, 6, 7},
			[]int{5, 5, 5, 4, 6, 7},
		},
	}

	for _, array := range arrays {
		numbers := array.numbers
		if !reflect.DeepEqual(dsa.Max_of_window(numbers, 3), array.maxes) {
			t.Error("Max_of_window failed")
		}
	}
	t.Log("Max_of_window succeed")
}
