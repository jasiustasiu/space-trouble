package booking

import (
	"encoding/json"
	"strings"
	"time"
)

const (
	dateFormat = "2006-01-02"
)

type Booking struct {
	FirstName     string      `json:"firstName"`
	LastName      string      `json:"lastName"`
	Gender        Gender      `json:"gender"`
	Birthday      Date        `json:"birthday"`
	LaunchpadID   string      `json:"launchpadId"`
	DestinationID Destination `json:"destinationId"`
	LaunchDate    Date        `json:"launchDate"`
}

type Gender string

const (
	Male   Gender = "M"
	Female        = "F"
)

type Destination string

const (
	Mars         Destination = "mars"
	Moon                     = "moon"
	Pluto                    = "pluto"
	AsteroidBelt             = "asteroid_belt"
	Europa                   = "europa"
	Titan                    = "titan"
	Ganymede                 = "ganymede"
)

type Date time.Time

func (d *Date) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"")
	t, err := time.Parse(dateFormat, s)
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
