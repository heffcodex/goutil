package utime

import (
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"google.golang.org/genproto/googleapis/type/interval"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func TestIntervalFromPB(t *testing.T) {
	t.Parallel()

	start := time.Now()
	end := start.Add(time.Second)

	pb := interval.Interval{
		StartTime: timestamppb.New(start),
		EndTime:   timestamppb.New(end),
	}

	i := IntervalFromPB(&pb)

	assert.True(t, start.Equal(i.StartTime.Time))
	assert.True(t, end.Equal(i.EndTime.Time))
}

func TestInclusiveInterval_PB(t *testing.T) {
	t.Parallel()

	start := Now()
	end := start.Add(time.Second)

	type test struct {
		start, end Time
	}

	tests := []test{
		{},
		{start: start},
		{end: end},
		{start: start, end: end},
	}

	for i, tt := range tests {
		i := i
		tt := tt

		t.Run(strconv.Itoa(i), func(t *testing.T) {
			t.Parallel()

			pb := InclusiveInterval{StartTime: tt.start, EndTime: tt.end}.PB()

			wantEnd := tt.end
			if !wantEnd.IsZero() {
				wantEnd = wantEnd.Add(time.Nanosecond)
			}

			assert.Equal(t, tt.start.PB(), pb.GetStartTime())
			assert.Equal(t, wantEnd.PB(), pb.GetEndTime())
		})
	}
}

func TestExclusiveInterval_PB(t *testing.T) {
	t.Parallel()

	start := Now()
	end := start.Add(time.Second)

	type test struct {
		start, end Time
	}

	tests := []test{
		{},
		{start: start},
		{end: end},
		{start: start, end: end},
	}

	for i, tt := range tests {
		i := i
		tt := tt

		t.Run(strconv.Itoa(i), func(t *testing.T) {
			t.Parallel()

			pb := ExclusiveInterval{StartTime: tt.start, EndTime: tt.end}.PB()

			assert.Equal(t, tt.start.PB(), pb.GetStartTime())
			assert.Equal(t, tt.end.PB(), pb.GetEndTime())
		})
	}
}

func TestInclusiveInterval_IsValid(t *testing.T) { //nolint: dupl // ignore for test
	t.Parallel()

	type test struct {
		start, end Time
		ok         bool
	}

	now := Now()

	tests := []test{
		{ok: true},
		{start: now, ok: false},
		{end: now, ok: true},
		{start: now, end: now, ok: true},
		{start: now.Add(time.Nanosecond), end: now, ok: false},
		{start: now, end: now.Add(time.Nanosecond), ok: true},
	}

	for i, tt := range tests {
		i := i
		tt := tt

		t.Run(strconv.Itoa(i), func(t *testing.T) {
			t.Parallel()

			i := InclusiveInterval{StartTime: tt.start, EndTime: tt.end}
			assert.Equal(t, tt.ok, i.IsValid())
		})
	}
}

func TestExclusiveInterval_IsValid(t *testing.T) { //nolint: dupl // ignore for test
	t.Parallel()

	type test struct {
		start, end Time
		ok         bool
	}

	now := Now()

	tests := []test{
		{ok: false},
		{start: now, ok: false},
		{end: now, ok: true},
		{start: now, end: now, ok: false},
		{start: now.Add(time.Nanosecond), end: now, ok: false},
		{start: now, end: now.Add(time.Nanosecond), ok: true},
	}

	for i, tt := range tests {
		i := i
		tt := tt

		t.Run(strconv.Itoa(i), func(t *testing.T) {
			t.Parallel()

			i := ExclusiveInterval{StartTime: tt.start, EndTime: tt.end}
			assert.Equal(t, tt.ok, i.IsValid())
		})
	}
}

func TestInterval_IsZero(t *testing.T) {
	t.Parallel()

	assert.True(t, Interval{}.IsZero())
	assert.False(t, Interval{StartTime: Now()}.IsZero())
	assert.False(t, Interval{EndTime: Now()}.IsZero())
	assert.False(t, Interval{StartTime: Now(), EndTime: Now()}.IsZero())
}

func TestInclusiveInterval_IsZero(t *testing.T) {
	t.Parallel()

	assert.True(t, InclusiveInterval{}.IsZero())
	assert.False(t, InclusiveInterval{StartTime: Now()}.IsZero())
	assert.False(t, InclusiveInterval{EndTime: Now()}.IsZero())
	assert.False(t, InclusiveInterval{StartTime: Now(), EndTime: Now()}.IsZero())
}

func TestExclusiveInterval_IsZero(t *testing.T) {
	t.Parallel()

	assert.True(t, ExclusiveInterval{}.IsZero())
	assert.False(t, ExclusiveInterval{StartTime: Now()}.IsZero())
	assert.False(t, ExclusiveInterval{EndTime: Now()}.IsZero())
	assert.False(t, ExclusiveInterval{StartTime: Now(), EndTime: Now()}.IsZero())
}

func TestInclusiveInterval_Contains(t *testing.T) { //nolint: dupl // it's ok
	t.Parallel()

	type test struct {
		start, end, v Time
		ok            bool
	}

	now := Now()

	tests := []test{
		{start: Time{}, end: Time{}, v: Time{}, ok: true},
		{start: Time{}, end: now, v: now, ok: true},
		{start: now, end: Time{}, v: Time{}, ok: false},
		{start: now, end: Time{}, v: now, ok: false},
		{start: now, end: now, v: now, ok: true},
		{start: now, end: now, v: Time{}, ok: false},
		{start: now, end: now, v: now.Add(time.Nanosecond), ok: false},
		{start: now, end: now, v: now.Add(-time.Nanosecond), ok: false},
		{start: now.Add(-time.Second), end: now, v: now.Add(-time.Nanosecond), ok: true},
	}

	for i, tt := range tests {
		i := i
		tt := tt

		t.Run(strconv.Itoa(i), func(t *testing.T) {
			t.Parallel()

			i := InclusiveInterval{StartTime: tt.start, EndTime: tt.end}
			assert.Equal(t, tt.ok, i.Contains(tt.v))
		})
	}
}

func TestExclusiveInterval_Contains(t *testing.T) { //nolint: dupl // it's ok
	t.Parallel()

	type test struct {
		start, end, v Time
		ok            bool
	}

	now := Now()

	tests := []test{
		{start: Time{}, end: Time{}, v: Time{}, ok: false},
		{start: Time{}, end: now, v: now, ok: false},
		{start: now, end: Time{}, v: Time{}, ok: false},
		{start: now, end: Time{}, v: now, ok: false},
		{start: now, end: now, v: now, ok: false},
		{start: now, end: now, v: Time{}, ok: false},
		{start: now, end: now, v: now.Add(time.Nanosecond), ok: false},
		{start: now, end: now, v: now.Add(-time.Nanosecond), ok: false},
		{start: now.Add(-time.Second), end: now, v: now.Add(-time.Nanosecond), ok: true},
	}

	for i, tt := range tests {
		i := i
		tt := tt

		t.Run(strconv.Itoa(i), func(t *testing.T) {
			t.Parallel()

			i := ExclusiveInterval{StartTime: tt.start, EndTime: tt.end}
			assert.Equal(t, tt.ok, i.Contains(tt.v))
		})
	}
}

func TestInclusiveInterval_Exclusive(t *testing.T) {
	t.Parallel()

	type test struct {
		in         InclusiveInterval
		start, end Time
	}

	now := Now()

	tests := []test{
		{
			in:    InclusiveInterval{},
			start: Time{}, end: Time{},
		},
		{
			in:    InclusiveInterval{StartTime: now, EndTime: now},
			start: now, end: now.Add(time.Nanosecond),
		},
	}

	for i, tt := range tests {
		i := i
		tt := tt

		t.Run(strconv.Itoa(i), func(t *testing.T) {
			t.Parallel()

			out := tt.in.Exclusive()

			assert.Equal(t, tt.start, out.StartTime)
			assert.Equal(t, tt.end, out.EndTime)
		})
	}
}

func TestExclusiveInterval_Inclusive(t *testing.T) {
	t.Parallel()

	type test struct {
		in         ExclusiveInterval
		start, end Time
	}

	now := Now()

	tests := []test{
		{
			in:    ExclusiveInterval{},
			start: Time{}, end: Time{},
		},
		{
			in:    ExclusiveInterval{StartTime: now, EndTime: now},
			start: now, end: now.Add(-time.Nanosecond),
		},
	}

	for i, tt := range tests {
		i := i
		tt := tt

		t.Run(strconv.Itoa(i), func(t *testing.T) {
			t.Parallel()

			out := tt.in.Inclusive()

			assert.Equal(t, tt.start, out.StartTime)
			assert.Equal(t, tt.end, out.EndTime)
		})
	}
}
