package main

import "fmt"

type t1 struct {
	a int
	b float32
	c string
}

func TestCutSlice() { //error
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	b := a[3:7]
	fmt.Println(b)
}

func SliceCopy() {
	//浅拷贝
	a := []int{1, 2, 3, 4, 5, 6}
	fmt.Printf("a地址1:%p\n", a)
	b := a[0:3]
	fmt.Println("a is:", a)
	fmt.Println("b is:", b)
	fmt.Printf("a地址2:%p\n", a)
	fmt.Printf("b地址:%p\n", b)
	b[0] = 100
	fmt.Println("a is:", a)
	fmt.Println("b is:", b)
	c := make([]int, 3)
	c = a[0:3]
	fmt.Println("c is:", c)
	fmt.Printf("c地址:%p\n", c)
	s1 := []int{1, 2, 3, 4, 5, 6}
	s2 := s1
	fmt.Printf("s1地址%p,s2地址%p\n", s1, s2)
	s1[1] = 100
	s2[0] = 200
	fmt.Printf("s1地址%p,s2地址%p\n", s1, s2)
	fmt.Println("s1 is:", s1)
	fmt.Println("s2 is:", s2)

	s1 = append(s1, 666)
	// 当元素数量超过容量
	// 切片会在底层申请新的数组
	// 从而达成两个切片地址不同，深拷贝的功能
	// 但是还是要用copy做深拷贝
	s2[2] = 300
	fmt.Printf("s1地址%p,s2地址%p\n", s1, s2)
	fmt.Println("s1 is:", s1)
	fmt.Println("s2 is:", s2)
	// copy 函数提供深拷贝功能
	// 但需要在拷贝前申请空间
	s3 := make([]int, 4)
	s4 := make([]int, 5)
	fmt.Println(copy(s3, s1)) //4
	fmt.Println(copy(s4, s1)) //5
	s3[0] = 300
	s4[0] = 400
	fmt.Println("s1 is:", s1)
	fmt.Println("s3 is:", s3)
	fmt.Println("s4 is:", s4)
}

func SliceStruct() { //因为range得到的value是值的copy，对value取地址是不行的
	var src []t1
	src = append(src, t1{
		a: 2,
		b: 3.5,
		c: "123123123",
	})
	src = append(src, t1{
		a: 3,
		b: 4.5,
		c: "321321321",
	})
	cp := make([]*t1, 0, len(src))
	for i := 0; i < len(src); i++ { //正常
		cp = append(cp, &src[i])
		fmt.Printf("%p\n", &src[i])
		fmt.Println(cp)
		fmt.Println()
	}

	//for _, v := range src {  // 会全部都是最后一个v
	//	cp = append(cp, &v)
	//	fmt.Printf("%p\n", &v)
	//	fmt.Println(cp)
	//	fmt.Println()
	//}
	for _, v := range cp {
		fmt.Println(v.a)
		fmt.Println(v.b)
		fmt.Println(v.c)
		fmt.Println()
	}
}
