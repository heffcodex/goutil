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

func EndOfWeek() time.Weekday { return (StartOfWeek + 6) % 7 }

var (
	_ json.Marshaler             = (*Time)(nil)
	_ json.Unmarshaler           = (*Time)(nil)
	_ encoding.TextMarshaler     = (*Time)(nil)
	_ encoding.TextUnmarshaler   = (*Time)(nil)
	_ encoding.BinaryMarshaler   = (*Time)(nil)
	_ encoding.BinaryUnmarshaler = (*Time)(nil)
)

type Time struct{ time.Time }

func FromStdTime(t time.Time) Time {
	return Time{Time: t}
}

func FromPB(t *timestamppb.Timestamp) Time {
	return Time{Time: t.AsTime()}
}

func (t Time) StdTime() time.Time {
	return t.Time
}

func (t Time) PB() *timestamppb.Timestamp {
	return timestamppb.New(t.Time)
}

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

func (t Time) StartOfDay() Time {
	year, month, day := t.Date()
	return FromStdTime(time.Date(year, month, day, 0, 0, 0, 0, t.Location()))
}

func (t Time) EndOfDay() Time {
	year, month, day := t.Date()
	return FromStdTime(time.Date(year, month, day, 23, 59, 59, int(time.Second-time.Nanosecond), t.Location()))
}

func (t Time) StartOfWeek() Time { // TODO: replace with faster implementation
	_t := t.Time

	for _t.Weekday() != StartOfWeek {
		_t = _t.AddDate(0, 0, -1)
	}

	return FromStdTime(_t).StartOfDay()
}

func (t Time) EndOfWeek() Time { // TODO: replace with faster implementation
	_t := t.Time
	eow := EndOfWeek()

	for _t.Weekday() != eow {
		_t = _t.AddDate(0, 0, 1)
	}

	return FromStdTime(_t).EndOfDay()
}

func (t Time) StartOfMonth() Time {
	year, month, _ := t.Date()
	return FromStdTime(time.Date(year, month, 1, 0, 0, 0, 0, t.Location()))
}

func (t Time) EndOfMonth() Time {
	year, month, _ := t.AddDate(0, 1, 0).Date()
	return FromStdTime(time.Date(year, month, 1, 0, 0, 0, -1, t.Location()))
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
