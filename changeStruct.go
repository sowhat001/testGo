package main

type test struct {
	a int
}

func noPointer(t test) {
	t.a = 444
}

func pointer(t *test) {
	t.a = 555
}
