package tinygo_buffers

var (
	// WhitespaceBuffer is a byte slice representing a whitespace character
	WhitespaceBuffer = []byte(" ")

	// NewlineBuffer is a byte slice representing a newline character
	NewlineBuffer = []byte("\n")

	// TabBuffer is a byte slice representing a tab character
	TabBuffer = []byte("\t")

	// TwoPointsBuffer is a byte slice representing two points
	TwoPointsBuffer = []byte(":")

	// DotBuffer is a byte slice representing a dot character
	DotBuffer = []byte(".")

	// HexPrefix is the prefix for error codes
	HexPrefix = []byte("0x")

	// Float64Buffer is the buffer used for float64 messages
	Float64Buffer = [8]byte{}

	// ASCIIHexDigits is a byte slice representing ASCII hex digits
	ASCIIHexDigits = []byte("0123456789ABCDEF")

	// ASCIIDecimalDigits is a byte slice representing ASCII decimal digits
	ASCIIDecimalDigits = []byte("0123456789")

	// UintToHexBuffer is a buffer used for converting uint64 to hex
	UintToHexBuffer = [16]byte{}

	// UintToDecimalBuffer is a buffer used for converting uint64 to decimal
	UintToDecimalBuffer = [20]byte{}
)