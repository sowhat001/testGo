package main

import (
	"fmt"
	"sync"
	"time"
)

func TestWaitGroup() {
	var wg sync.WaitGroup
	workerCount := 2
	for i := 0; i < workerCount; i++ {
		wg.Add(1) // 只能是1
		go testWaitGroup(i, &wg)
	}
	time.Sleep(1 * time.Second)
	wg.Wait()
	fmt.Println("all done!")
}

func testWaitGroup(workerID int, wg *sync.WaitGroup) {
	fmt.Printf("[%v] is running\n", workerID)
	defer wg.Done()
	time.Sleep(4 * time.Second) // 模拟 goroutine 正在执行
	fmt.Printf("[%v] is done\n", workerID)
}
