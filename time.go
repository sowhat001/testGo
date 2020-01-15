package main

import (
	"fmt"
	"reflect"
	"time"
)

func testTimeFormat() {
	t := time.Now().Format("2006-01-02")
	fmt.Println(t, reflect.TypeOf(t))
}
