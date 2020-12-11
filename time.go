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

func TestLoopTime() { //unixnano还是最好不当uuid
	for i := 0; i < 20; i++ {
		go func() {
			a := time.Now().UnixNano()
			fmt.Println(a)
		}()
	}
	for {
		_ = 1
	}
}

func TestDateValidator() {
	var err error
	date := "2019-02-29"
	// match, err := regexp.MatchString("[0-9]{4}-[0-9]{2}-[0-9]{2}", date)
	// fmt.Println(match, err)
	_, err = time.Parse("2006-01-02", date)
	fmt.Println(err)
}
