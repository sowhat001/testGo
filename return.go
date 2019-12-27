package main

import "fmt"

func noRetVal(v int64) {
	fmt.Println(v)
	return
	fmt.Println(v)
}

func retDefinedVal(v int64) (ret1 int64, ret2 int64) {
	ret1 = v
	ret2 = v + 2
	fmt.Println(ret1, ret2)
	return
}

//not enough arguments to return
//	have ()
//	want (int64, int64)
//func retUndefinedVal(v int64) (int64, int64) {
//	return
//}

// syntax error: mixed named and unnamed function parameters
//func retHalfDefinedVal (v int64) (ret1 int64, int64) {
//	return
//}
