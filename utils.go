//go:build tinygo && (rp2040 || rp2350)

package tinygo_buffers

import (
	"math"
	"encoding/binary"

	tinygotypes "github.com/ralvarezdev/tinygo-types"
)

// UintToHexIndex returns the index in the ASCIIHexDigits for a given uint value
//
// Parameters:
//
//	value: The uint value to convert.
// size: The size of the uint (8, 16, 32, or 64).
// pos: The position of the hex digit to retrieve (0-based).
//
// Returns:
//
// The index in the ASCIIHexDigits for the specified hex digit, or -1 if the position is out of range.
func UintToHexIndex(value uint64, size int, pos int) int {
	if pos < 0 || pos > (size/4)-1 {
		return -1
	}
	shift := (size/4-1 - pos) * 4
	return int((value >> shift) & 0x0F)
}

// Uint8ToHex converts a uint8 value to its hexadecimal representation
//
// Parameters:
//
//	value: The uint8 value to convert.
//
// Returns:
//
// A byte slice representing the hexadecimal representation of the uint8 value.
func Uint8ToHex(value uint8) []byte {
	for c := range UintToHexBuffer {
		index := UintToHexIndex(uint64(value), 8, c)
		if index >= 0 {
			UintToHexBuffer[c] = ASCIIHexDigits[index]
		}
	}
	return UintToHexBuffer[:2]
}

// Uint16ToHex converts a uint16 value to its hexadecimal representation
//
// Parameters:
//
//	value: The uint16 value to convert.
//
// Returns:
//
// A byte slice representing the hexadecimal representation of the uint16 value.
func Uint16ToHex(value uint16) []byte {
	for c := range UintToHexBuffer {
		index := UintToHexIndex(uint64(value), 16, c)
		if index >= 0 {
			UintToHexBuffer[c] = ASCIIHexDigits[index]
		}
	}
	return UintToHexBuffer[:4]
}

// Uint32ToHex converts a uint32 value to its hexadecimal representation
//
// Parameters:
//
//	value: The uint32 value to convert.
//
// Returns:
//
// A byte slice representing the hexadecimal representation of the uint32 value.
func Uint32ToHex(value uint32) []byte {
	for c := range UintToHexBuffer {
		index := UintToHexIndex(uint64(value), 32, c)
		if index >= 0 {
			UintToHexBuffer[c] = ASCIIHexDigits[index]
		}
	}
	return UintToHexBuffer[:8]
}

// Uint64ToHex converts a uint64 value to its hexadecimal representation
//
// Parameters:
//
//	value: The uint64 value to convert.
//
// Returns:
//
// A byte slice representing the hexadecimal representation of the uint64 value.
func Uint64ToHex(value uint64) []byte {
	for c := range UintToHexBuffer {
		index := UintToHexIndex(value, 64, c)
		if index >= 0 {
			UintToHexBuffer[c] = ASCIIHexDigits[index]
		}
	}
	return UintToHexBuffer[:16]
}

// UintToDecimal converts a uint8 value to its decimal representation
//
// Parameters:
//
//	value: The uint8 value to convert.
//
// Returns:
//
// A byte slice representing the decimal representation of the uint8 value.
func UintToDecimal(value uint64) []byte {
    // Fill buffer from the end
    i := len(UintToDecimalBuffer)
    v := value
    if v == 0 {
        i--
        UintToDecimalBuffer[i] = ASCIIDecimalDigits[0]
    }
    for v > 0 && i > 0 {
        i--
        UintToDecimalBuffer[i] = ASCIIDecimalDigits[v%10]
        v /= 10
    }
    return UintToDecimalBuffer[i:]
}

// UintToDecimalFixed converts a uint value to its decimal representation with fixed width
//
// Parameters:
//
//	value: The uint value to convert.
//	width: The fixed width for the decimal representation.
//
// Returns:
//
// A byte slice representing the decimal representation of the uint value with leading zeros if necessary.
func UintToDecimalFixed(value uint64, width int) []byte {
    buffer := UintToDecimal(uint64(value))
    pad := width - len(buffer)

	// Check if padding is needed
	if pad <= 0 {
		return buffer
	}
    
	// Move existing digits to the right
	copy(UintToDecimalBuffer[pad:], buffer)
    // Prepend leading zeros
    for i := 0; i < pad; i++ {
        UintToDecimalBuffer[i] = ASCIIDecimalDigits[0]
    }
    return UintToDecimalBuffer[:width]
}

