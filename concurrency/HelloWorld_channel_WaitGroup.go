package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	//开启N个后台打印线程
	for i := 0; i < 10; i++ {
		wg.Add(1)

		go func() {
			fmt.Println("你好，世界！")
			wg.Done()
		}()
	}

	//等待后台N个线程完成
	wg.Wait()
}
