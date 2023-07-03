package main

import (
	"encoding/binary"
	"math/bits"
)

func addDirect(x, y uint32) uint32 {
	return x + y
}

func goRightRotate(n, k uint32) uint32 {
	return bits.RotateLeft32(n, -int(k))
}

func littleSigma0(x uint32) uint32 {
	return bits.RotateLeft32(x, -7) ^
		bits.RotateLeft32(x, -18) ^
		x>>3
}

func littleSigma1(x uint32) uint32 {
	return bits.RotateLeft32(x, -17) ^
		bits.RotateLeft32(x, -19) ^
		x>>10
}

func bigSigma0(x uint32) uint32 {
	return bits.RotateLeft32(x, -2) ^
		bits.RotateLeft32(x, -13) ^
		bits.RotateLeft32(x, -22)
}

func bigSigma1(x uint32) uint32 {
	return bits.RotateLeft32(x, -6) ^
		bits.RotateLeft32(x, -11) ^
		bits.RotateLeft32(x, -25)
}

func messageSchedule(x []byte) []uint32 {
	// Create slice of length 64 to hold data
	W := make([]uint32, 64)
	//// in golang, byte is alias of uint8, and 64*8 = 512 (2+2 is 4 minus 1 that's 3 quick maths)
	if len(x) != 64 {
		panic("func messageSchedule internal error: message not 64bytes long!")
	}
	for i := 0; i < 64; i++ {
		if i < 16 {
			W[i] = binary.BigEndian.Uint32(x[i*4 : (i+1)*4])
			//W[i] = uint32(x[i*4])<<24 | uint32(x[i*4+1])<<16 | uint32(x[i*4+2])<<8 | uint32(x[i*4+3])
		} else {
			W[i] = littleSigma1(W[i-2]) + W[i-7] + littleSigma0(W[i-15]) + W[i-16]
		}
	}
	return W
}

func main() {
}
