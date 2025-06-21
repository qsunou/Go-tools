package slice

func Add[T any](slice []T, element T, index int) ([]T, error) {
	length := len(slice)

	if index < 0 || index > length {
		return nil, NewErrIndexOutOfRange(length, index)
	}
	return slice, nil
}
