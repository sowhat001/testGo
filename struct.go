package main

import "fmt"

type A struct {
	int64
}

type B struct {
	F float64
	S string
	A
}

type C struct {
	bool
	*B
}

func TestStruct() {
	c := C{
		bool: false,
		B: &B{
			F: 1.64,
			S: "testStruct",
			A: A{
				int64: 12345,
			},
		},
	}
	fmt.Println(c.bool, c.B, c.F, c.S, c.A, c.int64)
	c.A = A{int64: 54321}
	fmt.Println(c.bool, c.B, c.F, c.S, c.A, c.int64)
}

func NoPointer(a A) {
	a.int64 = 123
}

func Pointer(a *A) {
	a.int64 = 123
}
