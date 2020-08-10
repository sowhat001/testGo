package main

func Fibonacci() func() int {
	sum1, sum2 := 0, 1
	return func() int {
		sum1, sum2 = sum2, sum1+sum2
		return sum1
	}
}
