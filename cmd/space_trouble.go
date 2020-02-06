package main

import (
	"github.com/gin-gonic/gin"
	"space-trouble/pkg/booking"
	"space-trouble/pkg/spacex"
)

const (
	spacexBaseURL = "https://api.spacexdata.com"
	bookingURL = "/v1/bookings"
)

func main() {
	spacexAPI := spacex.NewAPI(spacexBaseURL)
	bookingService := booking.NewBookingService(spacexAPI)
	router := booking.NewRouter(bookingService)

	r := gin.Default()
	r.POST(bookingURL, router.CreateBookingRoute)
	r.GET(bookingURL, router.GetBookingsRoute)
	r.Run()
}
