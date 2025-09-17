package tinygo_buffers

import (
	"math"
	"encoding/binary"

	tinygoerrors "github.com/ralvarezdev/tinygo-errors"
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

// IntToDecimal converts an int64 value to its decimal representation
//
// Parameters:
//
//	value: The int64 value to convert.
//
// Returns:
//
// A byte slice representing the decimal representation of the int64 value, including a minus sign if negative.
func IntToDecimal(value int64) []byte {
    i := len(IntToDecimalBuffer)
    negative := value < 0
    v := value
    if negative {
        v = -v
    }
    if v == 0 {
        i--
        IntToDecimalBuffer[i] = ASCIIDecimalDigits[0]
    }
    for v > 0 && i > 0 {
        i--
        IntToDecimalBuffer[i] = ASCIIDecimalDigits[v%10]
        v /= 10
    }
    if negative && i > 0 {
        i--
        IntToDecimalBuffer[i] = '-'
    }
    return IntToDecimalBuffer[i:]
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

// Float64ToDecimal converts a float64 value to its decimal representation with specified precision
//
// Parameters:
//
//	value: The float64 value to convert.
//	precision: The number of digits after the decimal point.
//
// Returns:
//
// The number of bytes written to the buffer and an error code indicating success or failure
func Float64ToDecimal(value float64, precision int) ([]byte, tinygoerrors.ErrorCode) {
	// Get the integer and fractional parts
	intPart := int64(value)
    fracPart := value - float64(intPart)
    idx := 0

    // Convert integer part
    intBuf := IntToDecimal(intPart)
    copy(Float64ToDecimalBuffer[idx:], intBuf)
    idx += len(intBuf)

    // Add dot
    Float64ToDecimalBuffer[idx] = '.'
    idx++

	// Check precision limit
	if len(Float64ToDecimalBuffer)-idx < precision {
		return nil, ErrorCodeBuffersTooMuchPrecisionDigitsForFloat64
	}

    // Convert fractional part
    for i := 0; i < precision; i++ {
        fracPart *= 10
        digit := int(fracPart)
        Float64ToDecimalBuffer[idx] = byte('0' + digit)
        idx++
        fracPart -= float64(digit)
    }

    return Float64ToDecimalBuffer[:idx], tinygoerrors.ErrorCodeNil
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
func Uint16ToBytes(value uint16,  buffer []byte) tinygoerrors.ErrorCode {
	// Ensure the buffer has at least 2 bytes
	if len(buffer) < 2 {
		return ErrorCodeBuffersInvalidBufferSize
	}
	binary.BigEndian.PutUint16(buffer, value)
	return tinygoerrors.ErrorCodeNil
}

// Int16ToBytes converts an int16 value to an array of 2 bytes in big-endian order, storing the result in the provided buffer
//
// Parameters:
//
//	value: The int16 value to convert.
//  buffer: A byte slice to store the resulting bytes.	
// Returns:
//
// An error code indicating success or failure.
func Int16ToBytes(value int16,  buffer []byte) tinygoerrors.ErrorCode {
	return Uint16ToBytes(uint16(value), buffer)
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
func Uint32ToBytes(value uint32,  buffer []byte) tinygoerrors.ErrorCode {
	// Ensure the buffer has at least 4 bytes
	if len(buffer) < 4 {
		return ErrorCodeBuffersInvalidBufferSize
	}
	binary.BigEndian.PutUint32(buffer, value)
	return tinygoerrors.ErrorCodeNil
}

// Int32ToBytes converts an int32 value to an array of 4 bytes in big-endian order, storing the result in the provided buffer
//
// Parameters:
//
//	value: The int32 value to convert.
//  buffer: A byte slice to store the resulting bytes.			
//
// Returns:
//
// An error code indicating success or failure.
func Int32ToBytes(value int32,  buffer []byte) tinygoerrors.ErrorCode {
	return Uint32ToBytes(uint32(value), buffer)
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
func Uint64ToBytes(value uint64,  buffer []byte) tinygoerrors.ErrorCode {
	// Ensure the buffer has at least 8 bytes
	if len(buffer) < 8 {
		return ErrorCodeBuffersInvalidBufferSize
	}
	binary.BigEndian.PutUint64(buffer, value)
	return tinygoerrors.ErrorCodeNil
}

// Int64ToBytes converts an int64 value to an array of 8 bytes in big-endian order, storing the result in the provided buffer
//
// Parameters:
//
//	value: The int64 value to convert.
//  buffer: A byte slice to store the resulting bytes.
//
// Returns:
//
// An error code indicating success or failure.
func Int64ToBytes(value int64,  buffer []byte) tinygoerrors.ErrorCode {
	return Uint64ToBytes(uint64(value), buffer)
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
func Float32ToBytes(value float32,  buffer []byte) tinygoerrors.ErrorCode {
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
func Float64ToBytes(value float64,  buffer []byte) tinygoerrors.ErrorCode {
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
func Uint16ToBytesLE(value uint16,  buffer []byte) tinygoerrors.ErrorCode {
	// Ensure the buffer has at least 2 bytes
	if len(buffer) < 2 {
		return ErrorCodeBuffersInvalidBufferSize
	}
	binary.LittleEndian.PutUint16(buffer, value)
	return tinygoerrors.ErrorCodeNil
}

// Int16ToBytesLE converts an int16 value to an array of 2 bytes in little-endian order, storing the result in the provided buffer
//
// Parameters:
//
//	value: The int16 value to convert.
//  buffer: A byte slice to store the resulting bytes.	
//
// Returns:
//
// An error code indicating success or failure.
func Int16ToBytesLE(value int16,  buffer []byte) tinygoerrors.ErrorCode {
	return Uint16ToBytesLE(uint16(value), buffer)
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
func Uint32ToBytesLE(value uint32,  buffer []byte) tinygoerrors.ErrorCode {
	// Ensure the buffer has at least 4 bytes
	if len(buffer) < 4 {
		return ErrorCodeBuffersInvalidBufferSize
	}
	binary.LittleEndian.PutUint32(buffer, value)
	return tinygoerrors.ErrorCodeNil
}

// Int32ToBytesLE converts an int32 value to an array of 4 bytes in little-endian order, storing the result in the provided buffer
//
// Parameters:
//
//	value: The int32 value to convert.
//  buffer: A byte slice to store the resulting bytes.			
//
// Returns:	
//
// An error code indicating success or failure.
func Int32ToBytesLE(value int32,  buffer []byte) tinygoerrors.ErrorCode {
	return Uint32ToBytesLE(uint32(value), buffer)
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
func Uint64ToBytesLE(value uint64,  buffer []byte) tinygoerrors.ErrorCode {
	// Ensure the buffer has at least 8 bytes
	if len(buffer) < 8 {
		return ErrorCodeBuffersInvalidBufferSize
	}
	binary.LittleEndian.PutUint64(buffer, value)
	return tinygoerrors.ErrorCodeNil
}

// Int64ToBytesLE converts an int64 value to an array of 8 bytes in little-endian order, storing the result in the provided buffer
//
// Parameters:
//
//	value: The int64 value to convert.
//  buffer: A byte slice to store the resulting bytes.
//
// Returns:
//
// An error code indicating success or failure.
func Int64ToBytesLE(value int64,  buffer []byte) tinygoerrors.ErrorCode {
	return Uint64ToBytesLE(uint64(value), buffer)
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
func Float32ToBytesLE(value float32,  buffer []byte) tinygoerrors.ErrorCode {
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
func Float64ToBytesLE(value float64,  buffer []byte) tinygoerrors.ErrorCode {
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
func BytesToUint16(data []byte) (uint16, tinygoerrors.ErrorCode) {
	if len(data) < 2 {
		return 0, ErrorCodeBuffersInvalidBufferSize
	}
	return (uint16(data[0]) << 8) | uint16(data[1]), tinygoerrors.ErrorCodeNil
}

// BytesToInt16 converts a byte slice to an int16 value
//
// Parameters:	
//
//	data: A byte slice containing at least 2 bytes.
//
// Returns:
//
// The int16 value represented by the first 2 bytes of the input slice, or an error code if the input is invalid.
func BytesToInt16(data []byte) (int16, tinygoerrors.ErrorCode) {
	u, err := BytesToUint16(data)
	if err != tinygoerrors.ErrorCodeNil {
		return 0, err
	}
	return int16(u), tinygoerrors.ErrorCodeNil
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
func BytesToUint32(data []byte) (uint32, tinygoerrors.ErrorCode) {
	if len(data) < 4 {
		return 0, ErrorCodeBuffersInvalidBufferSize
	}
	return (uint32(data[0]) << 24) | (uint32(data[1]) << 16) | (uint32(data[2]) << 8) | uint32(data[3]), tinygoerrors.ErrorCodeNil
}

// BytesToInt32 converts a byte slice to an int32 value
//
// Parameters:
//
//	data: A byte slice containing at least 4 bytes.
//
// Returns:
//
// The int32 value represented by the first 4 bytes of the input slice, or an error code if the input is invalid.
func BytesToInt32(data []byte) (int32, tinygoerrors.ErrorCode) {
	u, err := BytesToUint32(data)
	if err != tinygoerrors.ErrorCodeNil {
		return 0, err
	}
	return int32(u), tinygoerrors.ErrorCodeNil
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
func BytesToUint64(data []byte) (uint64, tinygoerrors.ErrorCode) {
	if len(data) < 8 {
		return 0, ErrorCodeBuffersInvalidBufferSize
	}
	return (uint64(data[0]) << 56) | (uint64(data[1]) << 48) | (uint64(data[2]) << 40) | (uint64(data[3]) << 32) |
		(uint64(data[4]) << 24) | (uint64(data[5]) << 16) | (uint64(data[6]) << 8) | uint64(data[7]), tinygoerrors.ErrorCodeNil
}

// BytesToInt64 converts a byte slice to an int64 value
//
// Parameters:
//
//	data: A byte slice containing at least 8 bytes.
//
// Returns:
//
// The int64 value represented by the first 8 bytes of the input slice, or an error code if the input is invalid.
func BytesToInt64(data []byte) (int64, tinygoerrors.ErrorCode) {
	u, err := BytesToUint64(data)
	if err != tinygoerrors.ErrorCodeNil {
		return 0, err
	}
	return int64(u), tinygoerrors.ErrorCodeNil
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
func BytesToFloat32(data []byte) (float32, tinygoerrors.ErrorCode) {
	u, err := BytesToUint32(data)
	if err != tinygoerrors.ErrorCodeNil {
		return 0, err
	}
	return math.Float32frombits(u), tinygoerrors.ErrorCodeNil
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
func BytesToFloat64(data []byte) (float64, tinygoerrors.ErrorCode) {
	u, err := BytesToUint64(data)
	if err != tinygoerrors.ErrorCodeNil {
		return 0, err
	}
	return math.Float64frombits(u), tinygoerrors.ErrorCodeNil
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
func BytesToUint16LE(data []byte) (uint16, tinygoerrors.ErrorCode) {
	if len(data) < 2 {
		return 0, ErrorCodeBuffersInvalidBufferSize
	}
	return (uint16(data[1]) << 8) | uint16(data[0]), tinygoerrors.ErrorCodeNil
}

// BytesToInt16LE converts a byte slice to an int16 value in little-endian order
//
// Parameters:
//
//	data: A byte slice containing at least 2 bytes.
//
// Returns:
//
// The int16 value represented by the first 2 bytes of the input slice in little-endian order, or an error code if the input is invalid.
func BytesToInt16LE(data []byte) (int16, tinygoerrors.ErrorCode) {
	u, err := BytesToUint16LE(data)
	if err != tinygoerrors.ErrorCodeNil {
		return 0, err
	}
	return int16(u), tinygoerrors.ErrorCodeNil
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
func BytesToUint32LE(data []byte) (uint32, tinygoerrors.ErrorCode) {
	if len(data) < 4 {
		return 0, ErrorCodeBuffersInvalidBufferSize
	}
	return (uint32(data[3]) << 24) | (uint32(data[2]) << 16) | (uint32(data[1]) << 8) | uint32(data[0]), tinygoerrors.ErrorCodeNil
}

// BytesToInt32LE converts a byte slice to an int32 value in little-endian order
//
// Parameters:
//
//	data: A byte slice containing at least 4 bytes.
//
// Returns:
//
// The int32 value represented by the first 4 bytes of the input slice in little-endian order, or an error code if the input is invalid.
func BytesToInt32LE(data []byte) (int32, tinygoerrors.ErrorCode) {
	u, err := BytesToUint32LE(data)
	if err != tinygoerrors.ErrorCodeNil {
		return 0, err
	}
	return int32(u), tinygoerrors.ErrorCodeNil
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
func BytesToUint64LE(data []byte) (uint64, tinygoerrors.ErrorCode) {
	if len(data) < 8 {
		return 0, ErrorCodeBuffersInvalidBufferSize
	}
	return (uint64(data[7]) << 56) | (uint64(data[6]) << 48) | (uint64(data[5]) << 40) | (uint64(data[4]) << 32) |
		(uint64(data[3]) << 24) | (uint64(data[2]) << 16) | (uint64(data[1]) << 8) | uint64(data[0]), tinygoerrors.ErrorCodeNil
}

// BytesToInt64LE converts a byte slice to an int64 value in little-endian order
//
// Parameters:
//
//	data: A byte slice containing at least 8 bytes.
//
// Returns:		
//
// The int64 value represented by the first 8 bytes of the input slice in little-endian order, or an error code if the input is invalid.
func BytesToInt64LE(data []byte) (int64, tinygoerrors.ErrorCode) {
	u, err := BytesToUint64LE(data)
	if err != tinygoerrors.ErrorCodeNil {
		return 0, err
	}
	return int64(u), tinygoerrors.ErrorCodeNil
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
func BytesToFloat32LE(data []byte) (float32, tinygoerrors.ErrorCode) {
	u, err := BytesToUint32LE(data)
	if err != tinygoerrors.ErrorCodeNil {
		return 0, err
	}
	return math.Float32frombits(u), tinygoerrors.ErrorCodeNil
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
func BytesToFloat64LE(data []byte) (float64, tinygoerrors.ErrorCode) {
	u, err := BytesToUint64LE(data)
	if err != tinygoerrors.ErrorCodeNil {
		return 0, err
	}
	return math.Float64frombits(u), tinygoerrors.ErrorCodeNil
}