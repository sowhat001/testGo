package main

import "fmt"

func TestFibonacci() {
	for i := 0; i < 10; i++ {
		fmt.Printf("第 %d 个斐波那契数列项是 %d", i+1, fibonacci())
	}
}

func fibonacci() func() int {
	sum1, sum2 := 0, 1
	return func() int {
		sum1, sum2 = sum2, sum1+sum2
		return sum1
	}
}
