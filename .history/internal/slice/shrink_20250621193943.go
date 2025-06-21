package slice

func calCapacity(slice_cap, slice_len int) (int, bool) {
	if slice_cap <= 64 {
		return slice_cap, false
	}
	if slice_cap > 2048 && (slice_cap/slice_len >= 2) {
		factor := 0.625
		return int(float32(slice_cap) * float32(factor)), true
	}
	if slice_cap <= 2048 && (slice_cap/slice_len >= 4) {
		return slice_cap / 2, true
	}
	return slice_cap, false
}
