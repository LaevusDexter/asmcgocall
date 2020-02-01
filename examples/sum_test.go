package main

import (
	"testing"
)

func BenchmarkSumAsmcgocall(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sumAsmcgocall(1, 2, 3, 4, 5)
	}
}

func BenchmarkSumCgocall(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sumCgocall(1, 2, 3, 4, 5)
	}
}

func BenchmarkEmptyAsmcgocall(b *testing.B) {
	for i := 0; i < b.N; i++ {
		emptyAsmcgocall()
	}
}

func BenchmarkEmptyCgocall(b *testing.B) {
	for i := 0; i < b.N; i++ {
		emptyCgocall()
	}
}