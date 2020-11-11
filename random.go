package main

import (
	"fmt"
	"math/rand"
	"time"
)

func TestRandomInt() {
	num := 270
	fmt.Println("Default Seed, 每次运行都是一样的")
	for i := 0; i < 100; i++ {
		fmt.Print(rand.Intn(num))
		fmt.Print("\t")
	}
	fmt.Println()
	fmt.Println("时间戳 seed，每次运行不一样")
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 100; i++ {
		fmt.Print(rand.Intn(num))
		fmt.Print("\t")
	}
}

func TestRandomNumWithSelect() {
	ch := make(chan int)
	go func() {
		for i := 0; i < 100; i++ {
			select {
			case ch <- 0:
			case ch <- 1:
			}
		}
	}()

	for v := range ch {
		fmt.Println(v)
	}
}
