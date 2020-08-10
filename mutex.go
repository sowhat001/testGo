package main

import (
	"fmt"
	"sync"
	"time"
)

func TestMutex() {
	var l *sync.Mutex
	l = new(sync.Mutex)
	l.Lock()
	go func() {
		l.Unlock()
		fmt.Println("unlock")
	}()
	time.Sleep(1000)
}
