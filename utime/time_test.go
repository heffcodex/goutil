package utime

import (
	"encoding/json"
	"testing"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/stretchr/testify/require"
)

func TestFromStdTime(t *testing.T) {
	now := time.Now()
	ts := FromStdTime(now)

	require.Equal(t, now, ts.Time)
}

func TestFromPB(t *testing.T) {
	now := time.Now().UTC().Round(time.Nanosecond)
	pb := timestamppb.New(now)
	ts := FromPB(pb)

	require.Equal(t, now, ts.Time)
}

func TestNow(t *testing.T) {
	now := time.Now()
	ts := Now()

	require.True(t, now.Before(ts.Time))
}

func TestTime_StdTime(t *testing.T) {
	now := time.Now()
	ts := Time{Time: now}

	require.Equal(t, now, ts.StdTime())
}

func TestTime_PB(t *testing.T) {
	now := time.Now().UTC()
	ts := Time{Time: now}

	require.Equal(t, now, ts.PB().AsTime())
}

func TestTime_MarshalJSON(t *testing.T) {
	forTime := func(t *testing.T, tt time.Time) {
		s := struct {
			T Time `json:"t"`
		}{
			T: Time{Time: tt},
		}

		data, err := json.Marshal(s)
		require.NoError(t, err)

		expected := `{"t":null}`
		if !tt.IsZero() {
			expected = `{"t":"` + tt.Format(MarshalFormat) + `"}`
		}

		require.Equal(t, expected, string(data))
	}

	t.Run("now", func(t *testing.T) { forTime(t, time.Now()) })
	t.Run("zero", func(t *testing.T) { forTime(t, time.Time{}) })
}

func TestTimestamp_UnmarshalJSON(t *testing.T) {
	s := struct {
		T Time `json:"t"`
	}{}

	t.Run(
		"now", func(t *testing.T) {
			now := time.Now().Round(time.Second)
			err := json.Unmarshal([]byte(`{"t":"`+now.Format(MarshalFormat)+`"}`), &s)
			require.NoError(t, err)
			require.Equal(t, now, s.T.Time)
		},
	)

	t.Run(
		"zero", func(t *testing.T) {
			err := json.Unmarshal([]byte(`{"t":null}`), &s)
			require.NoError(t, err)
			require.True(t, s.T.IsZero())
		},
	)
}

func TestTimestamp_MarshalText(t *testing.T) {
	t.Run(
		"now", func(t *testing.T) {
			now := time.Now()
			ts := Time{Time: now}
			data, err := ts.MarshalText()

			require.NoError(t, err)
			require.Equal(t, now.Format(MarshalFormat), string(data))
		},
	)

	t.Run(
		"zero", func(t *testing.T) {
			ts := Time{}
			data, err := ts.MarshalText()

			require.NoError(t, err)
			require.Empty(t, data)
		},
	)
}

func TestTimestamp_UnmarshalText(t *testing.T) {
	ts := Time{}

	t.Run(
		"now", func(t *testing.T) {
			now := time.Now().Round(time.Second)
			err := ts.UnmarshalText([]byte(now.Format(MarshalFormat)))
			require.NoError(t, err)
			require.Equal(t, now, ts.Time)
		},
	)

	t.Run(
		"zero", func(t *testing.T) {
			err := ts.UnmarshalText([]byte(""))
			require.NoError(t, err)
			require.True(t, ts.IsZero())
		},
	)
}

func TestTimestamp_MarshalBinary(t *testing.T) {
	t.Run(
		"now", func(t *testing.T) {
			now := time.Now()
			ts := Time{Time: now}

			data, err := ts.MarshalBinary()
			require.NoError(t, err)

			expected, err := now.MarshalBinary()
			require.NoError(t, err)

			require.Equal(t, expected, data)
		},
	)

	t.Run(
		"zero", func(t *testing.T) {
			ts := Time{}

			data, err := ts.MarshalBinary()
			require.NoError(t, err)

			expected, err := time.Time{}.MarshalBinary()
			require.NoError(t, err)

			require.Equal(t, expected, data)
		},
	)
}

func TestTimestamp_UnmarshalBinary(t *testing.T) {
	ts := Time{}

	t.Run(
		"now", func(t *testing.T) {
			now := time.Now().Round(time.Second)

			nowBin, err := now.MarshalBinary()
			require.NoError(t, err)

			err = ts.UnmarshalBinary(nowBin)
			require.NoError(t, err)

			require.Equal(t, now, ts.Time)
		},
	)

	t.Run(
		"zero", func(t *testing.T) {
			nowBin, err := time.Time{}.MarshalBinary()
			require.NoError(t, err)

			err = ts.UnmarshalBinary(nowBin)
			require.NoError(t, err)

			require.True(t, ts.IsZero())
		},
	)
}

func TestTimestamp_StartOfDay(t *testing.T) {
	now := time.Now()
	ts := Time{Time: now}
	startOfDay := ts.StartOfDay()

	require.Equal(
		t,
		time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location()),
		startOfDay.Time,
	)
}

func TestTimestamp_EndOfDay(t *testing.T) {
	now := time.Now()
	ts := Time{Time: now}
	endOfDay := ts.EndOfDay()

	require.Equal(
		t,
		time.Date(now.Year(), now.Month(), now.Day(), 23, 59, 59, 999999999, now.Location()),
		endOfDay.Time,
	)
}

func TestTimestamp_StartOfWeek(t *testing.T) {
	now := time.Now()
	ts := Time{Time: now}
	startOfWeek := ts.StartOfWeek()

	require.Equal(t, startOfWeek.Time.Weekday(), StartOfWeek)
}

func TestTimestamp_EndOfWeek(t *testing.T) {
	now := time.Now()
	ts := Time{Time: now}
	endOfWeek := ts.EndOfWeek()

	require.Equal(t, endOfWeek.Time.Weekday(), EndOfWeek())
}

func TestTimestamp_StartOfMonth(t *testing.T) {
	now := time.Now()
	ts := Time{Time: now}
	startOfMonth := ts.StartOfMonth()

	require.Equal(
		t,
		time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location()),
		startOfMonth.Time,
	)
}

func TestTimestamp_EndOfMonth(t *testing.T) {
	now := time.Now()
	ts := Time{Time: now}
	endOfMonth := ts.EndOfMonth()

	require.Equal(
		t,
		time.Date(now.Year(), now.Month()+1, 0, 23, 59, 59, 999999999, now.Location()),
		endOfMonth.Time,
	)
}
