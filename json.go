package main

import (
	"encoding/json"
	"fmt"
)

func TestJson() {
	m := make(map[int]*A)
	m[23] = &A{int64(123)}
	m[34] = &A{int64(4535)}
	res, err := json.Marshal(m)
	if err == nil {
		fmt.Println(string(res))
	}
	res, _ = json.Marshal(nil)
	fmt.Println(string(res))
}
