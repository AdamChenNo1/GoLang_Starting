package LinkedList

import (
	LinkedList "algorithms/linked_list/LinkedList/src"
	"testing"
)

func TestSGLinkedList(t *testing.T) {
	list1 := LinkedList.NewSGLinkedList()
	for i := 0; i < 10; i++ {
		list1.InsertLast(i)
	}
	if list1.Size() != 10 {
		t.Error("SGLinkedList InsertLast failed")
	} else if list1.String() != "0->1->2->3->4->5->6->7->8->9->nil" {
		t.Error("SGLinkedList String failed")
	} else {
		t.Log("SGLinkedList succeed")
	}
}
func TestDLLinkedList(t *testing.T) {
	list2 := LinkedList.NewDLLinkedList()
	for i := 0; i < 10; i++ {
		list2.InsertLast(i)
	}
	if list2.Size() != 10 {
		t.Error("DLLinkedList InsertLast failed")
	} else if list2.String() != "0->1->2->3->4->5->6->7->8->9->nil" {
		t.Error("DLLinkedList String failed")
	} else if list2.InvString() != "nil<-0<-1<-2<-3<-4<-5<-6<-7<-8<-9" {
		t.Error("DLLinkedList InvString failed")
	} else {
		t.Log("DLLinkedList succeed")
	}
}

func TestSingleRing(t *testing.T) {
	data := []interface{}{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	ring1 := LinkedList.NewSingleRingFromSlice(data...)
	if ring1.Size() != 10 {
		t.Error("New SingleRing failed")
	} else if ring1.String() != "0->1->2->3->4->5->6->7->8->9->nil" {
		t.Error("SingleRing String failed")
	} else {
		t.Log("SingleRing succeed")
	}
}
