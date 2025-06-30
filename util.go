package btk

import "strconv"

func hex(hex string) (r, g, b, a uint8) {
	red, err := strconv.ParseUint(hex[0:2], 16, 8)
	assert(err)
	green, err := strconv.ParseUint(hex[2:4], 16, 8)
	assert(err)
	blue, err := strconv.ParseUint(hex[4:6], 16, 8)
	assert(err)
	return uint8(red), uint8(green), uint8(blue), 0xFF
}

func assert(err error) {
	if err != nil {
		panic(err)
	}
}
