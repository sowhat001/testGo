package main

import "fmt"

func retVarDefinedNull(new int) (aaa []int, match map[int]int) {
	_ = new
	return
}

func retVarDefined(new int) (aaa []int, match map[int]int) {
	aaa = append(aaa, new)
	match = make(map[int]int)
	match[new] = new
	return
}

func definedMapSpace(new int) {
	m := make(map[string]int, new)
	for i := 0; i < new*2; i++ {
		m[fmt.Sprintf("%d", i)] = i
	}
	fmt.Println(m)
}
