package main

import (
	"testing"
)

func TestSingleRet(t *testing.T) {
	res := singleRetAsmcgocall()

	if singleRetExpected != res {
		t.Fatal("singleRetAsmcgocall...", "expected:", singleRetExpected, "got:", res)
	}

	res = singleRetCgocall()

	if singleRetExpected != res {
		t.Fatal("singleRetCgocall...", "expected:", singleRetExpected, "got:", res)
	}
}

func BenchmarkSingleRetAsmcgocall(b *testing.B) {
	for i := 0; i < b.N; i++ {
		singleRetAsmcgocall()
	}
}

func BenchmarkSignleRetCgocall(b *testing.B) {
	for i := 0; i < b.N; i++ {
		singleRetCgocall()
	}
}
