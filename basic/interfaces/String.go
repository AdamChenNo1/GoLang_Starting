package interfaces

import (
	"strconv"
)


// 通过这个方法 Human 实现了 fmt.Stringer
func (h Human) String() string {
	return "❰"+h.Name+" - "+strconv.Itoa(h.Age)+" years -  ✆ " +h.Phone+"❱"
}
