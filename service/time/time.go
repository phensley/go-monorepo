package main

import (
	"fmt"
	"time"

	"github.com/phensley/go-monorepo/pkg/logz"
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

// DoTime ..
func (s *Service) DoTime() {
	logz.Info(fmt.Sprintf("the time is %s", time.Now()))
}

func main() {
	s := NewService("time")
	s.DoTime()
}
