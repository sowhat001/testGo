package main

import (
	"encoding/json"
	"fmt"
)

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

func mapJson() { // key是interface类型的map是不能marshal和unmarshal的
	m := make(map[string][]int)
	m2 := make(map[interface{}]interface{})
	m3 := make(map[interface{}]int)
	m4 := make(map[string]interface{})
	m["123"] = []int{1, 4}
	m["21"] = []int{2, 4}
	m["1234"] = []int{566, 666}

	b, err := json.Marshal(m)
	temp := string(b)
	fmt.Println(temp, "err: ", err)

	err = json.Unmarshal([]byte(temp), &m2)
	fmt.Println(m2, err) //报错

	err = json.Unmarshal([]byte(temp), &m3)
	fmt.Println(m3, err) //报错

	err = json.Unmarshal([]byte(temp), &m4)
	fmt.Println(m4, err)

	b, err = json.Marshal(m2) // 报错
	temp = string(b)
	fmt.Println(temp, "err: ", err)

	b, err = json.Marshal(m3) // 报错
	temp = string(b)
	fmt.Println(temp, "err: ", err)

	b, err = json.Marshal(m4)
	temp = string(b)
	fmt.Println(temp, "err: ", err)

}

func mapDelete() { //遍历删除是安全的
	x := map[int]int{}
	for i := 0; i < 10000; i++ {
		x[i] = i
	}
	fmt.Println("初始化后,长度:", len(x))

	// 遍历时删除所有的偶数
	for k := range x {
		if k%2 == 0 {
			delete(x, k)
		}
	}
	fmt.Println("删除所有的偶数后,长度:", len(x))

	// 遍历时删除所有的元素
	for k := range x {
		delete(x, k)
	}
	fmt.Println("删除所有的元素后,长度:", len(x))
}
