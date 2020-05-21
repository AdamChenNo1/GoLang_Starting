package main

import (
	LinkedList "algorithms/linked_list/LinkedList/src"
	"fmt"
)

func main() {
	list := LinkedList.NewSGLinkedListFromSlice(1, 2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Println("Before:"+list.String()+" Size:", list.Size())
	list.DeleteElementAt(2, 9)
	fmt.Println("After:"+list.String()+" Size:", list.Size())
	list2 := LinkedList.NewDLLinkedListFromSlice(1, 2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Println("Before:"+list2.String()+" Size:", list2.Size())
	list2.DeleteElementAt(4, 8)
	fmt.Println("After:"+list2.String()+" Size:", list2.Size())
}
