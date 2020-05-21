package LinkedList

import (
	Node "algorithms/linked_list/Node/src"
)

type SingleRing struct {
	head *Node.Node
	size int
}

func NewSingleRing() *SingleRing {
	return &SingleRing{nil, 0}
}

//NewSingleRingFromSlice 从slice构造单链表，可传入单个元素、数组或者其组合
//eg. NewSingleRingFromSlice(1,2,3)
//eg. NewSingleRingFromSlice(1,[]int{2,3},[]string{"hello","world","this","is","adam"})
//eg. NewSingleRingFromSlice([]interface{}{"hello","world","this","is","adam"}...)
//eg. NewSingleRingFromSlice([]interface{1,2,3}...)
func NewSingleRingFromSlice(slice ...interface{}) *SingleRing {
	length := len(slice)
	ring := NewSingleRing()
	if length == 0 {
		return ring
	} else {
		ring.InsertLast(slice[0])
		if length == 1 {
			return ring
		} else {
			p := ring.head
			for i := 1; i < length; i++ {
				p.Next = &Node.Node{slice[i], ring.head}
				p = p.Next
			}
			ring.size += (length - 1)
		}
	}
	return ring
}

func (ring SingleRing) First() Node.Node {
	return *ring.head
}

func (ring SingleRing) Size() int {
	return ring.size
}

func (ring *SingleRing) InsertLast(v interface{}) {
	p := ring.head
	if p.IsNil() {
		ring.head = &Node.Node{v, nil}
		ring.head.Next = ring.head
	} else {
		for ; !p.Next.IsNil(); p = p.Next {
		}
		p.Next = &Node.Node{v, ring.head}
	}
	ring.size++
}

//String() 打印为字符串
func (ring SingleRing) String() string {
	p := ring.head
	for i := 1; i < ring.size; i++ {
		p = p.Next
	}
	p.Next = nil
	return ring.head.String()
}
