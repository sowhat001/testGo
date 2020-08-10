package main

import (
	"fmt"
	"math/rand"
	"time"
)

func TestRandInt(num int) {
	fmt.Println("Default Seed, 每次运行都是一样的")
	for i := 0; i < 100; i++ {
		fmt.Print(rand.Intn(num))
		fmt.Print("\t")
	}
	fmt.Println()
	fmt.Println("时间戳seed，每次运行不一样")
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 100; i++ {
		fmt.Print(rand.Intn(num))
		fmt.Print("\t")
	}
}
