package booking

import (
	"errors"
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

func (m *RepositoryMock) Save(b Booking) error {
	args := m.Called(b)
	return args.Error(0)
}

type AvailabilityServiceMock struct {
	AvailabilityService
	mock.Mock
}

func (m *AvailabilityServiceMock) IsLaunchpadAvailable(launchpadID string, launchDate date.Date) (bool, error) {
	args := m.Called(launchpadID, launchDate)
	return args.Bool(0), args.Error(1)
}

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
	repositoryMock := new(RepositoryMock)
	availabilityServiceMock := new(AvailabilityServiceMock)
	svc := &service{
		repository:           repositoryMock,
		availabilityServices: []AvailabilityService{availabilityServiceMock},
	}
	availabilityServiceMock.On("IsLaunchpadAvailable", booking.LaunchpadID, booking.LaunchDate).Return(true, nil)
	repositoryMock.On("Save", booking).Return(nil)
	err := svc.CreateBooking(booking)
	assert.Nil(t, err)
	availabilityServiceMock.AssertCalled(t, "IsLaunchpadAvailable", booking.LaunchpadID, booking.LaunchDate)
	repositoryMock.AssertCalled(t, "Save", booking)
}

func TestShouldNotCreateBookingBecauseLaunchpadIsBusy(t *testing.T) {
	booking := Booking{
		FirstName:     "Kim",
		LastName:      "Dzong Un",
		Gender:        "M",
		Birthday:      date.Date(time.Date(1984, 1, 8, 0, 0, 0, 0, time.UTC)),
		LaunchpadID:   "ccafs_slc_40",
		DestinationID: "europa",
		LaunchDate:    date.Date(time.Date(2049, 2, 4, 12, 0, 0, 0, time.UTC)),
	}
	repositoryMock := new(RepositoryMock)
	availabilityServiceMock := new(AvailabilityServiceMock)
	svc := &service{
		repository:           repositoryMock,
		availabilityServices: []AvailabilityService{availabilityServiceMock},
	}
	availabilityServiceMock.On("IsLaunchpadAvailable", booking.LaunchpadID, booking.LaunchDate).Return(false, nil)
	err := svc.CreateBooking(booking)
	assert.NotNil(t, err)
	availabilityServiceMock.AssertCalled(t, "IsLaunchpadAvailable", booking.LaunchpadID, booking.LaunchDate)
	repositoryMock.AssertNotCalled(t, "Save", mock.Anything)
}

func TestShouldNotCreateBookingBecauseCheckingForLaunchpadReturnedError(t *testing.T) {
	booking := Booking{
		FirstName:     "Kim",
		LastName:      "Dzong Un",
		Gender:        "M",
		Birthday:      date.Date(time.Date(1984, 1, 8, 0, 0, 0, 0, time.UTC)),
		LaunchpadID:   "ccafs_slc_40",
		DestinationID: "europa",
		LaunchDate:    date.Date(time.Date(2049, 2, 4, 12, 0, 0, 0, time.UTC)),
	}
	repositoryMock := new(RepositoryMock)
	availabilityServiceMock := new(AvailabilityServiceMock)
	svc := &service{
		repository:           repositoryMock,
		availabilityServices: []AvailabilityService{availabilityServiceMock},
	}
	availabilityServiceMock.On("IsLaunchpadAvailable", booking.LaunchpadID, booking.LaunchDate).Return(true, errors.New("some error"))
	err := svc.CreateBooking(booking)
	assert.NotNil(t, err)
	availabilityServiceMock.AssertCalled(t, "IsLaunchpadAvailable", booking.LaunchpadID, booking.LaunchDate)
	repositoryMock.AssertNotCalled(t, "Save", mock.Anything)
}
