package LinkedList

import (
	"fmt"
	"reflect"
)

type DLNode struct {
	Value interface{}
	Next  *DLNode
	Prev  *DLNode
}

func (head *DLNode) String() string {
	if tp := reflect.ValueOf(head); tp.IsNil() {
		return "nil"
	} else {
		return fmt.Sprint(head.Value) + "->" + head.Next.String()
	}
}

func (head *DLNode) InvString() string {
	s := "nil"
	for cur := head; !reflect.ValueOf(cur).IsNil(); cur = cur.Next {
		s += "<-" + fmt.Sprint(cur.Value)
	}
	return s
}

func (node *DLNode) IsNil() bool {
	return reflect.ValueOf(node).IsNil()
}
