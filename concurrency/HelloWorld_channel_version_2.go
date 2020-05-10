package main

import "fmt"

func main() {
	done := make(chan int, 1) // 带缓存的管道

	go func() {
		fmt.Println("你好, 世界")
		done <- 1
	}()

	<-done
}