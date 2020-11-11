package main

import "fmt"

// break 同样适用于 select 和 switch
// 所以 for 嵌套 select 或者 switch 的时候需要带 label 的 break

// 多变量循环写法
func TestMultiVariableLoop() {
	for i, j := 0, 10; i < j; i, j = i+1, j-1 {
		fmt.Printf("%d, %d\n", i, j)
	}
}

// 对 range 的形参取地址是不准的
func TestRangeAddress() {
	nums := []int{1, 2, 3}
	for i := range nums {
		fmt.Println(&i)
	}
	for i := 0; i < len(nums); i++ {
		fmt.Println(&nums[i])
	}
}
