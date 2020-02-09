package spacex

import (
	"fmt"
	"space-trouble/internal/date"
	"space-trouble/internal/httpError"
	"space-trouble/pkg/booking"
)

type Service interface {
	IsLaunchpadAvailable(launchpadID string, launchDate date.Date, out chan<- booking.AvailabilityResponse)
}

func NewService(api API) Service {
	return &service{
		api: api,
	}
}

type service struct {
	api API
}

func (s *service) IsLaunchpadAvailable(launchpadID string, launchDate date.Date, out chan<- booking.AvailabilityResponse) {
	launchpad, err := s.api.GetLaunchPad(launchpadID)
	if err != nil {
		out <- booking.AvailabilityResponse{Available: false, Err: httpError.NewHTTPError(400, fmt.Sprintf("Launch pad with id %v does not exist", launchpadID)),}
		return
	}
	if launchpad.Status != "active" {
		out <- booking.AvailabilityResponse{Available: false, Err: httpError.NewHTTPError(400, fmt.Sprintf("Launch pad with id %v is not active", launchpadID)),}
		return
	}
	launches, err := s.api.ListUpcomingLaunches(launchpadID)
	if err != nil {
		out <- booking.AvailabilityResponse{Available: false, Err: err,}
		return
	}
	launchDateStr := launchDate.Format(date.Format)
	for _, launch := range launches {
		if launch.LaunchDateLocal.Format(date.Format) == launchDateStr {
			out <- booking.AvailabilityResponse{Available: false, Err: httpError.NewHTTPError(400, "Given launch pad is already booked for selected day"),}
			return
		}
	}
	out <- booking.AvailabilityResponse{Available: true, Err: nil }
}
