package main

import (
	"fmt"
	"sync"
)

func main() {
	var mu sync.Mutex

	go func() {
		fmt.Println("你好, 世界")
		mu.Lock()
	}()

	mu.Unlock()
}