// Uint16ToBytes converts a uint16 value to an array of 2 bytes in big-endian order, storing the result in the provided buffer
//
// Parameters:
//
//	value: The uint16 value to convert.
//  buffer: A byte slice to store the resulting bytes.
//
// Returns:
//
// An error code indicating success or failure.
func Uint16ToBytes(value uint16,  buffer []byte) tinygotypes.ErrorCode {
	// Ensure the buffer has at least 2 bytes
	if len(buffer) < 2 {
		return ErrorCodeBuffersInvalidBufferSize
	}
	binary.BigEndian.PutUint16(buffer, value)
	return tinygotypes.ErrorCodeNil
}

// Uint32ToBytes converts a uint32 value to an array of 4 bytes in big-endian order, storing the result in the provided buffer
//
// Parameters:
//
//	value: The uint32 value to convert.
//  buffer: A byte slice to store the resulting bytes.
//
// Returns:
//
// An error code indicating success or failure.
func Uint32ToBytes(value uint32,  buffer []byte) tinygotypes.ErrorCode {
	// Ensure the buffer has at least 4 bytes
	if len(buffer) < 4 {
		return ErrorCodeBuffersInvalidBufferSize
	}
	binary.BigEndian.PutUint32(buffer, value)
	return tinygotypes.ErrorCodeNil
}

// Uint64ToBytes converts a uint64 value to an array of 8 bytes in big-endian order, storing the result in the provided buffer
//
// Parameters:
//
//	value: The uint64 value to convert.
//  buffer: A byte slice to store the resulting bytes.	
//
// Returns:
//
// An error code indicating success or failure.
func Uint64ToBytes(value uint64,  buffer []byte) tinygotypes.ErrorCode {
	// Ensure the buffer has at least 8 bytes
	if len(buffer) < 8 {
		return ErrorCodeBuffersInvalidBufferSize
	}
	binary.BigEndian.PutUint64(buffer, value)
	return tinygotypes.ErrorCodeNil
}

// Float32ToBytes converts a float32 value to an array of 4 bytes in big-endian order, storing the result in the provided buffer
//
// Parameters:
//
//	value: The float32 value to convert.
//  buffer: A byte slice to store the resulting bytes.
//
// Returns:
//
// An error code indicating success or failure.
func Float32ToBytes(value float32,  buffer []byte) tinygotypes.ErrorCode {
	return Uint32ToBytes(math.Float32bits(value), buffer)
}

// Float64ToBytes converts a float64 value to an array of 8 bytes in big-endian order, storing the result in the provided buffer
//
// Parameters:
//
//	value: The float64 value to convert.
//  buffer: A byte slice to store the resulting bytes.
//
// Returns:
//
// An error code indicating success or failure.
func Float64ToBytes(value float64,  buffer []byte) tinygotypes.ErrorCode {
	return Uint64ToBytes(math.Float64bits(value), buffer)
}

// Uint16ToBytesLE converts a uint16 value to an array of 2 bytes in little-endian order, storing the result in the provided buffer
//
// Parameters:
//
//	value: The uint16 value to convert.
//  buffer: A byte slice to store the resulting bytes.
//
// Returns:
//
// An error code indicating success or failure.
func Uint16ToBytesLE(value uint16,  buffer []byte) tinygotypes.ErrorCode {
	// Ensure the buffer has at least 2 bytes
	if len(buffer) < 2 {
		return ErrorCodeBuffersInvalidBufferSize
	}
	binary.LittleEndian.PutUint16(buffer, value)
	return tinygotypes.ErrorCodeNil
}

