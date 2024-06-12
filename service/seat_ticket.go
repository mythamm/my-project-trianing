package service

type SeatTicketRequest struct {
	Ticket_id string `json:"ticket_id" validate:"required"`
	Zone      int    `json:"zone"`
}

type SeatTicketResponse struct {
	All_available_seat int    `json:"all_avialable_seat"`
	Last_time_updated  string `json:"last_time_updated"`
	Availble_seat      string `json:"avialable_seat"`
	Zone               string `json:"zone"`
}

type SeatTicketService interface {
	CheckAllSeatAvailable() (SeatTicketResponse, error)
	CheckAvailableByZone() (SeatTicketResponse, error)
	UpdatePaymentStatus() (error)
}
