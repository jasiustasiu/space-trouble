package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
	"space-trouble/pkg/booking"
	"space-trouble/pkg/spacex"
)

const (
	spacexBaseURL = "https://api.spacexdata.com"
	bookingURL = "/v1/bookings"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "space_trouble_user"
	password = "tabeo123"
	dbname   = "space_trouble"
)

func main() {

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sqlx.Connect("postgres", psqlInfo)
	if err != nil {
		log.Fatalln(err)
	}


	spacexAPI := spacex.NewAPI(spacexBaseURL)
	spacexService := spacex.NewService(spacexAPI)
	bookingRepository := booking.NewRepository(db)
	bookingService := booking.NewBookingService(bookingRepository, spacexService)
	router := booking.NewRouter(bookingService)

	r := gin.Default()
	r.POST(bookingURL, router.CreateBookingRoute)
	r.GET(bookingURL, router.GetBookingsRoute)
	r.Run()
}
