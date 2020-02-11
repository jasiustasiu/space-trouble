package booking

import "space-trouble/internal/date"

type Booking struct {
	ID            *int64      `json:"id" db:"id"`
	FirstName     string      `json:"firstName" db:"first_name"`
	LastName      string      `json:"lastName" db:"last_name"`
	Gender        Gender      `json:"gender" db:"gender"`
	Birthday      date.Date   `json:"birthday" db:"birthday"`
	LaunchpadID   string      `json:"launchpadId" db:"launchpad_id"`
	DestinationID Destination `json:"destinationId" db:"destination_id"`
	LaunchDate    date.Date   `json:"launchDate" db:"launch_date"`
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

type AvailabilityResponse struct {
	Available bool
	Err       error
}
