// This is just placeholder file for my implementation of specific functions
// when possible, i'm using built in implementation

package main

func addUInt(x, y uint32) uint32 {
	return uint32(uint64(x) + uint64(y))
}

func rightRotate(n, k uint32) uint32 {
	// eg (with uint8 type)  00101100 with shift 3
	// size ^ is 8 (num of bits)
	var mask, size uint32
	size = 32
	// find n mod 8, no need to shift multiple times
	shift := k % size
	if shift == 0 {
		return n
	}
	// create mask for right side
	// mask is 00000111 (2^3-1)
	mask = (1 << shift) - 1
	// just move everything right
	// 00101100 >> 3 = 00000101
	lo := n >> shift
	// AND number and mask so we can get lower part we lost when we
	// shifted to calculate lo
	// n AND mask = 00000100, now move that 8-3 spaces to left
	// because of wrap around
	// 10000000
	hi := (n & mask) << (size - shift)
	// or just + add it
	return hi | lo
}

func roundDeclare(state []uint32, roundConst uint32, scheduleWord uint32) []uint32 {
	t1 := state[7] + bigSigma1(state[4]) + choice(state[4], state[5], state[6]) + roundConst + scheduleWord
	t2 := bigSigma0(state[0]) + majority(state[0], state[1], state[2])
	W := make([]uint32, 8)
	for i := 1; i < len(W); i++ {
		W[i] = state[i-1]
	}
	// Fancy, but slow!
	//W := append([]uint32{t1 + t2}, state...)
	//W[4] += t1
	// return W[0:7]
	W[4] += t1
	W[0] = t1 + t2
	return W
}
