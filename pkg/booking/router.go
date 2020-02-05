package booking

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"net/http"
)

type Router struct {
	service Service
}

func (r *Router) CreateBookingRoute(c *gin.Context) {
	var booking Booking
	if err := c.ShouldBindWith(&booking, binding.Form); err == nil {
		err = r.service.CreateBooking(booking)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		} else {
			c.JSON(http.StatusOK, gin.H{"message": "Booking confirmed"})
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}
