package utime

import (
	"time"

	"google.golang.org/genproto/googleapis/type/interval"
)

// Interval mimics the interval.Interval proto representation with addition of optional InclusiveEnd flag and some useful methods.
type Interval struct {
	StartTime    Time
	EndTime      Time
	InclusiveEnd bool
}

// IntervalFromPB creates an Interval from the given proto representation.
func IntervalFromPB(pb *interval.Interval) Interval {
	return Interval{
		StartTime:    FromPB(pb.GetStartTime()),
		EndTime:      FromPB(pb.GetEndTime()),
		InclusiveEnd: false,
	}
}

// PB converts the Interval to the proto representation.
func (i Interval) PB() *interval.Interval {
	endTime := i.EndTime
	if i.InclusiveEnd && !endTime.IsZero() {
		endTime = endTime.Add(time.Nanosecond)
	}

	return &interval.Interval{
		StartTime: i.StartTime.PB(),
		EndTime:   endTime.PB(),
	}
}

// IsValid checks if the interval does not overlap itself.
func (i Interval) IsValid() bool {
	return (i.InclusiveEnd && i.StartTime.Compare(i.EndTime) <= 0) || i.StartTime.Compare(i.EndTime) < 0
}

// Contains returns true if the given Time t is contained in the interval.
func (i Interval) Contains(t Time) bool {
	return i.IsValid() && i.StartTime.Compare(t) < 1 && ((i.InclusiveEnd && i.EndTime.Compare(t) >= 0) || i.EndTime.Compare(t) > 0)
}

// Inclusive returns a new interval ensuring that InclusiveEnd flag set to true and EndTime corrected respectively if needed.
func (i Interval) Inclusive() Interval {
	if i.InclusiveEnd {
		return i
	}

	endTime := i.EndTime
	if !endTime.IsZero() {
		endTime = endTime.Add(-time.Nanosecond)
	}

	return Interval{
		StartTime:    i.StartTime,
		EndTime:      endTime,
		InclusiveEnd: true,
	}
}
