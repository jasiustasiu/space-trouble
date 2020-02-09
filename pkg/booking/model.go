package booking

import "space-trouble/internal/date"

type Booking struct {
	FirstName     string      `json:"firstName"`
	LastName      string      `json:"lastName"`
	Gender        Gender      `json:"gender"`
	Birthday      date.Date   `json:"birthday"`
	LaunchpadID   string      `json:"launchpadId"`
	DestinationID Destination `json:"destinationId"`
	LaunchDate    date.Date   `json:"launchDate"`
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
