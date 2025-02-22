package s_utils

import (
	"fmt"
	"time"
)

// TimeToString Преобразует время в формат времени принятый в ksi
func TimeToString(t *time.Time, local ...bool) (str string) {
	if t == nil {
		str = ""
	} else {
		if len(local) == 1 {
			if local[0] {
				str = t.Local().Format("2006-01-02T15:04:05.000000Z07:00")
			} else {
				str = t.UTC().Format("2006-01-02T15:04:05.000000Z07:00")
			}
		} else {
			str = t.UTC().Format("2006-01-02T15:04:05.000000Z07:00")
		}
	}
	return
}

// TimeToStringF Преобразует время в формат времени принятый в ksi
func TimeToStringF(t *time.Time, format string, local ...bool) (str string) {
	if t == nil {
		str = ""
	} else {
		if len(local) == 1 {
			if local[0] {
				str = t.Local().Format(format)
			} else {
				str = t.UTC().Format(format)
			}
		} else {
			str = t.UTC().Format(format)
		}
	}
	return
}

// TimeParse Получение времени из строки без указания формата
func TimeParse(tStr string) (t time.Time, err error) {
	layouts := []string{
		"2006-01-02T15:04:05Z07:00",
		"2006-01-02T15:04:05.999999999Z07:00",
		"2006-01-02 15:04:05",
		"01/02 03:04:05PM '06 -0700",
		"Mon Jan _2 15:04:05 2006",
		"Mon Jan _2 15:04:05 MST 2006",
		"Mon Jan 02 15:04:05 -0700 2006",
		"02 Jan 06 15:04 MST",
		"02 Jan 06 15:04 -0700",
		"Monday, 02-Jan-06 15:04:05 MST",
		"Mon, 02 Jan 2006 15:04:05 MST",
		"Mon, 02 Jan 2006 15:04:05 -0700",
		"Jan _2 15:04:05",
		"Jan _2 15:04:05.000",
		"Jan _2 15:04:05.000000",
		"Jan _2 15:04:05.000000000",
		"2006-01-02",
		"3:04PM",
		"15:04:05",
	}

	for _, v := range layouts {
		t, err = time.Parse(v, tStr)
		if err != nil {
			continue
		} else {
			return
		}
	}
	return t, fmt.Errorf(`cannot parse time from given string: "%v"`, tStr)
}
