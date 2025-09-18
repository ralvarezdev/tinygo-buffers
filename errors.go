package tinygo_buffers

import (
	tinygoerrors "github.com/ralvarezdev/tinygo-errors"
)

const (
	// ErrorCodeBuffersStartNumber is the starting number for error code buffers
	ErrorCodeBuffersStartNumber = 4000
)

const (
	ErrorCodeBuffersInvalidBufferSize tinygoerrors.ErrorCode = ErrorCodeBuffersStartNumber + iota
	ErrorCodeBuffersTooMuchPrecisionDigitsForFloat64
)
