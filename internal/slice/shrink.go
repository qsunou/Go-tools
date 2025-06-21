package slice

// calCapacity calculates the new capacity of a slice based on its current length and capacity.
// It returns the new capacity and a boolean indicating whether the capacity has been changed.
func calCapacity(sliceCap, sliceLen int) (int, bool) {
	if sliceCap <= 64 {
		return sliceCap, false
	}
	if sliceCap > 2048 && (sliceCap/sliceLen >= 2) {
		factor := 0.625
		return int(float32(sliceCap) * float32(factor)), true
	}
	if sliceCap <= 2048 && (sliceCap/sliceLen >= 4) {
		return sliceCap / 2, true
	}
	return sliceCap, false
}

// Shrink reduces the capacity of a slice if it exceeds certain thresholds.
// It creates a new slice with the calculated capacity and copies the elements from the original slice.
func Shrink[T any](slice []T) []T {
	sliceCap, sliceLen := cap(slice), len(slice)
	n, changed := calCapacity(sliceCap, sliceLen)

	if !changed {
		return slice
	}

	newSlice := make([]T, 0, n)
	newSlice = append(newSlice, slice...)
	return newSlice
}