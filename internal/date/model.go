package date

import (
	"encoding/json"
	"strings"
	"time"
)

const (
	Format = "2006-01-02"
)

type Date time.Time

func (d *Date) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"")
	t, err := time.Parse(Format, s)
	if err != nil {
		return err
	}
	*d = Date(t)
	return nil
}

func (d *Date) MarshalJSON() ([]byte, error) {
	return json.Marshal(d)
}

func (d Date) Format(s string) string {
	t := time.Time(d)
	return t.Format(s)
}

func (d Date) Weekday() time.Weekday {
	t := time.Time(d)
	return t.Weekday()
}