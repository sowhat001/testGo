package main

import (
	"testing"

	"go.uber.org/dig"
)

func BenchmarkDig(b *testing.B) {
	for i := 0; i < b.N; i++ {
		c := dig.New()

		c.Provide(NewServer)
		c.Provide(NewConfig)

		c.Invoke(func(s *Server) {
			s.Run()
		})
	}
}

func BenchmarkPrimitive(b *testing.B) {
	for i := 0; i < b.N; i++ {
		c := NewConfig()
		s := NewServer(c)
		s.Run()
	}
}
