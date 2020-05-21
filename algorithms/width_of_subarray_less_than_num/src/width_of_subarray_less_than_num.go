package SumOfSubarray

import "math"

func Sum_of_subarray_less_than_num(number []int, num int) int {
	if num < 0 {
		panic("num must be non-negative")
	}
	if number == nil {
		panic("wrong type array")
	}
	return getSum_of_subarray_less_than_num(number, num)
}

func getSum_of_subarray_less_than_num(number []int, num int) int {
	max := number[0]
	min := number[0]
	length := len(number)
	for i := 1; i < length; i++ {
		a := number[i]
		if a > max {
			max = a
		}
		if a < min {
			min = a
		}
	}

	countMax := length //数组中与最大值之差小于num的值个数
	countMin := length //数组中与最小值之差小于num的值个数
	between := 0       //数组中介于最小值+num和最大值-num之间的元素个数
	threshold1 := max - num
	threshold2 := min + num
	for j := 0; j < length; j++ {
		b := number[j]
		if b < threshold1 {
			countMax--
		} else if b <= threshold2 {
			between++
		} else {
			countMin--
		}
	}
	return int(math.Pow(2, float64(countMax)) + math.Pow(2, float64(countMin)) - math.Pow(2, float64(between)))
}
