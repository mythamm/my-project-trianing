package handler

import (
	"fmt"
	"haxagonal-train/common"
	"haxagonal-train/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

type bookingHandler struct {
	bookingService service.BookingService
	BookingRequest service.BookingRequest
}

func NewBookingHandler(bookingService service.BookingService) bookingHandler {
	return bookingHandler{
		bookingService: bookingService,
	}
}

func (h *bookingHandler) GetBookingAll(c *gin.Context) {
	fmt.Println("----- Get All Booking ----")
	var books []service.BookingResponse

	books, _ = h.bookingService.GetAllBooking()
	fmt.Println("Book response : ", books)

	c.JSON(http.StatusOK, books)
}

func (h *bookingHandler) CreateNewBooking(c *gin.Context) {
	fmt.Println("----- Create New Booking ----")
	req := service.BookingRequest{}
	response := common.CommonResponse{}

	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": err.Error(), 
		})
		return
	}

	// Validate the request
	var validate = validator.New()
	if err := validate.Struct(req); err != nil {
		response.Status_code = http.StatusBadRequest
		response.Status_desc = fmt.Sprintf("Validation error: %s", err)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	res, err := h.bookingService.CreateNewBooking(&req)
	fmt.Println("Create New Booking : ", res)

	if err != nil {
		c.JSON(http.StatusBadGateway, res)
	} else {
		c.JSON(http.StatusOK, res)
	}
}

func (h *bookingHandler) GetBookingByUser(c *gin.Context) {
	fmt.Println("----- Get All Booking By Userid ----")
	var books []service.BookingResponse

	req := service.GetBookingByUserReq{}
	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": err.Error(), 
		})
		return
	}

	// Validate the request
	var validate = validator.New()
	if err := validate.Struct(req); err != nil {
		response := common.CommonResponse{}
		response.Status_code = http.StatusBadRequest
		response.Status_desc = fmt.Sprintf("Validation error: %s", err)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	books, _ = h.bookingService.GetAllBookingByUser(req.User_id)
	fmt.Println("Book response : ", books)

	c.JSON(http.StatusOK, books)
}
