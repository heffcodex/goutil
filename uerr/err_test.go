package uerr

import (
	"errors"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFirst(t *testing.T) {
	t.Parallel()

	type test struct {
		chain []ErrFunc
		want  error
	}

	err1 := errors.New("err1")
	err2 := errors.New("err2")

	tests := []test{
		{
			chain: nil,
			want:  nil,
		},
		{
			chain: []ErrFunc{
				func() error { return nil },
				func() error { return err1 },
				func() error { return err2 },
			},
			want: err1,
		},
	}

	for i, tt := range tests {
		i := i
		tt := tt

		t.Run(strconv.Itoa(i), func(t *testing.T) {
			t.Parallel()
			require.ErrorIs(t, First(tt.chain...), tt.want)
		})
	}
}
