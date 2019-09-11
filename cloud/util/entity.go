package util

import (
	"strings"
	"time"
)

type Time time.Time

func (t Time) Format() string {
	return time.Time(t).Format("2006-01-02 15:04:05")
}

func (t Time) MarshalText() ([]byte, error) {
	return []byte(t.Format()), nil
}

func (t Time) MarshalJSON() ([]byte, error) {
	return []byte(`"` + t.Format() + `"`), nil
}

func (t *Time) UnmarshalJSON(b []byte) error {
	v, err := time.ParseInLocation("2006-01-02 15:04:05", strings.Trim(string(b), "\""), time.Local)
	if err != nil {
		return err
	}
	*t = Time(v)
	return nil
}

func ConutRate(a int, b int) int {
	if a > b {
		return 1
	}
	return int((float64(a) / float64(b)) * 100)
}
