package main

import (
	"github.com/phensley/go-monorepo/pkg/logz"

	"github.com/phensley/go-monorepo/service/echo/internal"
)

// Service ..
type Service struct {
	Name string
}

// NewService ..
func NewService(name string) *Service {
	return &Service{
		Name: name,
	}
}

// Echo ..
func (s *Service) Echo(msg string) {
	internal.DoEcho(msg)
	logz.Info("foo")
}

func main() {
	s := NewService("echo")
	s.Echo("my message")
}
