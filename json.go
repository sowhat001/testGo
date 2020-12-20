package main

import (
	"encoding/json"
	"fmt"
)

func TestParseString() {
	m := make(map[int64]string)
	a := "{\"1\":\"\"}"
	// a := "{1:\"\",2:\"\", 14:\"\", 15:\"\"}"
	err := json.Unmarshal([]byte(a), &m)
	fmt.Println(m, err)
	fmt.Println(m[1])
	fmt.Println(len(m))
}
