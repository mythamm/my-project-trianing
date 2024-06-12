package service

import (
	"fmt"
	"haxagonal-train/repository"
	"time"
)

type seatTicketService struct {
	seatTicketRepository repository.SeatTicketRepository
}

// CheckAvailableByZone implements SeatTicketService.
func (s seatTicketService) CheckAvailableByZone() (SeatTicketResponse, error) {
	panic("unimplemented")
}

// UpdatePaymentStatus implements SeatTicketService.
func (s seatTicketService) UpdatePaymentStatus() error {
	panic("unimplemented")
}

func NewSeatTicketService(seatTicketRepository repository.SeatTicketRepository) seatTicketService {
	return seatTicketService{
		seatTicketRepository: seatTicketRepository,
	}
}

// CheckAllSeatAvailable implements SeatTicketService.
func (s seatTicketService) CheckAllSeatAvailable() (SeatTicketResponse, error) {
	response := SeatTicketResponse{}
	seat, err := s.seatTicketRepository.CheckAllSeatAvailable()

	if err != nil {
		fmt.Println("Error : ", err)
		return response, err
	}
	fmt.Println("seat data : ", seat)

	response.All_available_seat = seat
	response.Last_time_updated = time.Now().Format("2006-01-02 15:04:05")

	return response, nil
}
