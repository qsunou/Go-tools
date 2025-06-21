package slice

import "github.com/qsunou/Go-tools/internal/errs"

func Delete[T any](slice []T, index int) ([]T, T, error) {
	length := len(slice)

	if index < 0 || index > length {
		var zero T
		return nil, errs.NewErrIndexOutOfRange(length, index)
	}
	element := slice[index]

	for i := index; i+1 < length; i++ {
		slice[i] = slice[i+1]
	}

	return slice, element, nil
}
