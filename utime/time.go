package utime

import (
	"encoding"
	"encoding/json"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"
)

var (
	MarshalFormat = time.RFC3339
	StartOfWeek   = time.Monday
)

func EndOfWeek(startOfWeek time.Weekday) time.Weekday {
	return (startOfWeek + 6) % 7
}

func LocalWeekday(startOfWeek time.Weekday, wd time.Weekday) int {
	return int(7+wd-startOfWeek) % 7
}

var (
	_ json.Marshaler             = (*Time)(nil)
	_ json.Unmarshaler           = (*Time)(nil)
	_ encoding.TextMarshaler     = (*Time)(nil)
	_ encoding.TextUnmarshaler   = (*Time)(nil)
	_ encoding.BinaryMarshaler   = (*Time)(nil)
	_ encoding.BinaryUnmarshaler = (*Time)(nil)
)

type Time struct{ time.Time }

// constructors:

func FromStdTime(t time.Time) Time {
	return Time{Time: t}
}

func FromPB(t *timestamppb.Timestamp) Time {
	return Time{Time: t.AsTime()}
}

func Now() Time {
	return Time{Time: time.Now()}
}

func Date(year int, month time.Month, day, hour, min, sec, nsec int, loc *time.Location) Time {
	return Time{Time: time.Date(year, month, day, hour, min, sec, nsec, loc)}
}

// converters:

func (t Time) StdTime() time.Time {
	return t.Time
}

func (t Time) PB() *timestamppb.Timestamp {
	return timestamppb.New(t.Time)
}

// (un)marshalers:

func (t Time) MarshalJSON() ([]byte, error) {
	if t.IsZero() {
		return []byte(`null`), nil
	}

	return []byte(`"` + t.Format(MarshalFormat) + `"`), nil
}

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

func (t Time) MarshalText() (text []byte, err error) {
	if t.IsZero() {
		return nil, nil
	}

	return []byte(t.Format(MarshalFormat)), nil
}

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

func (t Time) MarshalBinary() (data []byte, err error) {
	return t.Time.MarshalBinary()
}

func (t *Time) UnmarshalBinary(data []byte) error {
	_t := time.Time{}
	if err := _t.UnmarshalBinary(data); err != nil {
		return err
	}

	t.Time = _t.Local()
	return nil
}

// time.Time wrappers:

func (t Time) After(u Time) bool {
	return t.Time.After(u.Time)
}

func (t Time) Before(u Time) bool {
	return t.Time.Before(u.Time)
}

func (t Time) Equal(u Time) bool {
	return t.Time.Equal(u.Time)
}

func (t Time) AddDate(years, months, days int) Time {
	return Time{Time: t.Time.AddDate(years, months, days)}
}

func (t Time) Add(d time.Duration) Time {
	return Time{Time: t.Time.Add(d)}
}

func (t Time) Sub(u Time) time.Duration {
	return t.Time.Sub(u.Time)
}

func (t Time) UTC() Time {
	return Time{Time: t.Time.UTC()}
}

func (t Time) Local() Time {
	return Time{Time: t.Time.Local()}
}

func (t Time) In(loc *time.Location) Time {
	return Time{Time: t.Time.In(loc)}
}

func (t Time) ZoneBounds() (start, end Time) {
	_start, _end := t.Time.ZoneBounds()
	return Time{Time: _start}, Time{Time: _end}
}

func (t Time) Truncate(d time.Duration) Time {
	return Time{Time: t.Time.Truncate(d)}
}

func (t Time) Round(d time.Duration) Time {
	return Time{Time: t.Time.Round(d)}
}

// utility functions:

func (t Time) Between(start, end Time) bool {
	return (t.After(start) || t.Equal(start)) && (t.Before(end) || t.Equal(end))
}

func (t Time) StartOfDay() Time {
	year, month, day := t.Date()
	return FromStdTime(time.Date(year, month, day, 0, 0, 0, 0, t.Location()))
}

func (t Time) EndOfDay() Time {
	year, month, day := t.Date()
	return FromStdTime(time.Date(year, month, day, 23, 59, 59, int(time.Second-time.Nanosecond), t.Location()))
}

func (t Time) LocalWeekday() int {
	return LocalWeekday(StartOfWeek, t.Weekday())
}

func (t Time) StartOfWeek() Time {
	return t.AddDate(0, 0, -t.LocalWeekday()).StartOfDay()
}

func (t Time) EndOfWeek() Time {
	return t.AddDate(0, 0, 6-t.LocalWeekday()).EndOfDay()
}

func (t Time) StartOfMonth() Time {
	year, month, _ := t.Date()
	return Date(year, month, 1, 0, 0, 0, 0, t.Location())
}

func (t Time) EndOfMonth() Time {
	year, month, _ := t.AddDate(0, 1, 0).Date()
	return Date(year, month, 1, 0, 0, 0, -1, t.Location())
}

func (t *Time) RuMonthName() string {
	var longMonthNamesRuRU = map[string]string{
		"January":   "Январь",
		"February":  "Февраль",
		"March":     "Март",
		"April":     "Апрель",
		"May":       "Май",
		"June":      "Июнь",
		"July":      "Июль",
		"August":    "Август",
		"September": "Сентябрь",
		"October":   "Октябрь",
		"November":  "Ноябрь",
		"December":  "Декабрь",
	}

	return longMonthNamesRuRU[t.Format("January")]
}

func (t *Time) RuMonthNamePrepositional() string {
	var longMonthNamesRuRU = map[string]string{
		"January":   "Январе",
		"February":  "Феврале",
		"March":     "Марте",
		"April":     "Апреле",
		"May":       "Мае",
		"June":      "Июне",
		"July":      "Июле",
		"August":    "Августе",
		"September": "Сентябре",
		"October":   "Октябре",
		"November":  "Ноябре",
		"December":  "Декабре",
	}

	return longMonthNamesRuRU[t.Format("January")]
}
