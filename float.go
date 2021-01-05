package main

import "fmt"

func TestZero() {
	a := float32(0.00)
	b := float32(0.000000)
	fmt.Println(a > 0, b > 0)
}

func TestDivide() {
	a := float32(2.3)
	b := int(1)
	fmt.Println(a / float32(b))
}
