package goutil

import (
	"encoding"
	"encoding/json"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"
)

const (
	TimestampMarshalFormat = time.RFC3339
	StartOfWeek            = time.Monday
	EndOfWeek              = (StartOfWeek + 6) % 7
)

var (
	_ json.Marshaler             = (*Timestamp)(nil)
	_ json.Unmarshaler           = (*Timestamp)(nil)
	_ encoding.TextMarshaler     = (*Timestamp)(nil)
	_ encoding.TextUnmarshaler   = (*Timestamp)(nil)
	_ encoding.BinaryMarshaler   = (*Timestamp)(nil)
	_ encoding.BinaryUnmarshaler = (*Timestamp)(nil)
)

type Timestamp struct {
	t time.Time
}

func NewTimestamp(t time.Time) *Timestamp {
	return &Timestamp{t: t}
}

func NewTimestampFromPB(t *timestamppb.Timestamp) *Timestamp {
	return &Timestamp{t: t.AsTime()}
}

func (t *Timestamp) Time() time.Time {
	return t.t
}

func (t *Timestamp) PB() *timestamppb.Timestamp {
	return timestamppb.New(t.t)
}

func (t Timestamp) MarshalJSON() ([]byte, error) {
	if t.t.IsZero() {
		return []byte(`null`), nil
	}

	return []byte(`"` + t.t.Format(TimestampMarshalFormat) + `"`), nil
}

func (t *Timestamp) UnmarshalJSON(data []byte) error {
	strData := string(data)

	if strData == `""` || strData == `null` {
		t.t = time.Time{}
		return nil
	}

	_t, err := time.Parse(`"`+TimestampMarshalFormat+`"`, strData)
	if err != nil {
		return err
	}

	t.t = _t.Local()
	return nil
}

func (t Timestamp) MarshalText() (text []byte, err error) {
	if t.t.IsZero() {
		return nil, nil
	}

	return []byte(t.t.Format(TimestampMarshalFormat)), nil
}

func (t *Timestamp) UnmarshalText(text []byte) error {
	if len(text) == 0 {
		t.t = time.Time{}
		return nil
	}

	_t, err := time.Parse(TimestampMarshalFormat, string(text))
	if err != nil {
		return err
	}

	t.t = _t.Local()
	return nil
}

func (t Timestamp) MarshalBinary() (data []byte, err error) {
	return t.t.MarshalBinary()
}

func (t *Timestamp) UnmarshalBinary(data []byte) error {
	_t := time.Time{}
	if err := _t.UnmarshalBinary(data); err != nil {
		return err
	}

	t.t = _t.Local()
	return nil
}

func (t *Timestamp) StartOfDay() *Timestamp {
	year, month, day := t.t.Date()
	return NewTimestamp(time.Date(year, month, day, 0, 0, 0, 0, t.t.Location()))
}

func (t *Timestamp) EndOfDay() *Timestamp {
	year, month, day := t.t.Date()
	return NewTimestamp(time.Date(year, month, day, 23, 59, 59, int(time.Second-time.Nanosecond), t.t.Location()))
}

func (t *Timestamp) StartOfWeek() *Timestamp { // TODO: replace with faster implementation
	_t := t.t

	for _t.Weekday() != StartOfWeek {
		_t = _t.AddDate(0, 0, -1)
	}

	return NewTimestamp(_t).StartOfDay()
}

func (t *Timestamp) EndOfWeek() *Timestamp { // TODO: replace with faster implementation
	_t := t.t

	for _t.Weekday() != EndOfWeek {
		_t = _t.AddDate(0, 0, 1)
	}

	return NewTimestamp(_t).EndOfDay()
}

func (t *Timestamp) StartOfMonth() *Timestamp {
	year, month, _ := t.t.Date()
	return NewTimestamp(time.Date(year, month, 1, 0, 0, 0, 0, t.t.Location()))
}

func (t *Timestamp) EndOfMonth() *Timestamp {
	year, month, _ := t.t.AddDate(0, 1, 0).Date()
	return NewTimestamp(time.Date(year, month, 1, 0, 0, 0, -1, t.t.Location()))
}

func (t *Timestamp) RuMonthName() string {
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

	return longMonthNamesRuRU[t.t.Format("January")]
}

func (t *Timestamp) RuMonthNamePrepositional() string {
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

	return longMonthNamesRuRU[t.t.Format("January")]
}
