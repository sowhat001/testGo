package main

import (
	"fmt"
	"math/rand"
	"time"
)

func testRandInt(num int) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 100; i++ {
		fmt.Print(rand.Intn(num))
		fmt.Print("\t")
	}
}
