package slice

import (
	gotools "github.com/qsunou/Go-tools"
)

func getSliceMax[T gotools.RealNumber](slice []T) T {
	res := slice[0]
	for i := 1; i < len(slice); i++ {
		if slice[i] > res {
			res = slice[i]
		}
	}
	return res
}

func getSliceMin[T gotools.RealNumber](slice []T) T {
	res := slice[0]
	for i := 1; i < len(slice); i++ {
		if slice[i] < res {
			res = slice[i]
		}
	}
	return res
}

func getSliceSum[T gotools.RealNumber](slice []T) T {
	res := T(0)
	for i := 0; i < len(slice); i++ {
		res += slice[i]
	}
	return res
}
