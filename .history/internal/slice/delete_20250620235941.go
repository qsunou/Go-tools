package slice

import "github.com/qsunou/Go-tools/internal/errs"

func Delete[T any](slice []T, index int) ([]T, T, error) {
	length := len(slice)

	if index < 0 || index > length {
		var zero T
		return nil, zero, errs.NewErrIndexOutOfRange(length, index)
	}
	element := slice[index]

	for i := index; i+1 < length; i++ {
		slice[i] = slice[i+1]
	}
	slice = slice[:length-1] // Remove the last element since it's now a duplicate
	return slice, element, nil
}
