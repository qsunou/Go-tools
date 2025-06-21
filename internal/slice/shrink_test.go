package slice

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestShrink tests the Shrink function with various test cases to verify its correctness.
func TestShrink(t *testing.T) {
	testCases := []struct {
		name        string // Name of the test case for identification
		originCap   int    // Original capacity of the slice
		enqueueLoop int    // Number of elements appended to the slice
		expectCap   int    // Expected capacity after shrinking
	}{
		{
			name:        "less than 64", // Test case for slices with capacity less than or equal to 64
			originCap:   32,
			enqueueLoop: 6,
			expectCap:   32,
		},
		{
			name:        "less than 2048, less than 1/4", // Test case for slices with capacity less than 2048 and less than 1/4 utilization
			originCap:   1000,
			enqueueLoop: 20,
			expectCap:   500,
		},
		{
			name:        "less than 2048, more than 1/4", // Test case for slices with capacity less than 2048 and more than 1/4 utilization
			originCap:   1000,
			enqueueLoop: 400,
			expectCap:   1000,
		},
		{
			name:        "greater than 2048, less than half", // Test case for slices with capacity greater than 2048 and less than half utilization
			originCap:   3000,
			enqueueLoop: 60,
			expectCap:   1875,
		},
		{
			name:        "greater than 2048, more than half", // Test case for slices with capacity greater than 2048 and more than half utilization
			originCap:   3000,
			enqueueLoop: 2000,
			expectCap:   3000,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			l := make([]int, 0, tc.originCap)

			for i := 0; i < tc.enqueueLoop; i++ {
				l = append(l, i)
			}
			l = Shrink[int](l)
			assert.Equal(t, tc.expectCap, cap(l))
		})
	}
}