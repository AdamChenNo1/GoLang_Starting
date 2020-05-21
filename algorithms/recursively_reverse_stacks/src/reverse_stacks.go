package stack

import dsa "algorithms/stack/src"

func Reverse(s *dsa.ArrayStack) {
	if s.Size() > 1 {
		bottle := getBottle(s)
		Reverse(s)
		s.Push(bottle)
	}
}

func getBottle(s *dsa.ArrayStack) interface{} {
	if s.Size() == 1 {
		return s.Pop()
	} else {
		v := s.Pop()
		bottle := getBottle(s)
		s.Push(v)
		return bottle
	}
}
