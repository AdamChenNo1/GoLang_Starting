package maxOfWindow

import "container/list"

func Max_of_window(number []int, window int) []int {
	if len(number) < window {
		panic("array size must be not less than window size")
	} else {
		return getMaxList(number, window)
	}
}

func getMaxList(number []int, window int) []int {
	l := list.New()
	length := len(number)
	size := length - window + 1
	res := make([]int, 0, size)
	for i := 0; i < length; i++ {
		if l.Len() == 0 {
			l.PushBack(i)
		} else {
			value := number[i]
			Back := l.Back()
			if value < number[Back.Value.(int)] {
				l.PushBack(i)
			} else {
				for back := Back; l.Len() != 0 && value >= number[back.Value.(int)]; back = l.Back() {
					l.Remove(back)
				}
				l.PushBack(i)
			}

			front := l.Front()
			if front.Value.(int) == i-window {
				l.Remove(front)
			}
		}

		if i >= window-1 {
			res = append(res, number[l.Front().Value.(int)])
		}
	}
	return res
}
