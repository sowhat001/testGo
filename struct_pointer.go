package main

import (
	"fmt"
	"reflect"
)

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

type D struct {
	a int
	b *string
	c string
}

type E struct {
	a *int
	b string
	c *string
}

type Scope struct {
	Type int
	IDs  []int64
	Ext  []string
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

func TestFieldEmptyWithPointer() {
	d := D{
		a: 0,
		c: "",
	}
	e := E{
		a: &d.a,
		b: *d.b,
		c: &d.c,
	}
	fmt.Println(d)
	fmt.Println(e)
}

func TestDeepEqual() {
	var (
		scopes1 = make(map[int]Scope)
		scopes2 = make(map[int]Scope)
		ext1    = make(map[string]interface{})
		ext2    = make(map[string]interface{})
	)
	scopes1[1] = Scope{
		Type: 1,
		IDs:  []int64{2, 3, 4},
		Ext:  []string{"2", "3", "4"},
	}
	scopes2[1] = Scope{
		Type: 1,
		IDs:  []int64{2, 3, 4},
		Ext:  []string{"2", "3", "4"},
	}
	ext1["12312321"] = scopes1
	ext2["12312321"] = scopes2
	fmt.Println(reflect.DeepEqual(scopes1, scopes2))
	fmt.Println(reflect.DeepEqual(ext1, ext2))
}
