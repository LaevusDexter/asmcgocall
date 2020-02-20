package main

import (
	"testing"
)

func sum(prms ...int) (res int) {
	for _, i := range prms {
		res += i
	}

	return
}

var sumExpected = sum(1, 2, 3, 4, 5)

func TestSum(t *testing.T) {
	res := sumAsmcgocall(1, 2, 3, 4, 5)

	if sumExpected != int(res) {
		t.Fatal("sumAsmcgocall...", "expected:", sumExpected, "got:", res)
	}

	res = sumCgocall(1, 2, 3, 4, 5)
	if sumExpected != int(res) {
		t.Fatal("sumCgocall...", "expected:", sumExpected, "got:", res)
	}
}

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
