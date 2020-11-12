package main

import (
	"fmt"
	"sync"
	"time"
)

// 可以在其它 goroutine 里面做 unlock
// lock 需要先 new
func TestMutexInAnotherGoroutine() {
	l := new(sync.Mutex)
	l.Lock()
	go func() {
		l.Unlock()
		fmt.Println("unlock")
	}()
	time.Sleep(1000)
}

// 不 lock 直接 unlock 会panic
func TestDirectUnlock() {
	l := new(sync.Mutex)
	l.Unlock()
	fmt.Println("No Panic")
}

// 死锁
func TestDoubleLock() {
	l := new(sync.Mutex)
	l.Lock()
	l.Lock()
	fmt.Println("No Deadlock")
}
