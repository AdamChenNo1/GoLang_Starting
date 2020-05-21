package main

import (
	LinkedList "algorithms/linked_list/LinkedList/src"
	"fmt"
)

func main() {
	list1 := LinkedList.NewSGLinkedList()
	for i := 1; i < 10; i++ {
		list1.InsertLast(i)
	}
	fmt.Println("删除前：", list1.String())
	list1.DeleteReverseK(9)
	fmt.Println("----------------------")
	fmt.Println("删除后：", list1.String())
	list2 := LinkedList.NewDLLinkedList()
	for i := 1; i < 10; i++ {
		list2.InsertLast(i)
	}
	fmt.Println("----------------------")
	fmt.Println("删除前：", list2.String(), "大小：", list2.Size())
	fmt.Println("逆序：", list2.InvString(), "大小：", list2.Size())
	fmt.Println("----------------------")
	list2.DeleteReverseK(1)
	fmt.Println("删除后：", list2.String(), "大小：", list2.Size())
	fmt.Println("逆序：", list2.InvString(), "大小：", list2.Size())
}
