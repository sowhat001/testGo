package main

import "testing"

func BenchmarkNoDI(b *testing.B) {
	for i := 0; i < b.N; i++ {
		message := NewMessage("hello world")
		greeter := NewGreeter(message)
		event := NewEvent(greeter)
		event.Start()
	}
}

func BenchmarkDI(b *testing.B) {
	for i := 0; i < b.N; i++ {
		event := InitializeEvent("hello_world")
		event.Start()
	}
}
