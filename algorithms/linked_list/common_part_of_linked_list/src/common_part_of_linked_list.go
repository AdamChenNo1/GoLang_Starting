package CommonPartOfLL

import (
	dsa "algorithms/linked_list/Node/src"
	"fmt"
)

func PrintCommonPart(head1, head2 *dsa.Node) {
	cur1 := head1
	cur2 := head2
	for !cur1.IsNil() && !cur2.IsNil() {
		value1 := cur1.Value.(int)
		value2 := cur2.Value.(int)
		if value1 > value2 {
			cur2 = cur2.Next
		} else if value1 < value2 {
			cur1 = cur1.Next
		} else {
			fmt.Printf("%d ", value1)
			cur1 = cur1.Next
			cur2 = cur2.Next
		}
	}
}
