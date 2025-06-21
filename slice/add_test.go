package slice

import (
	"fmt"
	"testing"

	"github.com/qsunou/Go-tools/internal/errs"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	// Add mainly relies on internal/slice.Add to ensure correctness
	testCases := []struct {
		name      string
		slice     []int
		addVal    int
		index     int
		wantSlice []int
		wantErr   error
	}{
		{
			name:      "index 0",
			slice:     []int{123, 100},
			addVal:    233,
			index:     0,
			wantSlice: []int{233, 123, 100},
		},
		{
			name:    "index -1",
			slice:   []int{123, 100},
			index:   -1,
			wantErr: errs.NewErrIndexOutOfRange(2, -1),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res, err := Add(tc.slice, tc.addVal, tc.index)
			assert.Equal(t, tc.wantErr, err)
			if err != nil {
				return
			}
			assert.Equal(t, tc.wantSlice, res)
		})
	}
}

func ExampleAdd() {
	res, _ := Add[int]([]int{1, 2, 3, 4}, 233, 2)
	fmt.Println(res)
	_, err := Add[int]([]int{1, 2, 3, 4}, 233, -1)
	fmt.Println(err)
	// Output:
	// [1 2 233 3 4]
	// ekit: index out of range, length 4, index -1
}
