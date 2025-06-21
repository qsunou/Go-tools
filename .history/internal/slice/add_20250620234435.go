package slice

import "github.com/qsunou/Go-tools/internal/errs"

func Add[T any](slice []T, element T, index int) ([]T, error) {
	length := len(slice)

	if index < 0 || index > length {
		return nil, errs.NewErrIndexOutOfRange(length, index)
	}
	var zeroValue T
	slice = append(slice, zeroValue)
	for i := len(slice) - 1; i > index; i-- {
		if i-1 >= 0 {
			slice[i] = slice[i-1]
		}
	}

	return slice, nil
}
