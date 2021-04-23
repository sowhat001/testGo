package main

import (
	"fmt"
	"time"
)

// defer 的函数是在返回值计算之后，return 之前执行的
func TestReturnAfterDefer() {
	fmt.Println("ReturnAfterDefer return: ", testReturnAfterDefer(1)) // 输出 1
}

func testReturnAfterDefer(i int) int {
	defer func(i int) {
		fmt.Println("ReturnAfterDefer before defer: ", i) // 输出 1
		i++
		fmt.Println("ReturnAfterDefer defer finish: ", i) // 输出 2
	}(i)
	return i
}

// return 之后的 defer 不会被执行
func TestDeferAfterReturn() {
	fmt.Println("before defer")
	return
	defer fmt.Println("defer1")
	defer func() {
		fmt.Println("defer2")
	}()
}

// 先输出 1，再输出 0，defer 是压栈
func TestManyDefers() {
	i := 0
	defer fmt.Println(i)
	i++
	defer fmt.Println(i)
	return
}

// 因为 defer 的函数是在返回值计算之后，return 之前执行的，所以可以改变返回值，输出 2
func TestModifyReturnValue() {
	fmt.Println(testModifyReturnValue())
}

func testModifyReturnValue() (i int) {
	defer func() {
		i++
	}()
	return 1
}

// 虽然 time.Since 也是一个函数，但是有可能被内联了，导致值被直接计算出来了
func TestDeferVariable() { // 睡了2s 打印出来一个100ns
	t1 := time.Now()
	defer fmt.Println(time.Since(t1))
	time.Sleep(2 * time.Second)
}

func TestDeferFunc() { // 睡了2s 打印出来一个2s
	t1 := time.Now()
	defer func(t1 time.Time) {
		fmt.Println(time.Since(t1))
	}(t1)
	time.Sleep(2 * time.Second)
}

// 输出结果：
// Calling testRecursiveDeferAndPanic.
// Printing in testRecursiveDeferAndPanic 0
// Printing in testRecursiveDeferAndPanic 1
// Printing in testRecursiveDeferAndPanic 2
// Printing in testRecursiveDeferAndPanic 3
// Big Panic!!!
// Defer in testRecursiveDeferAndPanic 3
// Defer in testRecursiveDeferAndPanic 2
// Defer in testRecursiveDeferAndPanic 1
// Defer in testRecursiveDeferAndPanic 0
// Recovered in testRecover asdasdasdasd, 4
// TestPanicAndRecover Over

// 从 testRecursiveDeferAndPanic 函数中退出，先递归的将之前的 defer 执行，然后执行 testRecover 中的 defer，被 recover 捕捉到 panic 的内容
func TestPanicAndRecover() {
	testRecover()
	fmt.Println("TestPanicAndRecover Over")
}

func testRecover() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in testRecover", r)
		}
	}()
	fmt.Println("Calling testRecursiveDeferAndPanic.")
	testRecursiveDeferAndPanic(0)
	fmt.Println("Returned normally from testRecursiveDeferAndPanic.") // 调用不到，因为抛出了 panic，直接调到 defer 里面的 recover
}

func testRecursiveDeferAndPanic(i int) {
	if i > 3 {
		fmt.Println("Big Panic!!!")
		panic(fmt.Sprintf("asdasdasdasd, %v", i))
	}
	defer fmt.Println("Defer in testRecursiveDeferAndPanic", i)
	fmt.Println("Printing in testRecursiveDeferAndPanic", i)
	testRecursiveDeferAndPanic(i + 1)
}

func shouldNotExit() {
	for {
		time.Sleep(1 * time.Second)
		if time.Now().UnixNano()&0x03 == 0 {
			panic("The server is stolen")
		}
	}
}

func neverExit(name string, f func()) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("%v, %v\n", name, r)
			go neverExit(name, f)
		}
	}()
	f()
}

// 一直输出 The server is stolen
func TestRestartGoRoutine() {
	go neverExit("a", shouldNotExit)
	go neverExit("b", shouldNotExit)
	go neverExit("c", shouldNotExit)
	select {}
}

func TestDeferAfterPanic() {
	a := make([]int, 0)
	a = nil
	fmt.Println(a[0])

	// b := make(map[string]int)
	defer func() { // 没用
		if r := recover(); r != nil {
			fmt.Printf("recover, %v", r)
		}
	}()
}