// Uint32ToBytesLE converts a uint32 value to an array of 4 bytes in little-endian order, storing the result in the provided buffer
//
// Parameters:
//
//	value: The uint32 value to convert.
//  buffer: A byte slice to store the resulting bytes.
//
// Returns:	
//
// An error code indicating success or failure.
func Uint32ToBytesLE(value uint32,  buffer []byte) tinygotypes.ErrorCode {
	// Ensure the buffer has at least 4 bytes
	if len(buffer) < 4 {
		return ErrorCodeBuffersInvalidBufferSize
	}
	binary.LittleEndian.PutUint32(buffer, value)
	return tinygotypes.ErrorCodeNil
}

// Uint64ToBytesLE converts a uint64 value to an array of 8 bytes in little-endian order, storing the result in the provided buffer
//
// Parameters:
//
//	value: The uint64 value to convert.
//  buffer: A byte slice to store the resulting bytes.
//
// Returns:
//
// An error code indicating success or failure.
func Uint64ToBytesLE(value uint64,  buffer []byte) tinygotypes.ErrorCode {
	// Ensure the buffer has at least 8 bytes
	if len(buffer) < 8 {
		return ErrorCodeBuffersInvalidBufferSize
	}
	binary.LittleEndian.PutUint64(buffer, value)
	return tinygotypes.ErrorCodeNil
}

// Float32ToBytesLE converts a float32 value to an array of 4 bytes in little-endian order, storing the result in the provided buffer
//
// Parameters:
//
//	value: The float32 value to convert.
//  buffer: A byte slice to store the resulting bytes.
//
// Returns:
//
// An error code indicating success or failure.
func Float32ToBytesLE(value float32,  buffer []byte) tinygotypes.ErrorCode {
	return Uint32ToBytesLE(math.Float32bits(value), buffer)
}

// Float64ToBytesLE converts a float64 value to an array of 8 bytes in little-endian order, storing the result in the provided buffer
//
// Parameters:
//
//	value: The float64 value to convert.
//  buffer: A byte slice to store the resulting bytes.
//
// Returns:
//
// An error code indicating success or failure.
func Float64ToBytesLE(value float64,  buffer []byte) tinygotypes.ErrorCode {
	return Uint64ToBytesLE(math.Float64bits(value), buffer)
}

// BytesToUint16 converts a byte slice to a uint16 value
//
// Parameters:
//
//	data: A byte slice containing at least 2 bytes.
//
// Returns:
//
// The uint16 value represented by the first 2 bytes of the input slice, or an error code if the input is invalid.
func BytesToUint16(data []byte) (uint16, tinygotypes.ErrorCode) {
	if len(data) < 2 {
		return 0, ErrorCodeBuffersInvalidBufferSize
	}
	return (uint16(data[0]) << 8) | uint16(data[1]), tinygotypes.ErrorCodeNil
}

// BytesToUint32 converts a byte slice to a uint32 value
//
// Parameters:
//
//	data: A byte slice containing at least 4 bytes.
//
// Returns:
//
// The uint32 value represented by the first 4 bytes of the input slice, or an error code if the input is invalid.
func BytesToUint32(data []byte) (uint32, tinygotypes.ErrorCode) {
	if len(data) < 4 {
		return 0, ErrorCodeBuffersInvalidBufferSize
	}
	return (uint32(data[0]) << 24) | (uint32(data[1]) << 16) | (uint32(data[2]) << 8) | uint32(data[3]), tinygotypes.ErrorCodeNil
}

// BytesToUint64 converts a byte slice to a uint64 value
//
// Parameters:
//
//	data: A byte slice containing at least 8 bytes.
//
// Returns:
//
// The uint64 value represented by the first 8 bytes of the input slice, or an error code if the input is invalid.
func BytesToUint64(data []byte) (uint64, tinygotypes.ErrorCode) {
	if len(data) < 8 {
		return 0, ErrorCodeBuffersInvalidBufferSize
	}
	return (uint64(data[0]) << 56) | (uint64(data[1]) << 48) | (uint64(data[2]) << 40) | (uint64(data[3]) << 32) |
		(uint64(data[4]) << 24) | (uint64(data[5]) << 16) | (uint64(data[6]) << 8) | uint64(data[7]), tinygotypes.ErrorCodeNil
}

