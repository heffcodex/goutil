package uslice

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDiff(t *testing.T) { //nolint:dupl // ignore for test
	t.Parallel()

	type test struct {
		name                    string
		actual, desired         []int
		wantEq, wantRm, wantAdd []int
	}

	tests := []test{
		{
			name:    "empty both",
			actual:  nil,
			desired: nil,
			wantEq:  nil,
			wantRm:  nil,
			wantAdd: nil,
		},
		{
			name:    "empty actual",
			actual:  nil,
			desired: []int{1, 2, 3},
			wantEq:  nil,
			wantRm:  nil,
			wantAdd: []int{1, 2, 3},
		},
		{
			name:    "empty desired",
			actual:  []int{1, 2, 3},
			desired: nil,
			wantEq:  nil,
			wantRm:  []int{1, 2, 3},
			wantAdd: nil,
		},
		{
			name:    "same",
			actual:  []int{1, 1, 2, 3},
			desired: []int{3, 2, 1, 1},
			wantEq:  []int{1, 1, 2, 3},
			wantRm:  []int{},
			wantAdd: []int{},
		},
		{
			name:    "different",
			actual:  []int{1, 2, 3},
			desired: []int{4, 5, 6},
			wantEq:  []int{},
			wantRm:  []int{1, 2, 3},
			wantAdd: []int{4, 5, 6},
		},
		{
			name:    "semi-diff",
			actual:  []int{1, 1, 1, 1, 2, 3, 4, 5, 5, 6, 6},
			desired: []int{0, 1, 2, 2, 2, 2, 4, 6, 8, 9, 10},
			wantEq:  []int{1, 2, 4, 6},
			wantRm:  []int{1, 1, 1, 3, 5, 5, 6},
			wantAdd: []int{0, 2, 2, 2, 8, 9, 10},
		},
	}

	for _, test := range tests {
		test := test

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			gotEq, gotRm, gotAdd := Diff(test.actual, test.desired, KeyValue[int])

			assert.Equal(t, test.wantEq, gotEq)
			assert.Equal(t, test.wantRm, gotRm)
			assert.Equal(t, test.wantAdd, gotAdd)
		})
	}
}

func TestDiffIndex(t *testing.T) { //nolint:dupl // ignore for test
	t.Parallel()

	type test struct {
		name                    string
		actual, desired         []string
		wantEq, wantRm, wantAdd []int
	}

	tests := []test{
		{
			name:    "empty both",
			actual:  nil,
			desired: nil,
			wantEq:  nil,
			wantRm:  nil,
			wantAdd: nil,
		},
		{
			name:    "empty actual",
			actual:  nil,
			desired: []string{"a", "b", "c"},
			wantEq:  nil,
			wantRm:  nil,
			wantAdd: []int{0, 1, 2},
		},
		{
			name:    "empty desired",
			actual:  []string{"a", "b", "c"},
			desired: nil,
			wantEq:  nil,
			wantRm:  []int{0, 1, 2},
			wantAdd: nil,
		},
		{
			name:    "same",
			actual:  []string{"a", "a", "b", "c"},
			desired: []string{"b", "c", "a", "a"},
			wantEq:  []int{0, 1, 2, 3},
			wantRm:  []int{},
			wantAdd: []int{},
		},
		{
			name:    "different",
			actual:  []string{"a", "b", "c"},
			desired: []string{"d", "e", "f"},
			wantEq:  []int{},
			wantRm:  []int{0, 1, 2},
			wantAdd: []int{0, 1, 2},
		},
		{
			name:    "semi-diff",
			actual:  []string{"a", "a", "a", "a", "b", "c", "d", "e", "e", "f", "f"},
			desired: []string{"z", "a", "b", "b", "b", "b", "d", "f", "h", "i", "j"},
			wantEq:  []int{0, 4, 6, 9},
			wantRm:  []int{1, 2, 3, 5, 7, 8, 10},
			wantAdd: []int{0, 3, 4, 5, 8, 9, 10},
		},
	}

	for _, test := range tests {
		test := test

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			gotEq, gotRm, gotAdd := DiffIndex(test.actual, test.desired, KeyValue[string])

			assert.Equal(t, test.wantEq, gotEq)
			assert.Equal(t, test.wantRm, gotRm)
			assert.Equal(t, test.wantAdd, gotAdd)
		})
	}
}
