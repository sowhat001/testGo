// dig package
package main

import (
	"go.uber.org/dig"
)

type Config struct {
}

func NewConfig() *Config {
	return &Config{}
}

type Server struct {
	c *Config
}

func NewServer(c *Config) *Server {
	return &Server{c: c}
}

func (s *Server) Run() {
}

func main() {
	c := dig.New()

	c.Provide(NewServer)
	c.Provide(NewConfig)

	c.Invoke(func(s *Server) {
		s.Run()
	})
}
