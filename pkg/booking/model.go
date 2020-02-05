package booking

import "time"

type Booking struct {
	FirstName     string      `json:"firstName"`
	LastName      string      `json:"lastName"`
	Gender        Gender      `json:"gender"`
	Birthday      time.Time   `json:"birthday"`
	LaunchpadID   string      `json:"launchpadId"`
	DestinationID Destination `json:"destinationId"`
	LaunchDate    time.Time   `json:"launchDate"`
}

type Gender string

const (
	Male   Gender = "Male"
	Female        = "Female"
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
