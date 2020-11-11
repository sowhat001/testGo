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

func TestStructWithAnonymousField() {
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
	fmt.Println(c.bool, c.B, c.F, c.S, c.A, c.int64) // 没有变量名的直接用类型名代替
	c.A = A{int64: 54321}
	fmt.Println(c.bool, c.B, c.F, c.S, c.A, c.int64) // 修改成功
}

func TestModifyFieldWithPointerOrNot() {
	a1 := A{int64(2)}
	a2 := A{int64(2)}
	modifyFieldWithPointerOrNot(a1, &a2)
	fmt.Println("ModifyFieldWithoutPointer: ", a1)
	fmt.Println("ModifyFieldWithPointer: ", a2)
}

func modifyFieldWithPointerOrNot(a A, ap *A) {
	a.int64 += 1
	ap.int64 += 1
}

func (a *A) add(i int64) {
	a.int64 += i
}

func addFunc(a *A, i int64) {
	a.int64 += i
}

func TestCallMethodAndFuncWithPointerOrNot() {
	a := A{int64(2)}
	// 编译错误
	// addFunc(a, 2)
	// 下面都编译成功
	addFunc(&a, 2)
	fmt.Println(a)
	a.add(2)
	fmt.Println(a)
	(&a).add(2)
	fmt.Println(a)
}
