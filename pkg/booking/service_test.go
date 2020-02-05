package booking

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

var service = Service{}

func TestShouldCreateBooking(t *testing.T) {
	booking := Booking{
		FirstName:     "Kim",
		LastName:      "Dzong Un",
		Gender:        "M",
		Birthday:      time.Date(1984, 1, 8, 0, 0, 0, 0, time.UTC),
		LaunchpadID:   "ccafs_slc_40",
		DestinationID: "456",
		LaunchDate:    time.Date(2049, 2, 4, 12, 0, 0, 0, time.UTC),
	}

	err := service.CreateBooking(booking)
	assert.Nil(t, err)
}
