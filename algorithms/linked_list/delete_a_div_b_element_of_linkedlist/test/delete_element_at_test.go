package LinkedList

import (
	LinkedList "algorithms/linked_list/LinkedList/src"
	"testing"
)

func TestSGDeleteElementAt(t *testing.T) {
	datas := []struct {
		num    []int
		a, b   int
		String string
	}{
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			1, 9,
			"2->3->4->5->6->7->8->9->nil"},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			7, 8,
			"1->2->3->4->5->6->7->9->nil"},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			1, 1,
			"1->2->3->4->5->6->7->8->nil"},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			3, 5,
			"1->2->3->4->5->7->8->9->nil"},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			1, 2,
			"1->2->3->4->6->7->8->9->nil"},
	}
	for _, data := range datas {
		length := len(data.num)
		nums := make([]interface{}, 0, length)
		for _, num := range data.num {
			nums = append(nums, num)
		}
		list := LinkedList.NewSGLinkedListFromSlice(nums...)
		len1 := list.Size()
		list.DeleteElementAt(data.a, data.b)
		len2 := list.Size()
		if len1-len2 != 1 {
			t.Error("DeleteElementAt failed")
		} else if list.String() != data.String {
			t.Error("DeleteElementAt failed")
		} else {
			t.Log("DeleteElementAt succeed")
		}
	}
}

func TestDLDeleteElementAt(t *testing.T) {
	datas := []struct {
		num    []int
		a, b   int
		String string
	}{
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			1, 9,
			"2->3->4->5->6->7->8->9->nil"},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			7, 8,
			"1->2->3->4->5->6->7->9->nil"},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			1, 1,
			"1->2->3->4->5->6->7->8->nil"},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			3, 5,
			"1->2->3->4->5->7->8->9->nil"},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			1, 2,
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
		list.DeleteElementAt(data.a, data.b)
		len2 := list.Size()
		if len1-len2 != 1 {
			t.Error("DeleteElementAt failed")
		} else if list.String() != data.String {
			t.Error("DeleteElementAt failed")
		} else {
			t.Log("DeleteElementAt succeed")
		}
	}
}
