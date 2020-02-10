package spacex

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"space-trouble/internal/date"
	"testing"
	"time"
)

type APIMock struct {
	api
	mock.Mock
}

func (m *APIMock) ListUpcomingLaunches(siteID string) ([]Launch, error) {
	args := m.Called(siteID)
	return args.Get(0).([]Launch), args.Error(1)
}

func (m *APIMock) GetLaunchPad(siteID string) (LaunchPad, error) {
	args := m.Called(siteID)
	return args.Get(0).(LaunchPad), args.Error(1)
}

func TestShouldReturnTrueForActiveLaunchpadAndEmptyUpcomingLaunches(t *testing.T) {
	launchpadID := "id-1"
	apiMock := new(APIMock)
	svc := service{api: apiMock}
	launchpad := LaunchPad{
		Status: "active",
	}
	apiMock.On("GetLaunchPad", launchpadID).Return(launchpad, nil)
	apiMock.On("ListUpcomingLaunches", launchpadID).Return([]Launch{}, nil)
	available, err := svc.IsLaunchpadAvailable(launchpadID, date.Date(time.Date(2049, 6, 3, 0, 0, 0, 0, time.UTC)))
	assert.Equal(t, true, available)
	assert.Nil(t, err)
}

func TestShouldReturnFalseForInactiveLaunchpad(t *testing.T) {
	launchpadID := "id-1"
	apiMock := new(APIMock)
	svc := service{api: apiMock}
	launchpad := LaunchPad{
		Status: "retired",
	}
	apiMock.On("GetLaunchPad", launchpadID).Return(launchpad, nil)
	apiMock.On("ListUpcomingLaunches", launchpadID).Return([]Launch{}, nil)
	available, err := svc.IsLaunchpadAvailable(launchpadID, date.Date(time.Date(2049, 6, 3, 0, 0, 0, 0, time.UTC)))
	assert.Equal(t, false, available)
	assert.NotNil(t, err)
}

func TestShouldReturnFalseForLaunchAtGivenDay(t *testing.T) {
	launchpadID := "id-1"
	apiMock := new(APIMock)
	svc := service{api: apiMock}
	launchpad := LaunchPad{
		Status: "active",
	}
	launchDate := time.Date(2049, 6, 3, 0, 0, 0, 0, time.UTC)
	launch := Launch{LaunchDateLocal: launchDate}
	apiMock.On("GetLaunchPad", launchpadID).Return(launchpad, nil)
	apiMock.On("ListUpcomingLaunches", launchpadID).Return([]Launch{launch}, nil)
	available, err := svc.IsLaunchpadAvailable(launchpadID, date.Date(launchDate))
	assert.Equal(t, false, available)
	assert.NotNil(t, err)
}
