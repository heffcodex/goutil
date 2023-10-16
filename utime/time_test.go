package utime

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/timestamppb"
)

const (
	testNsec = 7
)

func testTime() time.Time {
	return time.Date(2022, 2, 3, 4, 5, 6, testNsec, time.UTC) // 3 Feb 2022 04:05:06.000000007 UTC
}

func TestEndOfWeek(t *testing.T) {
	require.Equal(t, time.Sunday, EndOfWeek(time.Monday))
	require.Equal(t, time.Monday, EndOfWeek(time.Tuesday))
	require.Equal(t, time.Tuesday, EndOfWeek(time.Wednesday))
	require.Equal(t, time.Wednesday, EndOfWeek(time.Thursday))
	require.Equal(t, time.Thursday, EndOfWeek(time.Friday))
	require.Equal(t, time.Friday, EndOfWeek(time.Saturday))
	require.Equal(t, time.Saturday, EndOfWeek(time.Sunday))
}

func TestLocalWeekday(t *testing.T) {
	sow := time.Monday

	require.Equal(t, 0, LocalWeekday(sow, time.Monday))
	require.Equal(t, 1, LocalWeekday(sow, time.Tuesday))
	require.Equal(t, 2, LocalWeekday(sow, time.Wednesday))
	require.Equal(t, 3, LocalWeekday(sow, time.Thursday))
	require.Equal(t, 4, LocalWeekday(sow, time.Friday))
	require.Equal(t, 5, LocalWeekday(sow, time.Saturday))
	require.Equal(t, 6, LocalWeekday(sow, time.Sunday))
}

func TestFromStd(t *testing.T) {
	ut := FromStd(testTime())
	require.Equal(t, testTime(), ut.Time)
}

func TestFromPB(t *testing.T) {
	tt := testTime().Round(time.Nanosecond)
	pb := timestamppb.New(tt)
	ut := FromPB(pb)

	require.Equal(t, tt, ut.Time)
}

func TestNow(t *testing.T) {
	now := time.Now()
	ut := Now()

	require.True(t, now.Before(ut.Time) || now.Equal(ut.Time))
}

func TestDate(t *testing.T) {
	tt := testTime()
	ut := Date(tt.Year(), tt.Month(), tt.Day(), tt.Hour(), tt.Minute(), tt.Second(), tt.Nanosecond(), tt.Location())

	require.Equal(t, tt, ut.Time)
}

func TestUnix(t *testing.T) {
	tt := testTime()
	ut := Unix(tt.Unix(), testNsec).UTC()

	require.Equal(t, tt, ut.Time)
}

func TestTime_Std(t *testing.T) {
	tt := testTime()
	ut := Time{Time: tt}

	require.Equal(t, tt, ut.Std())
}

func TestTime_PB(t *testing.T) {
	tt := testTime()
	ut := Time{Time: tt}

	require.Equal(t, tt, ut.PB().AsTime())
}

