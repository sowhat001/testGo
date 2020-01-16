package main

import "fmt"

func stringToByteArrayAndReverse() {
	str2 := "hello"
	data2 := []byte(str2)
	fmt.Println(data2)
	str2 = string(data2[:])
	fmt.Println(str2)
}
