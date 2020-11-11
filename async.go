package main

import (
	"fmt"
	"time"
)

func TestAsync() { // 调用async之后很快就返回，从channel里获取值的时候阻塞
	fmt.Println(time.Now())
	sum := asyncAdd()
	fmt.Println(time.Now())
	fmt.Println(<-sum)
}

func asyncAdd() chan int {
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

func TestPrimeFilter() {
	ch := generateNatural()
	for i := 0; i < 100; i++ {
		prime := <-ch // 新出现的素数
		fmt.Printf("%v: %v, %v\n", i+1, prime, ch)
		ch = primeFilter(ch, prime) // 基于新素数构造的过滤器
	}
}

func generateNatural() chan int {
	ch := make(chan int)
	go func() {
		for i := 2; ; i++ {
			ch <- i
		}
	}()
	return ch
}

// 管道过滤器: 删除能被素数整除的数
func primeFilter(in <-chan int, prime int) chan int {
	out := make(chan int)
	go func() {
		for {
			if i := <-in; i%prime != 0 {
				out <- i
			}
		}
	}()
	return out
}
