package main

import (
	"fmt"
)

type TestInterface interface {
	Bark()
	Print()
}

type Base struct {
	F float64
	S string
}

func (b *Base) M1(val int64) {
	fmt.Println("baseM1 ", val-10086, b.F, b.S)
}

func (b *Base) M2(val int64) {
	fmt.Println("baseM2 ", val+10086, b.F, b.S)
}

func (b *Base) Bark() {
	fmt.Println("Base bark ", b.F, b.S)
}

func (b *Base) Print() {
	fmt.Println("Base print ", b.F, b.S)
}

type Derived struct {
	I int64
	*Base
}

func (d *Derived) M1(val int64) {
	fmt.Println("derivedM1 ", val-10086, d.I, d.F, d.S)
}

func (d *Derived) M2(val int64) {
	fmt.Println("derivedM2 ", val+10086, d.I, d.F, d.S)
}

func (d *Derived) Bark() {
	fmt.Println("Derived bark ", d.I, d.F, d.S)
}

func (d *Derived) Print() {
	fmt.Println("Derived print ", d.I, d.F, d.S)
}

func TestOOP1(in TestInterface) {
	in.Bark()
	in.Print()
}

func TestOOP2(bas *Base) {
	bas.M1(55)
	bas.M2(55)
}

func TestOOP3(der *Derived) {
	der.M1(5)
	der.M2(5)
}

func TestOOP4(val interface{}) {
	switch v := val.(type) {
	case *Derived:
		v.M1(5)
		v.M2(5)
	case *Base:
		v.M1(5)
		v.M2(5)
	default:
		fmt.Println("unknown")
	}
}
