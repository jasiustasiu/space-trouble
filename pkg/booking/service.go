package booking

import "space-trouble/pkg/spacex"

func NewBookingService(spacexAPI *spacex.API) *Service {
	return &Service{
		spacexAPI: spacexAPI,
	}
}

type Service struct {
	spacexAPI *spacex.API
}

func (s * Service) CreateBooking(booking Booking) error {
	s.spacexAPI.ListUpcomingLaunches()
	return nil
}