func TestTime_MarshalJSON(t *testing.T) {
	forTime := func(t *testing.T, tt time.Time) {
		t.Helper()

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

	t.Run("time", func(t *testing.T) { forTime(t, testTime()) })
	t.Run("zero", func(t *testing.T) { forTime(t, time.Time{}) })
}

func TestTime_UnmarshalJSON(t *testing.T) {
	s := struct {
		T Time `json:"t"`
	}{}

	t.Run(
		"time", func(t *testing.T) {
			tt := testTime().Round(time.Second)
			err := json.Unmarshal([]byte(`{"t":"`+tt.Format(MarshalFormat)+`"}`), &s)
			require.NoError(t, err)
			require.Equal(t, tt, s.T.Time.UTC())
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

func TestTime_MarshalText(t *testing.T) {
	t.Run(
		"time", func(t *testing.T) {
			tt := testTime()
			ut := Time{Time: tt}
			data, err := ut.MarshalText()

			require.NoError(t, err)
			require.Equal(t, tt.Format(MarshalFormat), string(data))
		},
	)

	t.Run(
		"zero", func(t *testing.T) {
			ut := Time{}
			data, err := ut.MarshalText()

			require.NoError(t, err)
			require.Empty(t, data)
		},
	)
}

func TestTime_UnmarshalText(t *testing.T) {
	ut := Time{}

	t.Run(
		"time", func(t *testing.T) {
			tt := testTime().Round(time.Second)
			err := ut.UnmarshalText([]byte(tt.Format(MarshalFormat)))
			require.NoError(t, err)
			require.Equal(t, tt, ut.Time.UTC())
		},
	)

	t.Run(
		"zero", func(t *testing.T) {
			err := ut.UnmarshalText([]byte(""))
			require.NoError(t, err)
			require.True(t, ut.IsZero())
		},
	)
}

func TestTime_MarshalBinary(t *testing.T) {
	t.Run(
		"time", func(t *testing.T) {
			tt := testTime()
			ut := Time{Time: tt}

			data, err := ut.MarshalBinary()
			require.NoError(t, err)

			expected, err := tt.MarshalBinary()
			require.NoError(t, err)

			require.Equal(t, expected, data)
		},
	)

	t.Run(
		"zero", func(t *testing.T) {
			ut := Time{}

			data, err := ut.MarshalBinary()
			require.NoError(t, err)

			expected, err := time.Time{}.MarshalBinary()
			require.NoError(t, err)

			require.Equal(t, expected, data)
		},
	)
}

func TestTime_UnmarshalBinary(t *testing.T) {
	ut := Time{}

	t.Run(
		"time", func(t *testing.T) {
			tt := testTime().Round(time.Second)

			ttBin, err := tt.MarshalBinary()
			require.NoError(t, err)

			err = ut.UnmarshalBinary(ttBin)
			require.NoError(t, err)

			require.Equal(t, tt, ut.Time.UTC())
		},
	)

	t.Run(
		"zero", func(t *testing.T) {
			nowBin, err := time.Time{}.MarshalBinary()
			require.NoError(t, err)

			err = ut.UnmarshalBinary(nowBin)
			require.NoError(t, err)

			require.True(t, ut.IsZero())
		},
	)
}

func TestTime_Value(t *testing.T) {
	tt := testTime()
	ut := Time{Time: tt}

	val, err := ut.Value()
	require.NoError(t, err)
	require.NoError(t, ut.Scan(val))
	require.Equal(t, tt.Local(), ut.Time)
}

func TestTime_Scan(t *testing.T) {
	ut := Time{}

	t.Run(
		"time", func(t *testing.T) {
			tt := testTime()
			err := ut.Scan(tt)
			require.NoError(t, err)
			require.Equal(t, tt.Local(), ut.Time)
		},
	)

	t.Run(
		"zero", func(t *testing.T) {
			err := ut.Scan(nil)
			require.NoError(t, err)
			require.True(t, ut.IsZero())
		},
	)
}

func TestTime_After(t *testing.T) {
	tt := testTime()
	ut := Time{Time: tt}

	require.True(t, ut.After(Time{Time: tt.Add(-time.Second)}))
	require.False(t, ut.After(Time{Time: tt.Add(time.Second)}))
}

func TestTime_Before(t *testing.T) {
	tt := testTime()
	ut := Time{Time: tt}

	require.True(t, ut.Before(Time{Time: tt.Add(time.Second)}))
	require.False(t, ut.Before(Time{Time: tt.Add(-time.Second)}))
}

func TestTime_Equal(t *testing.T) {
	tt := testTime()
	ut := Time{Time: tt}

	require.True(t, ut.Equal(Time{Time: tt}))
	require.False(t, ut.Equal(Time{Time: tt.Add(time.Second)}))
}

func TestTime_AddDate(t *testing.T) {
	tt := testTime()
	ut := Time{Time: tt}

	tt.AddDate(1, 2, 3)
	ut.AddDate(1, 2, 3)

	require.Equal(t, tt, ut.Time)
}

func TestTime_Add(t *testing.T) {
	tt := testTime()
	ut := Time{Time: tt}

	tt.Add(time.Hour)
	ut.Add(time.Hour)

	require.Equal(t, tt, ut.Time)
}

func TestTime_Sub(t *testing.T) {
	tt := testTime()
	ut := Time{Time: tt}
	sub := time.Now()

	tt.Sub(sub)
	ut.Sub(Time{Time: sub})

	require.Equal(t, tt, ut.Time)
}

func TestTime_UTC(t *testing.T) {
	tt := testTime()
	ut := Time{Time: tt}

	tt.UTC()
	ut.UTC()

	require.Equal(t, tt, ut.Time)
}

func TestTime_Local(t *testing.T) {
	tt := testTime()
	ut := Time{Time: tt}

	tt.Local()
	ut.Local()

	require.Equal(t, tt, ut.Time)
}

func TestTime_In(t *testing.T) {
	tt := testTime()
	ut := Time{Time: tt}

	tt.In(time.UTC)
	ut.In(time.UTC)

	require.Equal(t, tt, ut.Time)
}

func TestTime_ZoneBounds(t *testing.T) {
	tt := testTime()
	ut := Time{Time: tt}

	a, b := tt.ZoneBounds()
	c, d := ut.ZoneBounds()

	require.Equal(t, a, c.Time)
	require.Equal(t, b, d.Time)
}

func TestTime_Truncate(t *testing.T) {
	tt := testTime()
	ut := Time{Time: tt}

	tt.Truncate(time.Hour)
	ut.Truncate(time.Hour)

	require.Equal(t, tt, ut.Time)
}

func TestTime_Round(t *testing.T) {
	tt := testTime()
	ut := Time{Time: tt}

	tt.Round(time.Hour)
	ut.Round(time.Hour)

	require.Equal(t, tt, ut.Time)
}

func TestTime_Between(t *testing.T) {
	tt := testTime()
	ut := Time{Time: tt}

	require.True(t, ut.Between(Time{Time: tt.Add(-time.Second)}, Time{Time: tt.Add(time.Second)}))
	require.False(t, ut.Between(Time{Time: tt.Add(time.Second)}, Time{Time: tt.Add(time.Second * 2)}))
}

func TestTime_LocalWeekday(t *testing.T) {
	ut := Time{Time: testTime()}
	require.Equal(t, time.Thursday, ut.Weekday())
	require.Equal(t, 3, ut.LocalWeekday())
}

func TestTime_StartOfDay(t *testing.T) {
	tt := testTime()
	ut := Time{Time: tt}
	startOfDay := ut.StartOfDay()

	require.Equal(
		t,
		time.Date(tt.Year(), tt.Month(), tt.Day(), 0, 0, 0, 0, tt.Location()),
		startOfDay.Time,
	)
}

func TestTime_EndOfDay(t *testing.T) {
	tt := testTime()
	ut := Time{Time: tt}
	endOfDay := ut.EndOfDay()

	require.Equal(
		t,
		time.Date(tt.Year(), tt.Month(), tt.Day(), 23, 59, 59, 999999999, tt.Location()),
		endOfDay.Time,
	)
}

func TestTime_StartOfWeek(t *testing.T) {
	tt := testTime()
	ut := Time{Time: tt}
	startOfWeek := ut.StartOfWeek()

	require.Equal(t, startOfWeek.Time.Weekday(), StartOfWeek)
	require.Equal(t, time.Date(2022, 1, 31, 0, 0, 0, 0, tt.Location()), startOfWeek.Time)
}

func TestTime_EndOfWeek(t *testing.T) {
	tt := testTime()
	ut := Time{Time: tt}
	endOfWeek := ut.EndOfWeek()

	require.Equal(t, endOfWeek.Time.Weekday(), EndOfWeek(StartOfWeek))
	require.Equal(t, time.Date(2022, 2, 6, 23, 59, 59, 999999999, tt.Location()), endOfWeek.Time)
}

func TestTime_StartOfMonth(t *testing.T) {
	tt := testTime()
	ut := Time{Time: tt}
	startOfMonth := ut.StartOfMonth()

	require.Equal(
		t,
		time.Date(tt.Year(), tt.Month(), 1, 0, 0, 0, 0, tt.Location()),
		startOfMonth.Time,
	)
}

func TestTime_EndOfMonth(t *testing.T) {
	tt := testTime()
	ut := Time{Time: tt}
	endOfMonth := ut.EndOfMonth()

	require.Equal(
		t,
		time.Date(tt.Year(), tt.Month()+1, 0, 23, 59, 59, 999999999, tt.Location()),
		endOfMonth.Time,
	)
}
