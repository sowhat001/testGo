package main

import (
	"fmt"
	"sync"
)

func testMutex() {
	var l *sync.Mutex
	l = new(sync.Mutex)
	l.Lock()
	go func() {
		l.Lock()
		l.Unlock()
		fmt.Println("unlock")
	}()
	fmt.Println("1")
}