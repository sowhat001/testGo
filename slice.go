package main

import "fmt"

func testCut() { //error
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	b := a[3:7:len(a)]
	fmt.Println(b)
}
