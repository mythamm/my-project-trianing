package handler

import (
	"fmt"
	"haxagonal-train/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type seatTicketHandler struct {
	seatTicketService service.SeatTicketService
	// seatTicketRequest service.SeatTicketRequest
}

func NewSeatTicketHandler(seatTicketService service.SeatTicketService) seatTicketHandler {
	return seatTicketHandler{seatTicketService: seatTicketService}
}

func (s *seatTicketHandler) GetAllSeatAvailable(c *gin.Context) {
	fmt.Println("----- Get All Seat Available ----")
	res, err := s.seatTicketService.CheckAllSeatAvailable()
	fmt.Println("response : ", res)

	if err != nil {
		c.JSON(http.StatusBadGateway , err)
	}

	c.JSON(http.StatusOK, res)
}

