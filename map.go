package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"sort"
)

func RetVarDefinedNil(new int) (aaa []int, match map[int]int) {
	_ = new
	return
}

func RetVarDefined(new int) (aaa []int, match map[int]int) {
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

func MapJsonConvert() {
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

func MapDelete() { //遍历删除是安全的
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

func MapSort() {
	x := map[int]int{}
	for i := 0; i < 5; i++ {
		x[i] = i%3 + rand.Intn(2123)/45
	}
	fmt.Println("初始化，map的值为", x)
	fmt.Println("遍历如下，不一定有序")
	for k, v := range x {
		fmt.Printf("%d, %d\n", k, v)
	}

	fmt.Println("按key排序")
	keys := make([]int, len(x))
	i := 0
	for k, _ := range x {
		keys[i] = k
		i++
	}
	sort.SliceStable(keys, func(i, j int) bool {
		return keys[i] < keys[j]
	})
	for _, v := range keys {
		fmt.Printf("%d, %d\n", v, x[v])
	}
}

func NilMap(m map[int64]int64) {
	print(m[2])
}
