package main

import (
	"fmt"
	"strconv"
	"strings"
)

func TestStrToByteArrAndReverse() {
	str := "阿萨德阿萨德阿萨德"
	data := []byte(str) // string -> []byte
	fmt.Printf("%T, %v\n", data, data)
	str = string(data) // []byte -> string
	fmt.Printf("%T, %v\n", str, str)
}

func TestStrToIntAndReverse() {
	i := 10
	s := strconv.Itoa(i) // int -> string
	fmt.Printf("%T, %v\n", s, s)
	if a, err := strconv.Atoi(s); err == nil { // string -> int
		fmt.Printf("%T, %v\n", a, a)
	} else {
		fmt.Println(err)
	}
}

func TestStrToInt32AndReverse() {
	i := int32(323232323)     // 这个数超出 int32 就会报错
	s := strconv.Itoa(int(i)) // // int32 -> string
	fmt.Printf("%T, %v\n", s, s)
	s = "2222222222222222"                                 // 超出 int32 的数字
	if a, err := strconv.ParseInt(s, 10, 32); err == nil { // string -> int32
		fmt.Printf("%T, %v\n", a, a) // type 还是 int64，需要手动转
	} else {
		fmt.Println(err) // 超出int32会报错
	}
}

func TestStrToInt64AndReverse() {
	i := int64(646464646464646464)
	s := strconv.FormatInt(i, 10) // int64 -> string
	fmt.Printf("%T, %v\n", s, s)
	if a, err := strconv.ParseInt(s, 10, 64); err == nil { // string -> int64
		fmt.Printf("%T, %v\n", a, a)
	} else {
		fmt.Println(err)
	}
}

// 'b' (-ddddp±ddd，二进制指数)
// 'e' (-d.dddde±dd，十进制指数)
// 'E' (-d.ddddE±dd，十进制指数)
// 'f' (-ddd.dddd，没有指数)
// 'g' (大指数情况下是'e'，否则是'f')
// 'G' (大指数情况下是'E'，否则是'f')
// prec: 'f'的情况下是小数点后的位数，指数的情况下是指数底的全部位数，-1代表不限制
func TestStrToFloat32AndReverse() {
	f := float32(3.1415926535)                        // 精度就7位小数
	s := strconv.FormatFloat(float64(f), 'f', -1, 32) // float32 -> string
	fmt.Printf("%T, %v\n", s, s)
	if a, err := strconv.ParseFloat(s, 32); err == nil { // string -> float32
		fmt.Printf("%T, %v\n", a, a) // type 还是 float64
		fmt.Println("float32: ", float32(a))
	} else {
		fmt.Println(err)
	}
}

func TestStrToFloat64AndReverse() {
	f := 10.123413546564567                  // 默认float64
	s := strconv.FormatFloat(f, 'f', -1, 64) // float64 -> string
	fmt.Printf("%T, %v\n", s, s)
	if a, err := strconv.ParseFloat(s, 64); err == nil { // string -> float64
		fmt.Printf("%T, %v\n", a, a)
	} else {
		fmt.Println(err)
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