// BytesToFloat32 converts a byte slice to a float32 value
//
// Parameters:
//
//	data: A byte slice containing at least 4 bytes.
//
// Returns:
//
// The float32 value represented by the first 4 bytes of the input slice, or an error code if the input is invalid.
func BytesToFloat32(data []byte) (float32, tinygotypes.ErrorCode) {
	u, err := BytesToUint32(data)
	if err != tinygotypes.ErrorCodeNil {
		return 0, err
	}
	return math.Float32frombits(u), tinygotypes.ErrorCodeNil
}

// BytesToFloat64 converts a byte slice to a float64 value
//
// Parameters:
//
//	data: A byte slice containing at least 8 bytes.
//
// Returns:
//
// The float64 value represented by the first 8 bytes of the input slice, or an error code if the input is invalid.
func BytesToFloat64(data []byte) (float64, tinygotypes.ErrorCode) {
	u, err := BytesToUint64(data)
	if err != tinygotypes.ErrorCodeNil {
		return 0, err
	}
	return math.Float64frombits(u), tinygotypes.ErrorCodeNil
}

// BytesToUint16LE converts a byte slice to a uint16 value in little-endian order
//
// Parameters:
//
//	data: A byte slice containing at least 2 bytes.
//
// Returns:
//
// The uint16 value represented by the first 2 bytes of the input slice in little-endian order, or an error code if the input is invalid.
func BytesToUint16LE(data []byte) (uint16, tinygotypes.ErrorCode) {
	if len(data) < 2 {
		return 0, ErrorCodeBuffersInvalidBufferSize
	}
	return (uint16(data[1]) << 8) | uint16(data[0]), tinygotypes.ErrorCodeNil
}

// BytesToUint32LE converts a byte slice to a uint32 value in little-endian order
//
// Parameters:
//
//	data: A byte slice containing at least 4 bytes.
//
// Returns:
//
// The uint32 value represented by the first 4 bytes of the input slice in little-endian order, or an error code if the input is invalid.
func BytesToUint32LE(data []byte) (uint32, tinygotypes.ErrorCode) {
	if len(data) < 4 {
		return 0, ErrorCodeBuffersInvalidBufferSize
	}
	return (uint32(data[3]) << 24) | (uint32(data[2]) << 16) | (uint32(data[1]) << 8) | uint32(data[0]), tinygotypes.ErrorCodeNil
}

// BytesToUint64LE converts a byte slice to a uint64 value in little-endian order
//
// Parameters:
//
//	data: A byte slice containing at least 8 bytes.
//
// Returns:
//
// The uint64 value represented by the first 8 bytes of the input slice in little-endian order, or an error code if the input is invalid.
func BytesToUint64LE(data []byte) (uint64, tinygotypes.ErrorCode) {
	if len(data) < 8 {
		return 0, ErrorCodeBuffersInvalidBufferSize
	}
	return (uint64(data[7]) << 56) | (uint64(data[6]) << 48) | (uint64(data[5]) << 40) | (uint64(data[4]) << 32) |
		(uint64(data[3]) << 24) | (uint64(data[2]) << 16) | (uint64(data[1]) << 8) | uint64(data[0]), tinygotypes.ErrorCodeNil
}

// BytesToFloat32LE converts a byte slice to a float32 value in little-endian order
//
// Parameters:
//
//	data: A byte slice containing at least 4 bytes.
//
// Returns:
//
// The float32 value represented by the first 4 bytes of the input slice in little-endian order, or an error code if the input is invalid.
func BytesToFloat32LE(data []byte) (float32, tinygotypes.ErrorCode) {
	u, err := BytesToUint32LE(data)
	if err != tinygotypes.ErrorCodeNil {
		return 0, err
	}
	return math.Float32frombits(u), tinygotypes.ErrorCodeNil
}

// BytesToFloat64LE converts a byte slice to a float64 value in little-endian order
//
// Parameters:
//
//	data: A byte slice containing at least 8 bytes.
//
// Returns:
//
// The float64 value represented by the first 8 bytes of the input slice in little-endian order, or an error code if the input is invalid.
func BytesToFloat64LE(data []byte) (float64, tinygotypes.ErrorCode) {
	u, err := BytesToUint64LE(data)
	if err != tinygotypes.ErrorCodeNil {
		return 0, err
	}
	return math.Float64frombits(u), tinygotypes.ErrorCodeNil
}