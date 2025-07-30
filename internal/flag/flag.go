package flag

const (
	FlagRed   = 1 << 0 // 0001
	FlagGreen = 1 << 1 // 0010
	FlagBlue  = 1 << 2 // 0100
)

func PrintColors(flags int) {
	if flags&FlagRed != 0 {
		println("Red enabled")
	}
	if flags&FlagGreen != 0 {
		println("Green enabled")
	}
	if flags&FlagBlue != 0 {
		println("Blue enabled")
	}
}

// usage:
// PrintColors(FlagRed | FlagBlue) // shows "Red enabled", "Blue enabled"
