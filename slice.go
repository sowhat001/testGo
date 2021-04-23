package main

import (
	"fmt"
)

// 不保证同步性，需要加锁操作

func TestSliceCopy() {
	// 浅拷贝
	a := []int{1, 2, 3, 4, 5, 6}
	fmt.Printf("a 地址: %p\n", a)
	b := a[0:3]
	fmt.Println("a is: ", a)
	fmt.Println("b is: ", b)
	fmt.Printf("a 地址 :%p\n", a)
	fmt.Printf("b 地址 :%p\n", b) // a, b 地址相同
	b[0] = 100
	fmt.Println("a is: ", a)
	fmt.Println("b is: ", b) // a[0], b[0] 都被修改
	c := make([]int, 3)
	c = a[0:3]
	fmt.Println("c is: ", c)    // c[0] == 100
	fmt.Printf("c 地址: %p\n", c) // a, b, c 地址都相同

	s1 := []int{1, 2, 3, 4, 5, 6}
	s2 := s1
	fmt.Printf("s1 地址 %p,s2 地址 %p\n", s1, s2) // s1, s2 地址相同
	s1[1] = 100
	s2[0] = 200
	fmt.Printf("s1 地址 %p,s2 地址 %p\n", s1, s2) // s1, s2 地址相同
	fmt.Println("s1 is :", s1)
	fmt.Println("s2 is :", s2) // s1, s2 值完全相同

	// 当元素数量超过容量
	// 切片会在底层申请新的数组
	// 从而达成两个切片地址不同，深拷贝的功能
	// 但是还是要用copy做深拷贝
	s1 = append(s1, 666)
	s2[2] = 300
	fmt.Printf("s1 地址 %p, s2 地址 %p\n", s1, s2) // s1, s2 地址不同
	fmt.Println("s1 is: ", s1)                 //
	fmt.Println("s2 is: ", s2)                 // s1[2] != 300，与 s1 值不同
	// copy 函数提供深拷贝功能
	// 但需要在拷贝前申请空间
	// copy 返回值是 copy 的元素数量，是两个 slice 中 size 最小的那个 size
	s3 := make([]int, 4)
	s4 := make([]int, 5)
	fmt.Println(copy(s3, s1)) // 输出 4
	fmt.Println(copy(s4, s1)) // 输出 5
	s3[0] = 300
	s4[0] = 400
	fmt.Println("s1 is:", s1)
	fmt.Println("s3 is:", s3)
	fmt.Println("s4 is:", s4) // s1, s3, s4 值各不相同
}

func TestSliceAppend() {
	a := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(a)                                            // [1 2 3 4 5 6]
	a = append(a[:3], append([]int{1}, a[3:]...)...)          // 在第 4 个位置插入 1
	fmt.Println(a)                                            // [1 2 3 1 4 5 6]
	a = append(a[:3], append([]int{10, 20, 30}, a[3:]...)...) // 在第 4 个位置插入切片
	fmt.Println(a)                                            // [1 2 3 10 20 30 1 4 5 6]
	a = append(a, 199)                                        // 追加 1 个元素
	fmt.Println(a)                                            // [1 2 3 10 20 30 1 4 5 6 199]
	a = append(a, 11, 21, 31)                                 // 追加多个元素, 手写解包方式
	fmt.Println(a)                                            // [1 2 3 10 20 30 1 4 5 6 199 11 21 31]
	a = append(a, []int{111, 222, 333}...)                    // 追加一个切片, 切片需要解包
	fmt.Println(a)                                            // [1 2 3 10 20 30 1 4 5 6 199 11 21 31 111 222 333]
	a = append([]int{0}, a...)                                // 在开头添加 1 个元素
	fmt.Println(a)                                            // [0 1 2 3 10 20 30 1 4 5 6 199 11 21 31 111 222 333]
	a = append([]int{-3, -2, -1}, a...)                       // 在开头添加 1 个切片
	fmt.Println(a)                                            // [-3 -2 -1 0 1 2 3 10 20 30 1 4 5 6 199 11 21 31 111 222 333]
	// 在开头添加会导致整个 slice 被重新 copy 一遍，效率低

	slist := make([]int, 0, 10)
	s3 := append(slist, 1)
	s4 := append(slist, 2)
	fmt.Println(s3) //[2]
	fmt.Println(s4) //[2]
}
