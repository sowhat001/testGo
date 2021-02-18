package main

import (
	"fmt"
	"math/rand"
	"strconv"
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

func TestRandomString() {
	for i := 0; i < 300; i++ {
		s1 := strconv.FormatInt(int64(i*1557/257-i+i-i+i*4682/452>>2), 10)
		s2 := strconv.FormatInt(int64(i*19087/2157-i+i*i*2/167-i*4+i*4682/452>>2), 10)
		fmt.Printf("\"%v[{@}][{=}]%v\",\n", s1, s2)
	}
}
