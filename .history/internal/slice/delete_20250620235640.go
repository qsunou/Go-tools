package slice

import "github.com/qsunou/Go-tools/internal/errs"

func Delete[T any](slice []T, index int) ([]T, T, error) {
	length := len(slice)

	if index < 0 || index > length {
		return nil, errs.NewErrIndexOutOfRange(length, index)
	}
	element := slice[index]

	for i := len(slice) - 1; i > index; i-- {
		if i-1 >= 0 {
			slice[i] = slice[i-1]
		}
	}

	return slice, element, nil
}
