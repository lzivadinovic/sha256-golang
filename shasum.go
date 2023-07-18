package main

import (
	"encoding/binary"
	"math/bits"
)

var (
	// K is roundConstant is initial value for compress function
	// it is equal to first 32 bits of decimal part of sqrt of first 32 prime numbers
	// Just use them as const's
	K = []uint32{0x428a2f98, 0x71374491, 0xb5c0fbcf, 0xe9b5dba5, 0x3956c25b, 0x59f111f1, 0x923f82a4, 0xab1c5ed5,
		0xd807aa98, 0x12835b01, 0x243185be, 0x550c7dc3, 0x72be5d74, 0x80deb1fe, 0x9bdc06a7, 0xc19bf174,
		0xe49b69c1, 0xefbe4786, 0x0fc19dc6, 0x240ca1cc, 0x2de92c6f, 0x4a7484aa, 0x5cb0a9dc, 0x76f988da,
		0x983e5152, 0xa831c66d, 0xb00327c8, 0xbf597fc7, 0xc6e00bf3, 0xd5a79147, 0x06ca6351, 0x14292967,
		0x27b70a85, 0x2e1b2138, 0x4d2c6dfc, 0x53380d13, 0x650a7354, 0x766a0abb, 0x81c2c92e, 0x92722c85,
		0xa2bfe8a1, 0xa81a664b, 0xc24b8b70, 0xc76c51a3, 0xd192e819, 0xd6990624, 0xf40e3585, 0x106aa070,
		0x19a4c116, 0x1e376c08, 0x2748774c, 0x34b0bcb5, 0x391c0cb3, 0x4ed8aa4a, 0x5b9cca4f, 0x682e6ff3,
		0x748f82ee, 0x78a5636f, 0x84c87814, 0x8cc70208, 0x90befffa, 0xa4506ceb, 0xbef9a3f7, 0xc67178f2}
)

func addDirect(x, y uint32) uint32 {
	return x + y
}

func goRightRotate(n, k uint32) uint32 {
	return bits.RotateLeft32(n, -int(k))
}

func choice(x, y, z uint32) uint32 {
	return (x & y) ^ (^x & z)
}

func majority(x, y, z uint32) uint32 {
	return (x & y) ^ (x & z) ^ (y & z)
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

// I would use this in implementation!
func roundMutate(state []uint32, roundConst uint32, scheduleWord uint32) {
	t1 := state[7] + bigSigma1(state[4]) + choice(state[4], state[5], state[6]) + roundConst + scheduleWord
	t2 := bigSigma0(state[0]) + majority(state[0], state[1], state[2])
	for i := len(state) - 1; i >= 1; i-- {
		state[i] = state[i-1]
	}
	state[4] += t1
	state[0] = t1 + t2
}

func compressBlock(initialState []uint32, block []byte) []uint32 {
	W := messageSchedule(block)
	state := make([]uint32, len(initialState))
	// "deep" copy from initialState to state
	copy(state, initialState)
	for i := 0; i < len(W); i++ {
		roundMutate(state, W[i], K[i])
	}
	// there is no native element-wise slice operation function
	// you can use gonum, but you need to change types to Vector
	for i := 0; i < len(state); i++ {
		initialState[i] += state[i]
	}
	return initialState
}

func main() {

}
