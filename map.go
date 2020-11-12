package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"sort"
	"time"
)

// map 是不能清空的

// map 要 make，slice 不用
func TestMapAndSliceMakeOrNot() {
	s1, m1 := retVarDefinedNil(2)
	fmt.Println("Do not make: ", s1 == nil, " ", m1 == nil)
	s2, m2 := retVarDefined(2)
	fmt.Println("Make: ", s2, " ", m2)
}

func retVarDefinedNil(new int) (aaa []int, match map[int]int) {
	_ = new
	return
}

func retVarDefined(new int) (aaa []int, match map[int]int) {
	aaa = append(aaa, new)
	match = make(map[int]int)
	match[new] = new
	return
}

// map 会自动扩展容量
func TestMapSpace() {
	space := 3
	m := make(map[string]int, space)
	for i := 0; i < space*2; i++ {
		m[fmt.Sprintf("%d", i)] = i
	}
	fmt.Println(m)
}

// interface{} 当 key 的 map，无法序列化/反序列化
func TestMapJsonConvert() {
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
	fmt.Println(m2, err) // 报错

	err = json.Unmarshal([]byte(temp), &m3)
	fmt.Println(m3, err) // 报错

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

	// map 中有结构体的时候，序列化会出错，结构体字段为空
	oldM := make(map[int]A)
	oldM[1] = A{int64(1)}
	oldM[2] = A{int64(2)}
	fmt.Printf("original m: %v\n", oldM)
	res, err := json.Marshal(oldM)
	fmt.Printf("m string value: %v, err: %v\n", string(res), err)
	newM := make(map[int]A)
	err = json.Unmarshal(res, &newM)
	fmt.Printf("m struct value: %v, err %v\n", newM, err)
	res, err = json.Marshal(nil)
	fmt.Printf("assign nil, value: %v, err: %v\n", string(res), err)
}

// 遍历删除是安全的
func TestMapDelete() {
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

func TestMapSort() {
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

func TestNilMap() {
	var m map[int64]int64
	// 编译不会错，输出 nil
	print(m[2])
	// 编译会错
	// m[2] = 1
}

// 并发读写 map 会直接报 fatal，不会 recover
// 输出 fatal error: concurrent map writes
func TestConcurrentMap() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("%v, panic recover!", r) // 不输出
		}
	}()
	m := make(map[int]int)
	go func() {
		for {
			m[1] = 1
		}

	}()
	go func() {
		for {
			m[2] = 2
		}
	}()
	time.Sleep(1 * time.Second)
}
