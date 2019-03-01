package logz

// example of wrapping a logging library, providing a global logging
// endpoint (to stderr) and a configurable service
// worth it to wrap, since api is nearly 1:1 with the wrapped library?

import (
	"log"

	"go.uber.org/zap"
)

// LogFunc for forwarding calls directly to internal logger
type LogFunc = func(msg string, fields ...Field)

// Field ..
type Field = zap.Field

// Option ..
type Option = zap.Option

// Service ..
type Service struct {
	logger *zap.Logger
	Info   LogFunc
}

var (
	svc *Service
	// Info ..
	Info LogFunc
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
	Info = svc.Info
}

// NewService ..
func NewService(options ...Option) (*Service, error) {
	logger, err := zap.NewProduction(options...)
	if err != nil {
		return nil, err
	}
	return &Service{
		logger: logger,
		Info:   logger.Info,
	}, nil
}
