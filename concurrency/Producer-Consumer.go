package main

import (
	"fmt"
	"time"
)

func Producer(factor int, out chan<- int) {
	for i := 0; ; i++ {
		out <- i * factor
	}
}

func Consumer(in <-chan int) {
	for v := range in {
		fmt.Println(v)
	}
}

func main() {
	ch := make(chan int, 64) //成果队列

	go Producer(3, ch) //生成3的倍数的序列
	go Producer(5, ch) //生成3的倍数的序列
	go Consumer(ch)    //消费生成的序列

	time.Sleep(5 * time.Second)
}
