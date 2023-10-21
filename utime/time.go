package utime

import (
	"database/sql"
	"database/sql/driver"
	"encoding"
	"encoding/json"
	"fmt"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"
)

var (
	// MarshalFormat to use for (un-)marshaling time.Time values.
	MarshalFormat = time.RFC3339

	// SQLFormat to use for (un-)marshaling time.Time values with sql drivers.
	SQLFormat = time.RFC3339Nano

	// StartOfWeek is the beginning of the week.
	// This overrides the default time.Sunday and changes behavior of weekday-related functions in this package.
	StartOfWeek = time.Monday
)

// EndOfWeek returns the circular-shifted time.Weekday according to the defined StartOfWeek global variable.
func EndOfWeek(startOfWeek time.Weekday) time.Weekday {
	return (startOfWeek + 6) % 7
}

// LocalWeekday returns the circular-shifted weekday number [0; 6] according to the given `startOfWeek`.
func LocalWeekday(startOfWeek, day time.Weekday) int {
	return int(7+day-startOfWeek) % 7
}

type iTime interface {
	json.Marshaler
	json.Unmarshaler
	encoding.TextMarshaler
	encoding.TextUnmarshaler
	encoding.BinaryMarshaler
	encoding.BinaryUnmarshaler
	driver.Valuer
	sql.Scanner

	Std() time.Time
	PB() *timestamppb.Timestamp
}

var _ iTime = (*Time)(nil)

// Time is a wrapper around time.Time with some useful methods.
type Time struct{ time.Time }

// constructors:

// FromStd creates a Time from a standard time.Time.
func FromStd(t time.Time) Time {
	return Time{Time: t}
}

// FromPB creates a Time from a protobuf timestamp.
func FromPB(t *timestamppb.Timestamp) Time {
	return FromStd(t.AsTime())
}

// Now returns the current time.
func Now() Time {
	return FromStd(time.Now())
}

// Date constructs a Time from the given date parts.
func Date(year int, month time.Month, day, hour, min, sec, nsec int, loc *time.Location) Time {
	return FromStd(time.Date(year, month, day, hour, min, sec, nsec, loc))
}

// Unix constructs a Time from the given unix epoch time.
func Unix(sec, nsec int64) Time {
	return FromStd(time.Unix(sec, nsec))
}

// converters:

// Std converts a Time to a standard time.Time.
func (t Time) Std() time.Time {
	return t.Time
}

// PB converts a Time to a protobuf timestamp.
func (t Time) PB() *timestamppb.Timestamp {
	return timestamppb.New(t.Time)
}

// (un)marshalers:

