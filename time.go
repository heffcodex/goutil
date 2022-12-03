package goutil

import (
	"encoding/json"
	"errors"
	"strconv"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"
)

const MaxTimestamp = 253402300799

var ErrTimestampOutOfRange = errors.New("timestamp is out of range")

var (
	_ json.Marshaler   = (*Timestamp)(nil)
	_ json.Unmarshaler = (*Timestamp)(nil)
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

func (t *Timestamp) MarshalJSON() ([]byte, error) {
	return []byte(strconv.FormatInt(t.t.Unix(), 10)), nil
}

func (t *Timestamp) UnmarshalJSON(data []byte) error {
	i, err := strconv.ParseInt(string(data), 10, 64)
	if err != nil {
		return err
	} else if i < 0 || i > MaxTimestamp {
		return ErrTimestampOutOfRange
	}

	t.t = time.Unix(i, 0)

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

func (t *Timestamp) StartOfWeek() *Timestamp {
	_t := t.t

	for _t.Weekday() != time.Monday {
		_t = _t.AddDate(0, 0, -1)
	}

	return NewTimestamp(_t).StartOfDay()
}

func (t *Timestamp) EndOfWeek() *Timestamp {
	_t := t.t

	for _t.Weekday() != time.Sunday {
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
