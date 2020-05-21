package LinkedList

import (
	Node "algorithms/linked_list/Node/src"
)

type DLLinkedList struct {
	head   *Node.DLNode
	tailer *Node.DLNode
	size   int
}

func NewDLLinkedList() *DLLinkedList {
	return &DLLinkedList{nil, nil, 0}
}

func (list DLLinkedList) First() Node.DLNode {
	return *list.head
}

func (list DLLinkedList) Last() Node.DLNode {
	return *list.tailer
}

func (list DLLinkedList) Size() int {
	return list.size
}

func (list *DLLinkedList) InsertLast(v interface{}) {
	if list.size == 0 {
		list.tailer = &Node.DLNode{v, nil, nil}
		list.head = list.tailer
	} else {
		list.tailer.Next = &Node.DLNode{v, nil, list.tailer}
		list.tailer = list.tailer.Next
	}
	list.size++
}

//String 打印为字符串
func (list DLLinkedList) String() string {
	return list.head.String()
}

//InvString 打印为逆向字符串
func (list DLLinkedList) InvString() string {
	return list.head.InvString()
}

//DeleteReverseK 删除导数第K个节点
func (list *DLLinkedList) DeleteReverseK(k int) {
	length := list.size
	if k < 0 || k > length {
		panic("k is out of bound")
	}
	if k > length/2 {
		p := list.head
		if k == length {
			list.head = p.Next
			list.head.Prev = nil
			p.Next = nil
			list.size--
			return
		}
		for i := 1; i < length-k; i++ {
			p = p.Next
		}
		trash := p.Next
		p.Next = trash.Next
		trash.Next.Prev = p
		trash.Next = nil
		trash.Prev = nil
	} else {
		p := list.tailer
		if k == 1 {
			list.tailer = p.Prev
			list.tailer.Next = nil
			p.Prev = nil
			list.size--
			return
		}
		for i := 2; i < k; i++ {
			p = p.Prev
		}
		trash := p.Prev
		p.Prev = trash.Prev
		trash.Prev.Next = p
		trash.Next = nil
		trash.Prev = nil
	}
	list.size--
}

//NewDLLinkedListFromSlice 从slice构造单链表，可传入单个元素、数组或者其组合
//eg. NewSDLinkedListFromSlice(1,2,3)
//eg. NewSDLinkedListFromSlice(1,[]int{2,3},[]string{"hello","world","this","is","adam"})
//eg. NewSDLinkedListFromSlice([]interface{}{"hello","world","this","is","adam"}...)
//eg. NewDLLinkedListFromSlice([]interface{1,2,3}...)
func NewDLLinkedListFromSlice(slice ...interface{}) *DLLinkedList {
	length := len(slice)
	list := NewDLLinkedList()
	if length == 0 {
		return list
	} else {
		list.InsertLast(slice[0])
		if length == 1 {
			return list
		} else {
			list.InsertLast(slice[1])
			if length == 2 {
				return list
			} else {
				t := list.tailer
				for i := 2; i < length; i++ {
					v := &Node.DLNode{slice[i], nil, t}
					t.Next = v
					t = v
				}
				list.tailer = t
				list.size += (length - 2)
			}
		}
	}
	return list
}

//DeleteElementAt 删除位于a/b处的元素，向下取整
func (list *DLLinkedList) DeleteElementAt(a, b int) {
	if a < 0 || a > b {
		panic("a and b must satisfy 0 <= a < b")
	}

	p := list.head
	t := list.tailer
	L := list.Size()
	if a == 0 {
		return
	}
	if b >= 2*a {
		if a*L <= b {
			list.head = p.Next
			list.head.Prev = nil
			p.Next = nil
			list.size--
			return
		}
		for i := 1; i <= L/2; i++ {
			x := b * i
			y := a * L
			if x < y && y <= x+b {
				v := p.Next
				p.Next = v.Next
				v.Next.Prev = p
				v.Next = nil
				v.Prev = nil
				list.size--
				return
			}
			p = p.Next
		}
	} else {
		if a <= b && (b-a)*L < b {
			list.tailer = t.Prev
			list.tailer.Next = nil
			t.Prev = nil
			list.size--
			return
		}
		for i := L - 1; i > L/2; i-- {
			x := b * i
			y := a * L
			if x < y && y <= x+b {
				t.Prev.Next = t.Next
				t.Next.Prev = t.Prev
				t.Next = nil
				t.Prev = nil
				list.size--
				return
			}
			t = t.Prev
		}
	}

}

func (list *DLLinkedList) Reverse() {
	length := list.Size()
	if length == 0 || length == 1 {
		return
	}
	p := list.head
	v := p.Next
	p.Next = nil
	for ; v != nil; v = v.Prev {
		p.Prev = v
		v.Prev = v.Next
		v.Next = p
		p = v
	}
	list.head, list.tailer = list.tailer, list.head
}

//ReversePart 逆转双链表的一部分
func (list *DLLinkedList) ReversePart(from, to int) {
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
		if to < length-1 {
			after.Prev = nil
		}
		Tailer := list.tailer
		list.tailer = end
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
			after.Prev = p
			list.size = length
			list.tailer = Tailer
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
	Tailer := list.tailer
	list.tailer = end
	list.head.Prev = nil
	list.tailer.Next = nil
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
	after.Prev = p
	before.Next = list.head
	list.head.Prev = before
	list.head = Head
	list.tailer = Tailer
	list.size = length
}
