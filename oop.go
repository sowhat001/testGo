package main

import (
	"fmt"
)

type testInterface interface {
	bark()
	print()
}

type base struct {
	f float64
	s string
}

func (b *base) M1(val int64) {
	fmt.Println("baseM1 ", val-10086, b.f, b.s)
}

func (b *base) M2(val int64) {
	fmt.Println("baseM2 ", val+10086, b.f, b.s)
}

func (b *base) bark() {
	fmt.Println("bark ", b.f, b.s)
}

func (b *base) print() {
	fmt.Println("print ", b.f, b.s)
}

type derived struct {
	int64
	*base
}

func (d *derived) M1(val int64) {
	fmt.Println("derivedM1 ", val-10086, d.int64, d.f, d.s)
}

func (d *derived) M2(val int64) {
	fmt.Println("derivedM2 ", val+10086, d.int64, d.f, d.s)
}

func (d *derived) bark() {
	fmt.Println("bark ", d.int64, d.f, d.s)
}

func (d *derived) print() {
	fmt.Println("print ", d.int64, d.f, d.s)
}

func testOOP1(in testInterface) {
	in.bark()
	in.print()
}

func testOOP2(bas *base) {
	bas.M1(55)
	bas.M2(55)
}

func testOOP3(der *derived) {
	der.M1(5)
	der.M2(5)
}

func testOOP4(val interface{}) {
	switch v := val.(type) {
	case *derived:
		v.M1(5)
		v.M2(5)
	case *base:
		v.M1(5)
		v.M2(5)
	default:
		fmt.Println("unknown")
	}
}
