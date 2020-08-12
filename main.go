package main

func main() {
	/* TestStruct */
	//testStruct()

	/* TestPointer */
	//a := A{
	//	int64: 2,
	//}
	//NoPointer(a)
	//fmt.Println(a)
	//Pointer(&a)
	//fmt.Println(a)

	/* TestOOP */
	//bas := &Base{
	//	F: 4.2,
	//	S: "123",
	//}
	//der := &Derived{
	//	I:    64,
	//	Base: bas,
	//}
	//TestOOP1(bas)
	//TestOOP1(der)
	//TestOOP2(bas)
	//// TestOOP2(der) // type no match
	//// TestOOP3(bas) // type no match
	//TestOOP3(der)
	//TestOOP4(bas)
	//TestOOP4(der)

	/* TestMap */
	//a, b := RetVarDefinedNil(2)
	//fmt.Println(a == nil, b == nil)
	//a, b = RetVarDefined(2)
	//fmt.Println(a, b) //slice不用make, map要make
	//definedMapSpace(3) //map会自动扩展容量
	//MapJsonConvert() // key是interface类型的map是不能marshal和unmarshal的
	//MapDelete()
	//MapSort()

	/* TestWaitGroup */
	//var wg sync.WaitGroup
	//workerCount := 2
	//for i := 0; i < workerCount; i++ {
	//	wg.Add(1) //只能是1
	//	go testWaitGroup(i, &wg)
	//}
	//time.Sleep(1 * time.Second)
	//wg.Wait()
	//fmt.Println("all done!")

	/* TestTime */
	//TestTimeFormat()

	/* TestSlice */
	//TestCutSlice()
	//SliceCopy()
	//SliceStruct()

	/* TestClosure */
	//t := Fibonacci()
	//for i := 0; i < 10; i++ {
	//	fmt.Println(t())
	//}

	/* TestRand */
	//TestRandInt(270)

	/* TestString */
	//StrToByteArrAndReverse()
	//StrToIntAndReverse()
	//StrToInt32AndReverse()
	//StrToInt64AndReverse()
	//StrToFloat32AndReverse() // 精度就6位小数
	//StrToFloat64AndReverse()

	/* TestMutex */
	//TestMutex()

	/* TestJson */
	TestJson()
}
