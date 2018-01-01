package bookmath_test

import (
	"bookmath"
	"testing"
)

func TestSumInts(t *testing.T) {
	tests := []struct {
		values   []int
		expected int64
	}{
		{[]int{1, 2, 3}, 6},
		{[]int{1, -1, 0}, 0},
	}

	for _, testCase := range tests {
		sum := bookmath.SumInts(testCase.values...)
		if sum != testCase.expected {
			t.Error("SumInts(%v), expected=%d, actual=%d", testCase.values, testCase.expected, sum)
		}
	}
}
