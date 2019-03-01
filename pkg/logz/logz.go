package logz

// example of wrapping a logging library, providing a global logging
// endpoint (to stderr) and a configurable service
// worth it to wrap, since api is nearly 1:1 with the wrapped library?

import (
	"log"

	"go.uber.org/zap"
)

// Service ..
type Service struct {
	logger *zap.Logger
}

// Field ..
type Field = zap.Field

// Option ..
type Option = zap.Option

var (
	svc *Service
)

var (
	// String constructs a string field with given key and value
	String = zap.String
	// Int constructs an int field with given key and value
	Int = zap.Int
	// Int64 constructs an int64 field with given key and value
	Int64 = zap.Int64
)

func init() {
	var err error
	svc, err = NewService()
	if err != nil {
		log.Fatalf("Failed to initialize static logz service: %s", err)
	}
}

// NewService ..
func NewService(options ...Option) (*Service, error) {
	logger, err := zap.NewProduction(options...)
	if err != nil {
		return nil, err
	}
	return &Service{
		logger: logger,
	}, nil
}

// Info ..
func (s *Service) Info(msg string, fields ...Field) {
	s.logger.Info(msg, fields...)
}

// Info ..
func Info(msg string, fields ...Field) {
	svc.Info(msg, fields...)
}
