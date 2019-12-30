package main

import "fmt"

type intptr struct {
	iptr *int64
	s    string
}

func testIntPtr() {
	var b int64 = 5
	a := intptr{
		iptr: &b,
		s:    "000",
	} // 只能这样初始化
	fmt.Println(a)
}
