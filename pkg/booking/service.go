package booking

import (
	"fmt"
	"net/http"
	"space-trouble/internal/date"
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

type AvailabilityService interface {
	IsLaunchpadAvailable(launchpadID string, launchDate date.Date) (bool, error)
}

type Service interface {
	CreateBooking(booking Booking) error
 	GetBookings() ([]Booking, error)
}

func NewBookingService(repository Repository, spacexService spacex.Service) Service {
	return &service{
		repository: repository,
		spacexService:  spacexService,
	}
}

type service struct {
	repository Repository
	spacexService  spacex.Service
}

func (s *service) CreateBooking(booking Booking) error {
	todayDestination := weekdaysToDestinations[booking.LaunchDate.Weekday()]
	if booking.DestinationID != todayDestination {
		return httpError.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("We departure to %v only for given day", todayDestination))
	}
	//TODO change to channels, refactor a bit
	available1, err := s.spacexService.IsLaunchpadAvailable(booking.LaunchpadID, booking.LaunchDate)
	available2, err := s.IsLaunchpadAvailable(booking.LaunchpadID, booking.LaunchDate)
	if available1 == false || available2 == false {
		return httpError.NewHTTPError(http.StatusConflict, "There is already a flight booked for given day")
	}
	return err
}

func (s *service) GetBookings() ([]Booking, error) {
	return s.repository.GetAll()
}

func (s *service) IsLaunchpadAvailable(launchpadID string, launchDate date.Date) (bool, error) {
	_, ok := s.repository.Get(launchpadID, launchDate)
	return ok, nil
}