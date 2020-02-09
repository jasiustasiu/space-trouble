package spacex

import (
	"fmt"
	"space-trouble/internal/date"
	"space-trouble/internal/httpError"
)

type Service interface {
	IsLaunchpadAvailable(launchpadID string, launchDate date.Date) (bool, error)
}

func NewService(api API) Service {
	return &service{
		api: api,
	}
}

type service struct {
	api API
}

func (s *service) IsLaunchpadAvailable(launchpadID string, launchDate date.Date) (bool, error) {
	launchpad, err := s.api.GetLaunchPad(launchpadID)
	if err != nil {
		return false, httpError.NewHTTPError(400, fmt.Sprintf("Launch pad with id %v does not exist", launchpadID))
	}
	if launchpad.Status != "active" {
		return false, httpError.NewHTTPError(400, fmt.Sprintf("Launch pad with id %v is not active", launchpadID))
	}

	launches, err := s.api.ListUpcomingLaunches(launchpadID)
	if err != nil {
		return false, err
	}
	launchDateStr := launchDate.Format(date.Format)
	for _, launch := range launches {
		if launch.LaunchDateLocal.Format(date.Format) == launchDateStr {
			return false, httpError.NewHTTPError(400, "Given launch pad is already booked for selected day")
		}
	}
	return true, nil
}
