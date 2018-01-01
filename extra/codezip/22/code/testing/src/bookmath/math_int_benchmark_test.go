package bookmath_test

import (
	"bookmath"
	"testing"
)

func BenchmarkSumInts(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bookmath.SumInts(1, 2, 3, 4, 5, 6, 7, 8, 9)
	}
}
