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

func TestTimeDuration() { // duration可以直接转化成int相加，除非叠加时间超过58年
	t := time.Now()
	time.Sleep(time.Second * 5)
	end := time.Since(t)
	fmt.Println(int64(end))
	fmt.Println(9223372036854775807)
}
