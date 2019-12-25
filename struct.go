package main

import "fmt"

type bbb struct {
	f float64
	s string
}

type aaa struct {
	int64
	*bbb
}

func testStruct() {
	a := aaa{
		int64: 3,
		bbb:   &bbb{
			f: 3.13,
			s: "123",
		},

	}
	fmt.Println(a.int64, a.s, a.f)
	a.int64 = 34545
	a.f = 3.14
	a.s = "4444"
	fmt.Println(a.int64, a.s, a.f)
}
