package slice

import (
	"fmt"
	"testing"

	gotools "github.com/qsunou/Go-tools"
	"github.com/stretchr/testify/assert"
)

func TestMax(t *testing.T) {
	testCases := []struct {
		name  string
		input []Integer
		want  Integer
	}{
		{
			name:  "value",
			input: []Integer{1},
			want:  1,
		},
		{
			name:  "values",
			input: []Integer{2, 3, 1},
			want:  3,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := getSliceMax[Integer](tc.input)
			assert.Equal(t, tc.want, res)
		})
	}

	assert.Panics(t, func() {
		getSliceMax[int](nil)
	})
	assert.Panics(t, func() {
		getSliceMax[int]([]int{})
	})

	testMaxTypes[uint](t)
	testMaxTypes[uint8](t)
	testMaxTypes[uint16](t)
	testMaxTypes[uint32](t)
	testMaxTypes[uint64](t)
	testMaxTypes[int](t)
	testMaxTypes[int8](t)
	testMaxTypes[int16](t)
	testMaxTypes[int32](t)
	testMaxTypes[int64](t)
	testMaxTypes[float32](t)
	testMaxTypes[float64](t)
}

func TestMin(t *testing.T) {
	testCases := []struct {
		name  string
		input []Integer
		want  Integer
	}{
		{
			name:  "value",
			input: []Integer{3},
			want:  3,
		},
		{
			name:  "values",
			input: []Integer{3, 1, 2},
			want:  1,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := getSliceMin[Integer](tc.input)
			assert.Equal(t, tc.want, res)
		})
	}

	assert.Panics(t, func() {
		getSliceMin[int](nil)
	})
	assert.Panics(t, func() {
		getSliceMin[int]([]int{})
	})

	testMinTypes[uint](t)
	testMinTypes[uint8](t)
	testMinTypes[uint16](t)
	testMinTypes[uint32](t)
	testMinTypes[uint64](t)
	testMinTypes[int](t)
	testMinTypes[int8](t)
	testMinTypes[int16](t)
	testMinTypes[int32](t)
	testMinTypes[int64](t)
	testMinTypes[float32](t)
	testMinTypes[float64](t)
}

func TestSum(t *testing.T) {
	testCases := []struct {
		name  string
		input []Integer
		want  Integer
	}{
		{
			name: "nil",
		},
		{
			name:  "empty",
			input: []Integer{},
		},
		{
			name:  "value",
			input: []Integer{1},
			want:  1,
		},
		{
			name:  "values",
			input: []Integer{1, 2, 3},
			want:  6,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := getSliceSum[Integer](tc.input)
			assert.Equal(t, tc.want, res)
		})
	}

	testSumTypes[uint](t)
	testSumTypes[uint8](t)
	testSumTypes[uint16](t)
	testSumTypes[uint32](t)
	testSumTypes[uint64](t)
	testSumTypes[int](t)
	testSumTypes[int8](t)
	testSumTypes[int16](t)
	testSumTypes[int32](t)
	testSumTypes[int64](t)
	testSumTypes[float32](t)
	testSumTypes[float64](t)
}

// testMaxTypes is used to test all types that satisfy the Max method constraint
func testMaxTypes[T gotools.RealNumber](t *testing.T) {
	res := getSliceMax[T]([]T{1, 2, 3})
	assert.Equal(t, T(3), res)
}

// testMinTypes is used to test all types that satisfy the Min method constraint
func testMinTypes[T gotools.RealNumber](t *testing.T) {
	res := getSliceMin[T]([]T{1, 2, 3})
	assert.Equal(t, T(1), res)
}

// testSumTypes is used to test all types that satisfy the Sum method constraint
func testSumTypes[T gotools.RealNumber](t *testing.T) {
	res := getSliceSum[T]([]T{1, 2, 3})
	assert.Equal(t, T(6), res)
}

type Integer int

func ExampleSum() {
	res := getSliceSum[int]([]int{1, 2, 3})
	fmt.Println(res)
	res = getSliceSum[int](nil)
	fmt.Println(res)
	// Output:
	// 6
	// 0
}

func ExampleMin() {
	res := getSliceMin[int]([]int{1, 2, 3})
	fmt.Println(res)
	// Output:
	// 1
}

func ExampleMax() {
	res := getSliceMax[int]([]int{1, 2, 3})
	fmt.Println(res)
	// Output:
	// 3
}
