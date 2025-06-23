package slice

func Map[Src any, Dst any](src []Src, m func(idx int, src Src) Dst) []Dst {
	dst := make([]Dst, len(src))
	for i, s := range src {
		dst[i] = m(i, s)
	}
	return dst
}

// FilterMap executes filtering and transformation.
// If the second return value of m is false, we ignore the first return value.
// Even if the second return value is false, subsequent elements will still be iterated.
func filterMap[Src any, Dst any](src []Src, m func(idx int, src Src) (Dst, bool)) []Dst {
	res := make([]Dst, 0, len(src))
	for i, s := range src {
		dst, ok := m(i, s)
		if ok {
			res = append(res, dst)
		}
	}
	return res
}

// MapSliceToMap converts []Ele to map[Key]Ele.
// The function fn to extract Key from Ele is provided by the user.
//
// Note:
// If i < j,
// Let:
//
//	key_i := fn(elements[i])
//	key_j := fn(elements[j])
//
// If key_i == key_j, in the result map resultMap,
// resultMap[key_i] = val_j
//
// Even if the input slice is nil, it guarantees that the returned map is an empty map rather than nil.
func ToMap[Ele any, Key comparable](
	elements []Ele,
	fn func(element Ele) Key,
) map[Key]Ele {
	return ToMapV(
		elements,
		func(element Ele) (Key, Ele) {
			return fn(element), element
		})
}

// MapSliceToMapWithValues converts []Ele to map[Key]Val.
// The function fn to extract Key and Val from Ele is provided by the user.
//
// Note:
// If i < j,
// Let:
//
//	key_i, val_i := fn(elements[i])
//	key_j, val_j := fn(elements[j])
//
// If key_i == key_j, in the result map resultMap,
// resultMap[key_i] = val_j
//
// Even if the input slice is nil, it guarantees that the returned map is an empty map rather than nil.
func ToMapV[Ele any, Key comparable, Val any](
	elements []Ele,
	fn func(element Ele) (Key, Val),
) (resultMap map[Key]Val) {
	resultMap = make(map[Key]Val, len(elements))
	for _, element := range elements {
		k, v := fn(element)
		resultMap[k] = v
	}
	return
}

// ConvertSliceToSet converts a slice into a `map` where the keys are the elements of the slice and the values are empty structs `struct{}` to save memory.
// **Function Description:**
// - Input: a slice containing elements of a comparable type `T`;
// - Output: a `map[T]struct{}` used for quickly checking the existence of an element;
// - Using the empty struct `struct{}` as the value helps avoid storing unnecessary data and reduces memory consumption.
func toMap[T comparable](src []T) map[T]struct{} {
	var dataMap = make(map[T]struct{}, len(src))
	for _, v := range src {
		dataMap[v] = struct{}{}
	}
	return dataMap
}
