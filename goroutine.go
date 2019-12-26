package main

import (
	"fmt"
	"sync"
	"time"
)

func testWaitGroup(workerID int, wg *sync.WaitGroup) {
	fmt.Printf("[%v] is running\n", workerID)
	defer wg.Done()
	time.Sleep(3 * time.Second) // 模拟 goroutine 正在执行
	fmt.Printf("[%v] is done\n", workerID)
}
