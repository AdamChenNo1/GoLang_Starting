package main

import (
	Node "algorithms/linked_list/Node/src"
	dsa "algorithms/linked_list/common_part_of_linked_list/src"
)

func main() {
	list1 := &Node.Node{0, nil}
	p1 := list1
	for i := 1; i < 10; i++ {
		p1.Next = &Node.Node{i, nil}
		p1 = p1.Next
	}

	list2 := &Node.Node{1, nil}
	p2 := list2
	for i := 1; i < 10; i += 2 {
		p2.Next = &Node.Node{i, nil}
		p2 = p2.Next
	}

	dsa.PrintCommonPart(list1, list2)
}
