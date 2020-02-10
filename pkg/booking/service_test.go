package booking

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"space-trouble/internal/date"
	"testing"
	"time"
)

type RepositoryMock struct {
	repository
	mock.Mock
}

type AvailabilityServiceMock struct {
	AvailabilityService
	mock.Mock
}

var (
	repositoryMock          = new(RepositoryMock)
	availabilityServiceMock = new(AvailabilityServiceMock)
	svc                     = service{
		repository:           repositoryMock,
		availabilityServices: []AvailabilityService{availabilityServiceMock},
	}
)

func TestShouldCreateBooking(t *testing.T) {
	booking := Booking{
		FirstName:     "Kim",
		LastName:      "Dzong Un",
		Gender:        "M",
		Birthday:      date.Date(time.Date(1984, 1, 8, 0, 0, 0, 0, time.UTC)),
		LaunchpadID:   "ccafs_slc_40",
		DestinationID: "europa",
		LaunchDate:    date.Date(time.Date(2049, 2, 4, 12, 0, 0, 0, time.UTC)),
	}
	availabilityServiceMock.On("IsLaunchpadAvailable", )
	//repositoryMock.On("")
	err := svc.CreateBooking(booking)
	assert.Nil(t, err)
}
