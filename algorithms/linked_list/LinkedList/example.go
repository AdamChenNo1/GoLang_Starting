package main

import (
	LinkedList "algorithms/linked_list/LinkedList/src"
	"fmt"
)

func main() {
	// list1 := LinkedList.NewSGLinkedList()
	// for i := 0; i < 10; i++ {
	// 	list1.InsertLast(i)
	// }
	// fmt.Println(list1.String())

	// list2 := LinkedList.NewDLLinkedList()
	// for i := 0; i < 10; i++ {
	// 	list2.InsertLast(i)
	// }
	// fmt.Println(list2.String())
	// fmt.Println(list2.InvString())
	// list3 := LinkedList.NewSGLinkedListFromSlice([]int{1, 2, 3, 4, 5, 6}...)
	// list3 := LinkedList.NewSGLinkedListFromSlice(1, 2, 3, 4, 5, 6)
	// list3 := LinkedList.NewSGLinkedListFromSlice([]interface{}{"hello", "world", "this", "is", "adam"}...)
	// fmt.Println(list3.String())
	// list4 := LinkedList.NewDLLinkedListFromSlice([]interface{}{"hello", "world", "this", "is", "adam"}...)
	// fmt.Println(list4.String())
	// fmt.Println(list4.InvString())
	datas := []struct {
		num    []int
		k      int
		String string
	}{
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			9,
			"2->3->4->5->6->7->8->9->nil"},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			2,
			"1->2->3->4->5->6->7->9->nil"},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			1,
			"1->2->3->4->5->6->7->8->nil"},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			4,
			"1->2->3->4->5->7->8->9->nil"},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			5,
			"1->2->3->4->6->7->8->9->nil"},
	}
	for _, data := range datas {
		length := len(data.num)
		nums := make([]interface{}, 0, length)
		for _, num := range data.num {
			nums = append(nums, num)
		}
		list := LinkedList.NewDLLinkedListFromSlice(nums...)
		len1 := list.Size()
		fmt.Println("删除第", data.k, "个元素前：", list.String(), "大小：", len1)
		list.DeleteReverseK(data.k)
		len2 := list.Size()
		fmt.Println("删除第", data.k, "个元素后：", list.String(), "大小：", len2)
		fmt.Println("-------------------------------------")
	}
}
