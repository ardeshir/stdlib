package bookmath

// SumInts adds up a bunch of ints
func SumInts(values ...int) (sum int64) {
	for _, value := range values {
		sum += int64(value)
	}
	return sum
}
