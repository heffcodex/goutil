package goutil

import (
	"errors"
	"strconv"
	"time"
)

const MaxTimestamp = 253402300799

type Timestamp struct {
	T time.Time
}

func (t *Timestamp) MarshalJSON() ([]byte, error) {
	return []byte(strconv.FormatInt(t.T.Unix(), 10)), nil
}

func (t *Timestamp) UnmarshalJSON(data []byte) error {
	i, err := strconv.ParseInt(string(data), 10, 64)
	if err != nil {
		return err
	} else if i < 0 || i > MaxTimestamp {
		return errors.New("invalid timestamp")
	}

	t.T = time.Unix(i, 0)

	return nil
}

func StartOfDay(t time.Time) time.Time {
	year, month, day := t.Date()
	return time.Date(year, month, day, 0, 0, 0, 0, t.Location())
}

func EndOfDay(t time.Time) time.Time {
	year, month, day := t.Date()
	return time.Date(year, month, day, 23, 59, 59, int(time.Second-time.Nanosecond), t.Location())
}

func StartOfWeek(t time.Time) time.Time {
	for t.Weekday() != time.Monday {
		t = t.AddDate(0, 0, -1)
	}

	return StartOfDay(t)
}

func EndOfWeek(t time.Time) time.Time {
	for t.Weekday() != time.Sunday {
		t = t.AddDate(0, 0, 1)
	}

	return EndOfDay(t)
}

func StartOfMonth(t time.Time) time.Time {
	year, month, _ := t.Date()
	return time.Date(year, month, 1, 0, 0, 0, 0, t.Location())
}

func EndOfMonth(t time.Time) time.Time {
	return StartOfMonth(t.AddDate(0, 1, 0)).Add(-time.Nanosecond)
}

func TimeRef(t time.Time) *time.Time {
	return &t
}

func TimeUnixOrNil(t *time.Time) *int64 {
	if t == nil || t.IsZero() {
		return nil
	}

	unix := t.Unix()

	return &unix
}

func TimeRefOrNil(unix *int64) *time.Time {
	if unix == nil || *unix < 1 {
		return nil
	}

	return TimeRef(time.Unix(*unix, 0))
}

func RuMonthName(t time.Time) string {
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

func RuMonthNamePrepositional(t time.Time) string {
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
