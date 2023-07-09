package main

import (
	"encoding/binary"
	"math/rand"
	"testing"
)

var (
	S   = rand.NewSource(10)
	R   = rand.New(S)
	res []uint32
)

func BenchmarkAddUInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		addUInt(R.Uint32(), R.Uint32())
	}
}

func BenchmarkAddDirect(b *testing.B) {
	for i := 0; i < b.N; i++ {
		addDirect(R.Uint32(), R.Uint32())
	}
}

func BenchmarkRightRotate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rightRotate(R.Uint32(), R.Uint32())
	}
}

func BenchmarkGoRotate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		goRightRotate(R.Uint32(), R.Uint32())
	}
}

func BenchmarkMessageSchedule(b *testing.B) {
	token := make([]byte, 64)
	R.Read(token)
	for i := 0; i < b.N; i++ {
		messageSchedule(token)
	}
}

func BenchmarkRoundDeclare(b *testing.B) {
	W := make([]uint32, 8)
	binary.Read(R, binary.BigEndian, &W)
	for i := 0; i < b.N; i++ {
		W = roundDeclare(W, 961987163, 3221900128)
	}
	res = W
}

func BenchmarkRoundMutate(b *testing.B) {
	W := make([]uint32, 8)
	binary.Read(R, binary.BigEndian, &W)
	for i := 0; i < b.N; i++ {
		roundMutate(W, 961987163, 3221900128)
	}
	res = W
}
