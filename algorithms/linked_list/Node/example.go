package main

import (
	dsa "algorithms/linked_list/Node/src"
	"fmt"
)

func main() {
	list1 := &dsa.Node{0, nil}
	p1 := list1
	for i := 1; i < 10; i++ {
		p1.Next = &dsa.Node{i, nil}
		p1 = p1.Next
	}

	fmt.Println(list1.String())

	list2 := &dsa.DLNode{0, nil, nil}
	p2 := list2
	for i := 1; i < 10; i++ {
		p2.Next = &dsa.DLNode{i, nil, nil}
		p2.Next.Prev = p2
		p2 = p2.Next
	}

	fmt.Println(list2.String())
	fmt.Println(list2.InvString())
}
