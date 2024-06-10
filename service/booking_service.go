package service

import (
	"fmt"
	"haxagonal-train/common"
	"haxagonal-train/repository"
)

type bookingService struct {
	bookingRepository    repository.BookingRepository
	seatTicketRepository repository.SeatTicketRepository
}

func NewBookingService(bookingRepository repository.BookingRepository, seatTicketRepository repository.SeatTicketRepository) bookingService {
	return bookingService{
		bookingRepository:    bookingRepository,
		seatTicketRepository: seatTicketRepository,
	}
}

func (s bookingService) GetAllBooking() ([]BookingResponse, error) {
	booking, err := s.bookingRepository.GetAll()
	if err != nil {
		fmt.Println("Error : ", err)
		return nil, err
	}
	fmt.Println("Booking data : ", booking)

	var bookingReponse []BookingResponse
	for _, booking := range booking {
		booking := BookingResponse{
			Ticket_name:  booking.Ticket_name,
			Booking_seat: booking.Booking_seat,
			Seat_id:      booking.Seat_id,
		}
		bookingReponse = append(bookingReponse, booking)
	}

	return bookingReponse, nil
}

func (s bookingService) CreateNewBooking(req *BookingRequest) (common.CommonResponse, error) {
	var response common.CommonResponse
	fmt.Println("Data :: ", req.Booking_seat)
	
	// create to Booking table
	data := repository.Booking{
		Booking_id:   req.Booking_id,
		User_id:      req.User_id,
		Ticket_name:  req.Ticket_name,
		Booking_seat: req.Booking_seat,
		// Payment_method: "",
		// Payment_status: "",
		Seat_id: req.Seat_id,
		Amount:  req.Amount,
	}
	err := s.bookingRepository.CreateNew(data)
	if err != nil {
		fmt.Println("Crete Booking Error : ", err)
		response.Status_code = 500
		response.Status_desc = "Crete Booking Error : " + err.Error()
		return response, err
	}

	// data to update seat_ticket info
	seat_data := repository.SeatTicketInfo{
		Seat_id: req.Seat_id,
		// Ticket_id:    req.Ticket_id,
		// Seat_no:      req.Seat_no,
		Booking_flag: "Y",
	}
	err = s.seatTicketRepository.UpdateBookingFlag(seat_data)
	if err != nil {
		fmt.Println("Update Booking Flag Error : ", err)
		response.Status_code = 500
		response.Status_desc = "Crete Booking Error : " + err.Error()
		return response, err
	}

	response.Status_code = 200
	response.Status_desc = "Create New Booking Success"

	return response, nil
}

// GetAllBookingByUser implements BookingService.
func (s bookingService) GetAllBookingByUser(id string) ([]BookingResponse, error) {
	booking, err := s.bookingRepository.GetById(id)
	if err != nil {
		fmt.Println("Error : ", err)
		return nil, err
	}
	fmt.Println("Booking data : ", booking)

	var bookingReponse []BookingResponse
	for _, booking := range booking {
		booking := BookingResponse{
			Ticket_name:  booking.Ticket_name,
			Booking_seat: booking.Booking_seat,
			Seat_id:      booking.Seat_id,
		}
		bookingReponse = append(bookingReponse, booking)
	}

	return bookingReponse, nil
}
