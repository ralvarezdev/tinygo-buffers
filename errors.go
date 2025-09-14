package tinygo_buffers

import (
	tinygotypes "github.com/ralvarezdev/tinygo-types"
)

const (
	// ErrorCodeBuffersStartNumber is the starting number for error code buffers
	ErrorCodeBuffersStartNumber = 4000
)

const (
	ErrorCodeBuffersInvalidBufferSize tinygotypes.ErrorCode = ErrorCodeBuffersStartNumber + iota
	
)