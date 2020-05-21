package LinkedList

import (
	Node "algorithms/linked_list/Node/src"
)

type SGLinkedList struct {
	head *Node.Node
	size int
}

func NewSGLinkedList() *SGLinkedList {
	return &SGLinkedList{nil, 0}
}

//NewSGLinkedListFromSlice 从slice构造单链表，可传入单个元素、数组或者其组合
//eg. NewSGLinkedListFromSlice(1,2,3)
//eg. NewSGLinkedListFromSlice(1,[]int{2,3},[]string{"hello","world","this","is","adam"})
//eg. NewSGLinkedListFromSlice([]interface{}{"hello","world","this","is","adam"}...)
//eg. NewSGLinkedListFromSlice([]interface{1,2,3}...)
func NewSGLinkedListFromSlice(slice ...interface{}) *SGLinkedList {
	length := len(slice)
	list := NewSGLinkedList()
	if length == 0 {
		return list
	} else {
		list.InsertLast(slice[0])
		if length == 1 {
			return list
		} else {
			p := list.head
			for i := 1; i < length; i++ {
				p.Next = &Node.Node{slice[i], nil}
				p = p.Next
			}
			list.size += (length - 1)
		}
	}
	return list
}

func (list SGLinkedList) First() Node.Node {
	return *list.head
}

func (list SGLinkedList) Size() int {
	return list.size
}

func (list *SGLinkedList) InsertLast(v interface{}) {
	p := list.head
	if p.IsNil() {
		list.head = &Node.Node{v, nil}
	} else {
		for ; !p.Next.IsNil(); p = p.Next {
		}
		p.Next = &Node.Node{v, nil}
	}
	list.size++
}

//String() 打印为字符串
func (list SGLinkedList) String() string {
	return list.head.String()
}

//DeleteReverseK 删除导数第K个节点
func (list *SGLinkedList) DeleteReverseK(k int) {
	length := list.size
	if k < 0 || k > length {
		panic("k is out of bound")
	}
	p := list.head
	if k == length {
		list.head = p.Next
		p.Next = nil
		list.size--
		return
	}
	for i := 1; i < length-k; i++ {
		p = p.Next
	}
	trash := p.Next
	p.Next = trash.Next
	trash.Next = nil
	list.size--
}

//DeleteElementAt 删除位于a/b处的元素，向上取整
func (list *SGLinkedList) DeleteElementAt(a, b int) {
	if a < 0 || a > b {
		panic("a and b must satisfy 0 <= a < b")
	}

	p := list.head
	L := list.Size()
	if a == 0 {
		return
	}
	if a*L <= b {
		list.head = p.Next
		p.Next = nil
		list.size--
		return
	}
	for i := 1; i < L; i++ {
		x := b * i
		y := a * L
		if x < y && y <= x+b {
			v := p.Next
			p.Next = v.Next
			v.Next = nil
			list.size--
			return
		}
		p = p.Next
	}
}

//Reverse 逆转单链表
func (list *SGLinkedList) Reverse() {
	length := list.size
	if length == 0 || length == 1 {
		return
	}
	p := list.head
	list.head = p.Next
	p.Next = nil
	list.size--
	list.Reverse()
	list.InsertLast(p.Value)
}

//ReversePart 逆转单链表的一部分
func (list *SGLinkedList) ReversePart(from, to int) {
	length := list.size
	if from > to {
		panic("start index must be smaller than the end")
	}

	if to <= 1 || from >= length || length <= 1 {
		return
	}
	if to > length-1 {
		to = length - 1
	}
	if from <= 0 {
		end := list.head
		for i := 0; i < to; i++ {
			end = end.Next
		}
		after := end.Next
		end.Next = nil
		list.size = to + 1
		list.Reverse()
		if to < length-1 {
			//连接2段
			p := list.head
			q := p.Next
			for !q.IsNil() {
				p = q
				q = q.Next
			}
			p.Next = after
			list.size = length
		}
		return
	}

	before := list.head
	for i := 1; i < from-1; i++ {
		before = before.Next
	}
	start := before.Next
	end := before
	for i := from - 1; i < to; i++ {
		end = end.Next
	}
	after := end.Next
	//分成三段
	before.Next = nil
	Head := list.head
	list.head = start
	end.Next = nil
	list.size = to - from + 1
	list.Reverse()
	//连接三段
	p := list.head
	q := p.Next
	for !q.IsNil() {
		p = q
		q = q.Next
	}
	p.Next = after
	before.Next = list.head
	list.head = Head
	list.size = length
}
