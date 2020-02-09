package booking

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"space-trouble/internal/httpError"
)

func NewRouter(service Service) *Router {
	return &Router{
		service:service,
	}
}

type Router struct {
	service Service
}

func (r *Router) CreateBookingRoute(c *gin.Context) {
	var booking Booking
	if err := c.BindJSON(&booking); err == nil {
		err = r.service.CreateBooking(booking)
		if err != nil {
			if httpErr, ok := err.(*httpError.HTTPError); ok {
				c.JSON(httpErr.Status, gin.H{"error": httpErr.Details})
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
			}
		} else {
			c.JSON(http.StatusOK, gin.H{"message": "Booking confirmed"})
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}

func (r *Router) GetBookingsRoute(c *gin.Context) {
	var booking Booking
	if err := c.BindJSON(&booking); err == nil {
		err = r.service.CreateBooking(booking)
		if err != nil {
			if httpErr, ok := err.(*httpError.HTTPError); ok {
				c.JSON(httpErr.Status, gin.H{"error": httpErr.Details})
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
			}
		} else {
			c.JSON(http.StatusOK, gin.H{"message": "Booking confirmed"})
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}
