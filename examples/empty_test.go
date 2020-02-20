package main

import (
	"testing"
)

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
