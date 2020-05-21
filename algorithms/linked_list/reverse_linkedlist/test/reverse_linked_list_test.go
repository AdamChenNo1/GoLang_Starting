package LinkedList

import (
	LinkedList "algorithms/linked_list/LinkedList/src"
	"testing"
)

func TestSGLinkedListReverse(t *testing.T) {
	datas := []struct {
		num    []int
		String string
	}{
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			"9->8->7->6->5->4->3->2->1->nil"},
		{[]int{1},
			"1->nil"},
	}
	for _, data := range datas {
		length := len(data.num)
		nums := make([]interface{}, 0, length)
		for _, num := range data.num {
			nums = append(nums, num)
		}
		list := LinkedList.NewSGLinkedListFromSlice(nums...)
		list.Reverse()
		if list.String() != data.String {
			t.Error("Reverse failed")
		} else {
			t.Log("Reverse succeed")
		}
	}
}
func TestDLLinkedListReverse(t *testing.T) {
	datas := []struct {
		num    []int
		String string
	}{
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			"9->8->7->6->5->4->3->2->1->nil"},
		{[]int{1},
			"1->nil"},
	}
	for _, data := range datas {
		length := len(data.num)
		nums := make([]interface{}, 0, length)
		for _, num := range data.num {
			nums = append(nums, num)
		}
		list := LinkedList.NewDLLinkedListFromSlice(nums...)
		list.Reverse()
		if list.String() != data.String {
			t.Error("Reverse failed")
		} else {
			t.Log("Reverse succeed")
		}
	}
}
