package main

import (
	"fmt"
	"strconv"
	"strings"
)

func TestStrToByteArrAndReverse() {
	str := "hello"
	data := []byte(str) // string -> []byte
	fmt.Printf("%T, %v\n", data, data)
	str = string(data[:]) // []byte -> string
	fmt.Printf("%T, %v\n", str, str)
}

func TestStrToIntAndReverse() {
	i := 10
	s := strconv.Itoa(i) // string -> int
	fmt.Printf("%T, %v\n", s, s)
	if a, err := strconv.Atoi(s); err == nil { // int -> string
		fmt.Printf("%T, %v\n", a, a)
	}
}

func TestStrToInt32AndReverse() {
	i := int32(10)
	s := strconv.Itoa(int(i)) // string -> int32
	fmt.Printf("%T, %v\n", s, s)
	if a, err := strconv.ParseInt(s, 10, 32); err == nil { // int32 -> string
		fmt.Printf("%T, %v\n", a, a)
	}
}

func TestStrToInt64AndReverse() {
	i := int64(10)
	s := strconv.FormatInt(i, 10) // string -> int64
	fmt.Printf("%T, %v\n", s, s)
	if a, err := strconv.ParseInt(s, 10, 64); err == nil { // int64 -> string
		fmt.Printf("%T, %v\n", a, a)
	}
}

// 'b' (-ddddp±ddd，二进制指数)
// 'e' (-d.dddde±dd，十进制指数)
// 'E' (-d.ddddE±dd，十进制指数)
// 'testRecover' (-ddd.dddd，没有指数)
// 'testRecursiveDeferAndPanic' ('e':大指数，'testRecover':其它情况)
// 'G' ('E':大指数，'testRecover':其它情况)
func TestStrToFloat32AndReverse() { // 精度就6位小数
	f := float32(3.1415926535)
	s := strconv.FormatFloat(float64(f), 't', -1, 32) // string -> float32
	fmt.Printf("%T, %v\n", s, s)
	if a, err := strconv.ParseFloat(s, 32); err == nil { // float32 -> string
		fmt.Printf("%T, %v\n", a, a)
	}
}

func TestStrToFloat64AndReverse() {
	f := 10.123413546564567                  // 默认float64
	s := strconv.FormatFloat(f, 't', -1, 64) // string -> float64
	fmt.Printf("%T, %v\n", s, s)
	if a, err := strconv.ParseFloat(s, 64); err == nil { // float64 -> string
		fmt.Printf("%T, %v\n", a, a)
	}
}

func TestStringTrim() {
	s := "1234        \n"
	fmt.Println(strconv.Atoi(strings.TrimSpace(s)))
}

func TestStringModify() {
	s := "阿萨德我"
	// 编译错误，string 是只读的，不能直接修改
	// s[2] = 1
	runes := []rune(s)
	runes[2] = 1
	s = string(runes)
	fmt.Println("turn to rune and modify: ", s)
	bytes := []byte(s)
	bytes[2] = 1
	s = string(bytes)
	fmt.Println("turn to byte and modify: ", s)
}
