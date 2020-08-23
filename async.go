package main

import (
	"fmt"
	"time"
)

func TestAsync() { // 调用async之后很快就返回，从channel里获取值的时候阻塞
	fmt.Println(time.Now())
	sum := AsyncAdd()
	fmt.Println(time.Now())
	fmt.Println(<-sum)
}

func AsyncAdd() chan int {
	result := make(chan int)
	go func() {
		add(result, 2, 3)
	}()
	return result
}

func add(result chan int, a, b int) {
	time.Sleep(3 * time.Second)
	result <- a + b
}
