package main

import (
	"fmt"
	"reflect"
	"time"
)

func TestTimeFormat() { // 必须是这个时间
	t := time.Now().Format("2006-01-02 15:04:05")
	fmt.Println(t, reflect.TypeOf(t))
}
