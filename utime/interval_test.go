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

func TestInterval_PB(t *testing.T) {
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
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			t.Run("inclusiveEnd", func(t *testing.T) {
				i := Interval{
					StartTime:    tt.start,
					EndTime:      tt.end,
					InclusiveEnd: true,
				}

				pb := i.PB()

				wantEnd := tt.end
				if !tt.end.IsZero() {
					wantEnd = tt.end.Add(time.Nanosecond)
				}

				assert.Equal(t, tt.start.PB(), pb.GetStartTime())
				assert.Equal(t, wantEnd.PB(), pb.GetEndTime())
			})

			t.Run("exclusiveEnd", func(t *testing.T) {
				i := Interval{
					StartTime:    tt.start,
					EndTime:      tt.end,
					InclusiveEnd: false,
				}

				pb := i.PB()

				assert.Equal(t, tt.start.PB(), pb.GetStartTime())
				assert.Equal(t, tt.end.PB(), pb.GetEndTime())
			})
		})
	}
}

func TestInterval_IsValid(t *testing.T) {
	type test struct {
		start, end Time
		ie, ok     bool
	}

	now := Now()

	tests := []test{
		{ie: false, ok: false},
		{ie: true, ok: true},
		{start: now, ie: false, ok: false},
		{start: now, ie: true, ok: false},
		{end: now, ie: false, ok: true},
		{end: now, ie: true, ok: true},
		{start: now, end: now, ie: false, ok: false},
		{start: now, end: now, ie: true, ok: true},
		{start: now.Add(time.Nanosecond), end: now, ie: false, ok: false},
		{start: now.Add(time.Nanosecond), end: now, ie: false, ok: false},
		{start: now, end: now.Add(time.Nanosecond), ie: false, ok: true},
		{start: now, end: now.Add(time.Nanosecond), ie: true, ok: true},
	}

	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			i := Interval{
				StartTime:    tt.start,
				EndTime:      tt.end,
				InclusiveEnd: tt.ie,
			}

			assert.Equal(t, tt.ok, i.IsValid())
		})
	}
}

func TestInterval_Contains(t *testing.T) {
	type test struct {
		start, end, v Time
		ie, ok        bool
	}

	now := Now()

	tests := []test{
		{start: Time{}, end: Time{}, v: Time{}, ie: true, ok: true},
		{start: Time{}, end: Time{}, v: Time{}, ie: true, ok: true},
		{start: Time{}, end: now, v: now, ie: true, ok: true},
		{start: now, end: Time{}, v: Time{}, ie: false, ok: false},
		{start: now, end: Time{}, v: Time{}, ie: true, ok: false},
		{start: now, end: Time{}, v: now, ie: false, ok: false},
		{start: now, end: Time{}, v: now, ie: true, ok: false},
		{start: now, end: now, v: now, ie: true, ok: true},
		{start: now, end: now, v: now, ie: false, ok: false},
		{start: now, end: now, v: Time{}, ie: true, ok: false},
		{start: now, end: now, v: Time{}, ie: false, ok: false},
		{start: now, end: now, v: now.Add(time.Nanosecond), ie: false, ok: false},
		{start: now, end: now, v: now.Add(time.Nanosecond), ie: true, ok: false},
		{start: now, end: now, v: now.Add(-time.Nanosecond), ie: false, ok: false},
		{start: now, end: now, v: now.Add(-time.Nanosecond), ie: true, ok: false},
		{start: now.Add(-time.Second), end: now, v: now.Add(-time.Nanosecond), ie: true, ok: true},
	}

	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			i := Interval{
				StartTime:    tt.start,
				EndTime:      tt.end,
				InclusiveEnd: tt.ie,
			}

			assert.Equal(t, tt.ok, i.Contains(tt.v))
		})
	}
}
