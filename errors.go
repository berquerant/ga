package ga

import (
	"fmt"
	"strings"
)

//go:generate stringer -type=ErrorCode -output errors_generated.go
type ErrorCode int

const (
	UnknownError ErrorCode = iota
	InvalidArgument
	CannotCrossover
	EvalError
	CannotMakeNextGeneration
	CannotBuildGeneration
	CannotSelectIndividual
)

// GAError is an error for this package.
type GAError interface {
	error
	// Unwrap returns the wrapped error.
	Unwrap() error
	// Trace returns stack trace.
	Trace() string
	// Code returns the error code of this.
	Code() ErrorCode
}

func NewGAError(err error, code ErrorCode, message string) GAError {
	return &gaError{
		err:     err,
		code:    code,
		message: message,
	}
}

type gaError struct {
	err     error
	message string
	code    ErrorCode
}

func (s *gaError) Code() ErrorCode { return s.code }
func (s *gaError) Unwrap() error   { return s.err }
func (s *gaError) Error() string {
	top := s
	for {
		if err, ok := top.err.(*gaError); ok {
			top = err
		} else {
			return top.message
		}
	}
}
func (s *gaError) Trace() string {
	var b strings.Builder
	top := s
	for {
		if err, ok := top.err.(*gaError); ok {
			b.WriteString(fmt.Sprintf("[%s] %s\n", top.code, top.message))
			top = err
		} else {
			b.WriteString(fmt.Sprintf("[%s] %s %v", top.code, top.message, top.err))
			return b.String()
		}
	}
}
