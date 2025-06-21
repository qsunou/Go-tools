package slice

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

func Shrink[T any](slice []T) []T {
	sliceCap, sliceLen := cap(slice), len(slice)
	n, changed := calCapacity(sliceCap, sliceLen)

	if !changed {
		return slice
	}

	new_slice := make([]T, 0, n)
	new_slice = append(new_slice, slice...)
	return new_slice
}
