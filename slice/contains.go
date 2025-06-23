package slice

// Contains checks if a given item exists in the source slice.
func Contains[T comparable](src []T, item T) bool {
	return ContainsFunc[T](src, func(src T) bool {
		return src == item
	})
}

// ContainsFunc checks if any element in the source slice satisfies the provided equality function.
func ContainsFunc[T any](src []T, equal func(src T) bool) bool {
	for _, v := range src {
		if equal(v) {
			return true
		}
	}
	return false
}

// ContainsAny checks if any element from the destination slice exists in the source slice.
func ContainsAny[T comparable](src, dst []T) bool {
	srcMap := toMap[T](src)
	for _, v := range dst {
		if _, exist := srcMap[v]; exist {
			return true
		}
	}
	return false
}

// ContainsAnyFunc checks if any element from the destination slice matches an element in the source slice based on the provided equality function.
// ContainsAny is better than ContainsAnyFunc in terms of performance, as it uses a map for lookups.
func ContainsAnyFunc[T any](src, dst []T, equal equalFunc[T]) bool {
	for _, valDst := range dst {
		for _, valSrc := range src {
			if equal(valSrc, valDst) {
				return true
			}
		}
	}
	return false
}

// ContainsAll checks if all elements from the destination slice exist in the source slice.
func ContainsAll[T comparable](src, dst []T) bool {
	srcMap := toMap[T](src)
	for _, v := range dst {
		if _, exist := srcMap[v]; !exist {
			return false
		}
	}
	return true
}

// ContainsAllFunc checks if all elements from the destination slice match an element in the source slice based on the provided equality function.
// ContainsAll is better than ContainsAllFunc in terms of performance, as it uses a map for lookups.
func ContainsAllFunc[T any](src, dst []T, equal equalFunc[T]) bool {
	for _, valDst := range dst {
		if !ContainsFunc[T](src, func(src T) bool {
			return equal(src, valDst)
		}) {
			return false
		}
	}
	return true
}
