package LinkedList

import (
	"fmt"
	"reflect"
)

type Node struct {
	Value interface{}
	Next *Node
}

func (head *Node) String() string{
	if head.IsNil() {
		return "nil"
	}else {
		return fmt.Sprint(head.Value) + "->" + head.Next.String()
	}
}

func (node *Node) IsNil() bool {
	return reflect.ValueOf(node).IsNil()
} 

