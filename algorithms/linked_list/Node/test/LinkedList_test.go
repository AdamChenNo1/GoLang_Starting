package LinkedList

import (
	dsa "algorithms/linked_list/Node/src"
	"testing"
)

func TestNode(t *testing.T) {
	list := &dsa.Node{0, nil}
	p := list
	for i := 1; i < 10; i++ {
		p.Next = &dsa.Node{i, nil}
		p = p.Next
	}
	if list.String() != "0->1->2->3->4->5->6->7->8->9->nil" {
		t.Error("Node String failed!")
	} else {
		t.Log("Node String passed")
	}
}

func TestDLNode(t *testing.T) {
	list := &dsa.DLNode{0, nil, nil}
	p := list
	for i := 1; i < 10; i++ {
		p.Next = &dsa.DLNode{i, nil, nil}
		p.Next.Prev = p
		p = p.Next
	}
	if list.String() != "0->1->2->3->4->5->6->7->8->9->nil" {
		t.Error("Node String failed!")
	} else if list.InvString() != "nil<-0<-1<-2<-3<-4<-5<-6<-7<-8<-9" {
		t.Error("Node String failed!")
	} else {
		t.Log("Node String passed")
	}
}
