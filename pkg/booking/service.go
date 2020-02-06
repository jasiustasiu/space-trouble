package booking

import (
	"fmt"
	"space-trouble/internal/httpError"
	"space-trouble/pkg/spacex"
	"time"
)

var weekdaysToDestinations = map[time.Weekday]Destination{
	time.Sunday:    Mars,
	time.Monday:    Moon,
	time.Tuesday:   Pluto,
	time.Wednesday: AsteroidBelt,
	time.Thursday:  Europa,
	time.Friday:    Titan,
	time.Saturday:  Ganymede,
}

func NewBookingService(repository Repository, spacexAPI *spacex.API) *Service {
	return &Service{
		repository: repository,
		spacexAPI:  spacexAPI,
	}
}

type Service struct {
	repository Repository
	spacexAPI  *spacex.API
}

func (s *Service) CreateBooking(booking Booking) error {
	todayDestination := weekdaysToDestinations[booking.LaunchDate.Weekday()]
	if booking.DestinationID != todayDestination {
		return httpError.NewHTTPError(400, fmt.Sprintf("We departure to %v only for given day", todayDestination))
	}
	launchpad, err := s.spacexAPI.GetLaunchPad(booking.LaunchpadID)
	if err != nil {
		return httpError.NewHTTPError(400, fmt.Sprintf("Launch pad with id %v does not exist", booking.LaunchpadID))
	}
	if launchpad.Status != "active" {
		return httpError.NewHTTPError(400, fmt.Sprintf("Launch pad with id %v is not active", booking.LaunchpadID))
	}

	launches, err := s.spacexAPI.ListUpcomingLaunches(booking.LaunchpadID)
	if err != nil {
		return err
	}
	launchDateStr := booking.LaunchDate.Format(dateFormat)
	for _, launch := range launches {
		if launch.LaunchDateLocal.Format(dateFormat) == launchDateStr {
			return httpError.NewHTTPError(400, "Given launch pad is already booked for selected day")
		}
	}
	return nil
}

func (s *Service) GetBookings() ([]Booking, error) {
	return s.repository.GetAll()
}
