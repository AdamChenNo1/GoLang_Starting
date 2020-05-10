package main

import (
	"time"

	"github.com/samuel/go-zookeeper/zk"
)

func main() {
	c, _, err := zk.Connect([]string{"35.221.200.251"}, time.Second)
	if err != nil {
		panic(err)
	}
	// children, stat, ch, err := c.ChildrenW("/")
	// if err != nil {
	// 	panic(err)
	// }
	l := zk.NewLock(c, "/lock", zk.WorldACL(zk.PermAll))
	err = l.Lock()
	if err != nil {
		panic(err)
	}
	println("lock succ,do your bussiness logic")

	time.Sleep(time.Second * 10)

	l.Unlock()
	println("unlock succ,finish bussiness logic")
}
