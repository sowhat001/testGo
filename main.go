package main

func main() {
	bas := &base{
		f: 4.2,
		s: "123",
	}
	der := &derived{
		int64: 64,
		base:  bas,
	}
	testStruct()
	testOOP1(bas)
	// testOOP2(der) //cannot use der (type *derived) as type *base in argument to testOOP2
	testOOP3(der)
	testOOP4(bas)
	testOOP4(der)
}