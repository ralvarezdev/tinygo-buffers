//go:build tinygo && (rp2040 || rp2350)

package tinygo_buffers

// UintToHexIndex returns the index in the asciiHexDigits for a given uint value
//
// Parameters:
//
//	value: The uint value to convert.
// size: The size of the uint (8, 16, 32, or 64).
// pos: The position of the hex digit to retrieve (0-based).
//
// Returns:
//
// The index in the asciiHexDigits for the specified hex digit, or -1 if the position is out of range.
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
		index := l.UintToHexIndex(uint64(value), 8, c)
		if index >= 0 {
			UintToHexBuffer[c] = asciiHexDigits[index]
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
		index := l.UintToHexIndex(uint64(value), 16, c)
		if index >= 0 {
			UintToHexBuffer[c] = asciiHexDigits[index]
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
		index := l.UintToHexIndex(uint64(value), 32, c)
		if index >= 0 {
			UintToHexBuffer[c] = asciiHexDigits[index]
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
		index := l.UintToHexIndex(value, 64, c)
		if index >= 0 {
			UintToHexBuffer[c] = asciiHexDigits[index]
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
        UintToDecimalBuffer[i] = asciiDecimalDigits[0]
    }
    for v > 0 && i > 0 {
        i--
        UintToDecimalBuffer[i] = asciiDecimalDigits[v%10]
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
    buffer := l.UintToDecimal(uint64(value))
    pad := width - len(buffer)

	// Check if padding is needed
	if pad <= 0 {
		return buffer
	}
    
	// Move existing digits to the right
	copy(UintToDecimalBuffer[pad:], buffer)
    // Prepend leading zeros
    for i := 0; i < pad; i++ {
        UintToDecimalBuffer[i] = asciiDecimalDigits[0]
    }
    return UintToDecimalBuffer[:width]
}
