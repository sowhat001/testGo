package main

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
