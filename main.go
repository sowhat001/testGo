package main

func main() {
	//bas := &base{
	//	f: 4.2,
	//	s: "123",
	//}
	//der := &derived{
	//	int64: 64,
	//	base:  bas,
	//}
	//t := test{
	//	a: 2,
	//}
	//testStruct()
	//testOOP1(bas)
	//// testOOP2(der) //cannot use der (type *derived) as type *base in argument to testOOP2
	//testOOP3(der)
	//testOOP4(bas)
	//testOOP4(der)
	//noPointer(t)
	//fmt.Println(t)
	//pointer(&t)
	//fmt.Println(t)
	//
	//definedMapSpace(3)
	//
	//a, b := retVarDefinedNull(2)
	//fmt.Println(a == nil, b == nil)
	//
	//a, b = retVarDefined(2)
	//fmt.Println(a, b) //slice不用make, map要make
	//
	//var wg sync.WaitGroup
	//workerCount := 2
	//for i := 0; i < workerCount; i++ {
	//	wg.Add(1) //只能是1
	//	go testWaitGroup(i, &wg)
	//}
	//time.Sleep(1 * time.Second)
	//wg.Wait()
	//fmt.Println("all done!")
	//testTimeFormat()
	// testCut()
	mapJson()
	//mapDelete()
}
