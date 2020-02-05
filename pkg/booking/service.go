package booking

import (
	"fmt"
	"space-trouble/internal/httpError"
	"space-trouble/pkg/spacex"
	"time"
)

const (
	dateFormat = "2006-01-02"
)

var weekdaysToDestinations = map[time.Weekday]Destination{
	time.Sunday: Mars,
	time.Monday: Moon,
	time.Tuesday: Pluto,
	time.Wednesday: AsteroidBelt,
	time.Thursday: Europa,
	time.Friday: Titan,
	time.Saturday: Ganymede,
}

func NewBookingService(spacexAPI *spacex.API) *Service {
	return &Service{
		spacexAPI: spacexAPI,
	}
}

type Service struct {
	spacexAPI *spacex.API
}

func (s *Service) CreateBooking(booking Booking) error {
	launches, err := s.spacexAPI.ListUpcomingLaunches(booking.LaunchpadID)
	if err != nil {
		return err
	}
	if weekdaysToDestinations[booking.LaunchDate.Weekday()] != booking.DestinationID {
		return httpError.NewHTTPError(400, fmt.Sprintf("There is not any flight to %v for selected day", booking.DestinationID))
	}
	launchDateStr := booking.LaunchDate.Format(dateFormat)
	for _, launch := range launches {
		if launch.LaunchDateLocal.Format(dateFormat) == launchDateStr {
			return httpError.NewHTTPError(400, "Given launch pad is already booked for selected day")
		}
	}
	return nil
}
