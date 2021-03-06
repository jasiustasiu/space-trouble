package booking

import (
	"fmt"
	"net/http"
	"space-trouble/internal/date"
	"space-trouble/internal/httpError"
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
	DeleteBooking(id int64) error
}

func NewBookingService(repository Repository, availabilityService ...AvailabilityService) Service {
	return &service{
		repository:           repository,
		availabilityServices: availabilityService,
	}
}

type service struct {
	repository           Repository
	availabilityServices []AvailabilityService
}

func (s *service) CreateBooking(booking Booking) error {
	todayDestination := weekdaysToDestinations[booking.LaunchDate.Weekday()]
	if booking.DestinationID != todayDestination {
		return httpError.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("We departure to %v only for given day", todayDestination))
	}
	out := make(chan AvailabilityResponse, len(s.availabilityServices))
	for _, availabilityService := range s.availabilityServices {
		go isLaunchpadAvailable(availabilityService, booking.LaunchpadID, booking.LaunchDate, out)
	}
	for range s.availabilityServices {
		response := <-out
		if response.Err != nil {
			return response.Err
		}
		if !response.Available {
			return httpError.NewHTTPError(http.StatusConflict, "There is already a flight booked for given day")
		}
	}
	close(out)
	return s.repository.Save(booking)
}

func isLaunchpadAvailable(svc AvailabilityService, launchpadID string, launchDate date.Date, out chan<- AvailabilityResponse) {
	available, err := svc.IsLaunchpadAvailable(launchpadID, launchDate)
	out <- AvailabilityResponse{Available: available, Err: err,}
}

func (s *service) GetBookings() ([]Booking, error) {
	return s.repository.GetAll()
}

func (s *service) DeleteBooking(id int64) error {
	return s.repository.Delete(id)
}
