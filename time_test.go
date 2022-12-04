package goutil

import (
	"encoding/json"
	"testing"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/stretchr/testify/require"
)

func TestNewTimestamp(t *testing.T) {
	now := time.Now()
	ts := NewTimestamp(now)

	require.Equal(t, now, ts.Time())
}

func TestNewTimestampFromPB(t *testing.T) {
	now := time.Now().UTC().Round(time.Nanosecond)
	pb := timestamppb.New(now)
	ts := NewTimestampFromPB(pb)

	require.Equal(t, now, ts.Time())
}

func TestTimestamp_Time(t *testing.T) {
	now := time.Now()
	ts := Timestamp{t: now}

	require.Equal(t, now, ts.Time())
}

func TestTimestamp_PB(t *testing.T) {
	now := time.Now().UTC()
	ts := Timestamp{t: now}
	pb := ts.PB()

	require.Equal(t, now, pb.AsTime())
}

func TestTimestamp_MarshalJSON(t *testing.T) {
	now := time.Now()
	s := struct {
		Timestamp Timestamp `json:"timestamp"`
	}{
		Timestamp: Timestamp{t: time.Now()},
	}

	data, err := json.Marshal(s)
	require.NoError(t, err)

	require.Equal(t, `{"timestamp":"`+now.Format(TimestampMarshalFormat)+`"}`, string(data))
}

func TestTimestamp_UnmarshalJSON(t *testing.T) {
	now := time.Now().Round(time.Second)
	s := struct {
		Timestamp Timestamp `json:"timestamp"`
	}{}

	err := json.Unmarshal([]byte(`{"timestamp":"`+now.Format(TimestampMarshalFormat)+`"}`), &s)
	require.NoError(t, err)

	require.Equal(t, now, s.Timestamp.Time())
}

func TestTimestamp_MarshalText(t *testing.T) {
	now := time.Now()
	ts := Timestamp{t: now}
	data, err := ts.MarshalText()

	require.NoError(t, err)
	require.Equal(t, now.Format(TimestampMarshalFormat), string(data))
}

func TestTimestamp_UnmarshalText(t *testing.T) {
	now := time.Now().Round(time.Second)
	ts := Timestamp{}

	err := ts.UnmarshalText([]byte(now.Format(TimestampMarshalFormat)))
	require.NoError(t, err)

	require.Equal(t, now, ts.Time())
}

func TestTimestamp_MarshalBinary(t *testing.T) {
	now := time.Now()
	ts := Timestamp{t: now}

	data, err := ts.MarshalBinary()
	require.NoError(t, err)

	expected, err := now.MarshalBinary()
	require.NoError(t, err)

	require.Equal(t, expected, data)
}

func TestTimestamp_UnmarshalBinary(t *testing.T) {
	now := time.Now().Round(time.Second)
	ts := Timestamp{}

	nowBin, err := now.MarshalBinary()
	require.NoError(t, err)

	err = ts.UnmarshalBinary(nowBin)
	require.NoError(t, err)

	require.Equal(t, now, ts.Time())
}

func TestTimestamp_StartOfDay(t *testing.T) {
	now := time.Now()
	ts := Timestamp{t: now}
	startOfDay := ts.StartOfDay()

	require.Equal(
		t,
		time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location()),
		startOfDay.Time(),
	)
}

func TestTimestamp_EndOfDay(t *testing.T) {
	now := time.Now()
	ts := Timestamp{t: now}
	endOfDay := ts.EndOfDay()

	require.Equal(
		t,
		time.Date(now.Year(), now.Month(), now.Day(), 23, 59, 59, 999999999, now.Location()),
		endOfDay.Time(),
	)
}

func TestTimestamp_StartOfWeek(t *testing.T) {
	now := time.Now()
	ts := Timestamp{t: now}
	startOfWeek := ts.StartOfWeek()

	require.Equal(t, startOfWeek.Time().Weekday(), StartOfWeek)
}

func TestTimestamp_EndOfWeek(t *testing.T) {
	now := time.Now()
	ts := Timestamp{t: now}
	endOfWeek := ts.EndOfWeek()

	require.Equal(t, endOfWeek.Time().Weekday(), EndOfWeek)
}

func TestTimestamp_StartOfMonth(t *testing.T) {
	now := time.Now()
	ts := Timestamp{t: now}
	startOfMonth := ts.StartOfMonth()

	require.Equal(
		t,
		time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location()),
		startOfMonth.Time(),
	)
}

func TestTimestamp_EndOfMonth(t *testing.T) {
	now := time.Now()
	ts := Timestamp{t: now}
	endOfMonth := ts.EndOfMonth()

	require.Equal(
		t,
		time.Date(now.Year(), now.Month()+1, 0, 23, 59, 59, 999999999, now.Location()),
		endOfMonth.Time(),
	)
}
