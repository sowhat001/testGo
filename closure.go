package main

func fibonacci() func() int {
	count, sum1, sum2 := 0, 0, 1
	return func() int {
		switch count {
		case 0:
			count++
		default:
			sum1, sum2 = sum2, sum1+sum2
		}
		return sum1
	}
}
