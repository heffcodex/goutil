package uarray

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDiff(t *testing.T) {
	type test struct {
		name                    string
		actual, desired         []int
		wantEq, wantAdd, wantRm []int
	}

	tests := []test{
		{
			name:    "empty both",
			actual:  nil,
			desired: nil,
			wantEq:  nil,
			wantAdd: nil,
			wantRm:  nil,
		},
		{
			name:    "empty actual",
			actual:  nil,
			desired: []int{1, 2, 3},
			wantEq:  nil,
			wantAdd: []int{1, 2, 3},
			wantRm:  nil,
		},
		{
			name:    "empty desired",
			actual:  []int{1, 2, 3},
			desired: nil,
			wantEq:  nil,
			wantAdd: nil,
			wantRm:  []int{1, 2, 3},
		},
		{
			name:    "same",
			actual:  []int{1, 2, 3},
			desired: []int{3, 2, 1},
			wantEq:  []int{1, 2, 3},
			wantAdd: []int{},
			wantRm:  []int{},
		},
		{
			name:    "different",
			actual:  []int{1, 2, 3},
			desired: []int{4, 5, 6},
			wantEq:  []int{},
			wantAdd: []int{4, 5, 6},
			wantRm:  []int{1, 2, 3},
		},
		{
			name:    "semi-diff",
			actual:  []int{1, 2, 3, 4, 5, 6},
			desired: []int{0, 1, 2, 4, 6, 8, 9, 10},
			wantEq:  []int{1, 2, 4, 6},
			wantAdd: []int{0, 8, 9, 10},
			wantRm:  []int{3, 5},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			gotEq, gotAdd, gotRm := Diff(test.actual, test.desired, CmpValue[int])
			require.Equal(t, test.wantEq, gotEq)
			require.EqualValues(t, test.wantAdd, gotAdd)
			require.EqualValues(t, test.wantRm, gotRm)
		})
	}
}