// MarshalJSON implements the json.Marshaler interface for the defined MarshalFormat global variable.
func (t Time) MarshalJSON() ([]byte, error) {
	if t.IsZero() {
		return []byte(`null`), nil
	}

	return []byte(`"` + t.Format(MarshalFormat) + `"`), nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for the defined MarshalFormat global variable.
// Returns the local time.
func (t *Time) UnmarshalJSON(data []byte) error {
	strData := string(data)

	if strData == `""` || strData == `null` {
		t.Time = time.Time{}
		return nil
	}

	_t, err := time.Parse(`"`+MarshalFormat+`"`, strData)
	if err != nil {
		return err
	}

	t.Time = _t.Local()

	return nil
}

// MarshalText implements the encoding.TextMarshaler interface for the defined MarshalFormat global variable.
func (t Time) MarshalText() (text []byte, err error) {
	if t.IsZero() {
		return nil, nil
	}

	return []byte(t.Format(MarshalFormat)), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface for the defined MarshalFormat global variable.
// Returns the local time.
func (t *Time) UnmarshalText(text []byte) error {
	if len(text) == 0 {
		t.Time = time.Time{}
		return nil
	}

	_t, err := time.Parse(MarshalFormat, string(text))
	if err != nil {
		return err
	}

	t.Time = _t.Local()

	return nil
}

// MarshalBinary implements the encoding.BinaryMarshaler interface.
func (t Time) MarshalBinary() (data []byte, err error) {
	return t.Time.MarshalBinary()
}

// UnmarshalBinary implements the encoding.BinaryUnmarshaler interface.
// Returns the local time.
func (t *Time) UnmarshalBinary(data []byte) error {
	_t := time.Time{}
	if err := _t.UnmarshalBinary(data); err != nil {
		return err
	}

	t.Time = _t.Local()

	return nil
}

// Value implements the driver.Valuer interface for the defined SQLFormat global variable.
func (t Time) Value() (driver.Value, error) {
	if t.IsZero() {
		return nil, nil
	}

	return t.Format(SQLFormat), nil
}

// Scan implements the sql.Scanner interface for the defined SQLFormat global variable.
// Scans the local time.
func (t *Time) Scan(src any) error {
	var (
		stdT time.Time
		err  error
	)

	switch srcT := src.(type) {
	case time.Time:
		stdT = srcT
	case string:
		stdT, err = time.ParseInLocation(SQLFormat, srcT, time.UTC)
	case []byte:
		stdT, err = time.ParseInLocation(SQLFormat, string(srcT), time.UTC)
	case nil:
		// do nothing
	default:
		err = fmt.Errorf("unsupported data type: %T", src)
	}

	if err != nil {
		return err
	}

	if stdT.IsZero() {
		t.Time = time.Time{}
	} else {
		t.Time = stdT.Local()
	}

	return nil
}

// time.Time wrappers:

// After wraps the standard function.
func (t Time) After(u Time) bool {
	return t.Time.After(u.Time)
}

// Before wraps the standard function.
func (t Time) Before(u Time) bool {
	return t.Time.Before(u.Time)
}

// Equal wraps the standard function.
func (t Time) Equal(u Time) bool {
	return t.Time.Equal(u.Time)
}

// Compare wraps the standard function.
func (t Time) Compare(u Time) int {
	return t.Time.Compare(u.Time)
}

// AddDate wraps the standard function.
func (t Time) AddDate(years, months, days int) Time {
	return FromStd(t.Time.AddDate(years, months, days))
}

// Add wraps the standard function.
func (t Time) Add(d time.Duration) Time {
	return FromStd(t.Time.Add(d))
}

// Sub wraps the standard function.
func (t Time) Sub(u Time) time.Duration {
	return t.Time.Sub(u.Time)
}

// UTC wraps the standard function.
func (t Time) UTC() Time {
	return FromStd(t.Time.UTC())
}

// Local wraps the standard function.
func (t Time) Local() Time {
	return FromStd(t.Time.Local())
}

// In wraps the standard function.
func (t Time) In(loc *time.Location) Time {
	return FromStd(t.Time.In(loc))
}

// ZoneBounds wraps the standard function.
func (t Time) ZoneBounds() (start, end Time) {
	_start, _end := t.Time.ZoneBounds()
	return FromStd(_start), FromStd(_end)
}

// Truncate wraps the standard function.
func (t Time) Truncate(d time.Duration) Time {
	return FromStd(t.Time.Truncate(d))
}

// Round wraps the standard function.
func (t Time) Round(d time.Duration) Time {
	return FromStd(t.Time.Round(d))
}

// utility functions:

// Between returns true if time is between the given bounds.
func (t Time) Between(start, end Time) bool {
	return (t.After(start) || t.Equal(start)) && (t.Before(end) || t.Equal(end))
}

// StartOfDay returns the start of the day.
func (t Time) StartOfDay() Time {
	year, month, day := t.Date()
	return FromStd(time.Date(year, month, day, 0, 0, 0, 0, t.Location()))
}

// EndOfDay returns the end of the day.
func (t Time) EndOfDay() Time {
	year, month, day := t.Date()
	return FromStd(time.Date(year, month, day, 23, 59, 59, int(time.Second-time.Nanosecond), t.Location()))
}

// LocalWeekday returns the circular-shifted weekday number [0; 6] according to the defined StartOfWeek global variable.
func (t Time) LocalWeekday() int {
	return LocalWeekday(StartOfWeek, t.Weekday())
}

// StartOfWeek returns the start of the week according to the defined StartOfWeek global variable.
func (t Time) StartOfWeek() Time {
	return t.AddDate(0, 0, -t.LocalWeekday()).StartOfDay()
}

// EndOfWeek returns the end of the week according to the defined StartOfWeek global variable.
func (t Time) EndOfWeek() Time {
	return t.AddDate(0, 0, 6-t.LocalWeekday()).EndOfDay()
}

// StartOfMonth returns the start of the month.
func (t Time) StartOfMonth() Time {
	year, month, _ := t.Date()
	return Date(year, month, 1, 0, 0, 0, 0, t.Location())
}

// EndOfMonth returns the end of the month.
func (t Time) EndOfMonth() Time {
	year, month, _ := t.Date()
	return Date(year, month+1, 1, 0, 0, 0, -1, t.Location())
}

// RuMonthName returns the name of month in Russian.
func (t *Time) RuMonthName() string {
	longMonthNamesRuRU := map[time.Month]string{
		time.January:   "Январь",
		time.February:  "Февраль",
		time.March:     "Март",
		time.April:     "Апрель",
		time.May:       "Май",
		time.June:      "Июнь",
		time.July:      "Июль",
		time.August:    "Август",
		time.September: "Сентябрь",
		time.October:   "Октябрь",
		time.November:  "Ноябрь",
		time.December:  "Декабрь",
	}

	return longMonthNamesRuRU[t.Month()]
}

// RuMonthNamePrepositional returns the name of month in Russian in prepositional case.
func (t *Time) RuMonthNamePrepositional() string {
	longMonthNamesRuRU := map[time.Month]string{
		time.January:   "Январе",
		time.February:  "Феврале",
		time.March:     "Марте",
		time.April:     "Апреле",
		time.May:       "Мае",
		time.June:      "Июне",
		time.July:      "Июле",
		time.August:    "Августе",
		time.September: "Сентябре",
		time.October:   "Октябре",
		time.November:  "Ноябре",
		time.December:  "Декабре",
	}

	return longMonthNamesRuRU[t.Month()]
}
