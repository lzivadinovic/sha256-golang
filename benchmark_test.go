package main

import (
	"math/rand"
	"testing"
)

func BenchmarkAddUInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		addUInt(rand.Uint32(), rand.Uint32())
	}
}

func BenchmarkAddDirect(b *testing.B) {
	for i := 0; i < b.N; i++ {
		addDirect(rand.Uint32(), rand.Uint32())
	}
}

func BenchmarkRightRotate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rightRotate(rand.Uint32(), rand.Uint32())
	}
}

func BenchmarkGoRotate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		goRightRotate(rand.Uint32(), rand.Uint32())
	}
}

func BenchmarkMessageSchedule(b *testing.B) {
	S := rand.NewSource(10)
	R := rand.New(S)
	for i := 0; i < b.N; i++ {
		token := make([]byte, 64)
		R.Read(token)
		messageSchedule(token)
	}
}
