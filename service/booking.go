package service

import "haxagonal-train/common"

type BookingResponse struct {
	// User_id      string `db:"user_id"`
	Ticket_name string `json:"ticket_name"`
	// Booking_id   string `db:"booking_id"`
	Booking_seat string `json:"booking_seat"`
	Seat_id      int    `json:"seat_id"`
}

type BookingRequest struct {
	Booking_id     string `json:"booking_id" validate:"required"`
	User_id        string `json:"user_id"`
	Ticket_name    string `json:"ticket_name"`
	Booking_seat   string `json:"booking_seat"`
	Payment_method string `json:"payment_method"`
	Payment_status string `json:"payment_status"`
	Seat_id        int    `json:"seat_id"`
	Amount         int    `json:"amount"`
}

type GetBookingByUserReq struct {
	User_id        string `json:"user_id" validate:"required"`
}

type BookingService interface {
	GetAllBooking() ([]BookingResponse, error)
	CreateNewBooking(*BookingRequest) (common.CommonResponse, error)
	GetAllBookingByUser(string) ([]BookingResponse, error)
}
