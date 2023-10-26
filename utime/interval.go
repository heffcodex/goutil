package utime

import (
	"time"

	"google.golang.org/genproto/googleapis/type/interval"
)

type (
	// Interval is a common time interval.
	// Should not be used directly, see InclusiveInterval and ExclusiveInterval instead.
	Interval struct {
		StartTime Time
		EndTime   Time
	}

	// InclusiveInterval is the closed Interval [StartTime, EndTime].
	InclusiveInterval Interval

	// ExclusiveInterval is the right-open Interval [StartTime, EndTime).
	// Mimics the interval.Interval proto representation.
	ExclusiveInterval Interval
)

// IntervalFromPB creates an ExclusiveInterval from the given proto representation.
func IntervalFromPB(pb *interval.Interval) ExclusiveInterval {
	return ExclusiveInterval{
		StartTime: FromPB(pb.GetStartTime()),
		EndTime:   FromPB(pb.GetEndTime()),
	}
}

// PB converts the InclusiveInterval to the proto representation.
func (i InclusiveInterval) PB() *interval.Interval {
	return i.Exclusive().PB()
}

// PB converts the ExclusiveInterval to the proto representation.
func (i ExclusiveInterval) PB() *interval.Interval {
	return &interval.Interval{
		StartTime: i.StartTime.PB(),
		EndTime:   i.EndTime.PB(),
	}
}

// IsValid checks if the InclusiveInterval does not overlap itself.
func (i InclusiveInterval) IsValid() bool {
	return i.StartTime.Compare(i.EndTime) <= 0
}

// IsValid checks if the ExclusiveInterval does not overlap itself.
func (i ExclusiveInterval) IsValid() bool {
	return i.StartTime.Compare(i.EndTime) < 0
}

// IsZero checks if the interval both StartTime and EndTime are zero.
// Implementation is same for both InclusiveInterval and ExclusiveInterval.
func (i Interval) IsZero() bool {
	return i.StartTime.IsZero() && i.EndTime.IsZero()
}

// IsZero is a wrapper for Interval.IsZero().
func (i InclusiveInterval) IsZero() bool {
	return Interval(i).IsZero()
}

// IsZero is a wrapper for Interval.IsZero().
func (i ExclusiveInterval) IsZero() bool {
	return Interval(i).IsZero()
}

// Contains returns `true` if the given Time t is contained in the InclusiveInterval.
func (i InclusiveInterval) Contains(t Time) bool {
	return i.IsValid() && i.StartTime.Compare(t) < 1 && i.EndTime.Compare(t) >= 0
}

// Contains returns `true` if the given Time t is contained in the ExclusiveInterval.
func (i ExclusiveInterval) Contains(t Time) bool {
	return i.IsValid() && i.StartTime.Compare(t) < 1 && i.EndTime.Compare(t) > 0
}

// Exclusive is an equivalent conversion of the InclusiveInterval to the ExclusiveInterval.
// Conversion precision is time.Nanosecond.
func (i InclusiveInterval) Exclusive() ExclusiveInterval {
	endTime := i.EndTime
	if !endTime.IsZero() {
		endTime = endTime.Add(time.Nanosecond)
	}

	return ExclusiveInterval{
		StartTime: i.StartTime,
		EndTime:   endTime,
	}
}

// Inclusive is an equivalent conversion of the ExclusiveInterval to the InclusiveInterval.
// Conversion precision is time.Nanosecond.
func (i ExclusiveInterval) Inclusive() InclusiveInterval {
	endTime := i.EndTime
	if !endTime.IsZero() {
		endTime = endTime.Add(-time.Nanosecond)
	}

	return InclusiveInterval{
		StartTime: i.StartTime,
		EndTime:   endTime,
	}
}
