package LinkedList

import (
	LinkedList "algorithms/linked_list/LinkedList/src"
	"testing"
)

func TestSGLinkedListReversePart(t *testing.T) {
	datas := []struct {
		num      []int
		from, to int
		String   string
	}{
		{
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			1, 3,
			"1->4->3->2->5->6->7->8->9->nil",
		},
		{
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			-4, 15,
			"9->8->7->6->5->4->3->2->1->nil",
		},
	}
	for _, data := range datas {
		length := len(data.num)
		nums := make([]interface{}, 0, length)
		for _, num := range data.num {
			nums = append(nums, num)
		}
		list := LinkedList.NewSGLinkedListFromSlice(nums...)
		list.ReversePart(data.from, data.to)
		if list.String() != data.String {
			t.Error("ReversePart failed")
		} else {
			t.Log("ReversePart succeed")
		}
	}
}
func TestDLLinkedListReverse(t *testing.T) {
	datas := []struct {
		num      []int
		from, to int
		String   string
	}{{
		[]int{1, 2, 3, 4, 5, 6, 7, 8, 9},
		1, 3,
		"1->4->3->2->5->6->7->8->9->nil",
	},
		{
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			-4, 15,
			"9->8->7->6->5->4->3->2->1->nil",
		},
	}
	for _, data := range datas {
		length := len(data.num)
		nums := make([]interface{}, 0, length)
		for _, num := range data.num {
			nums = append(nums, num)
		}
		list := LinkedList.NewDLLinkedListFromSlice(nums...)
		list.ReversePart(data.from, data.to)
		if list.String() != data.String {
			t.Error("ReversePart failed")
		} else {
			t.Log("ReversePart succeed")
		}
	}
}
