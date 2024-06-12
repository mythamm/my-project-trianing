package repository

type SeatTicket struct {
	Ticket_id   int    `gorm:"ticket_id"`
	Ticket_name string `gorm:"ticket_name"`
}

type SeatTicketInfo struct {
	Seat_id      int    `gorm:"seat_id"`
	Ticket_id    int    `gorm:"ticket_id"`
	Zone         string `gorm:"zone"`
	Price        int    `gorm:"price"`
	Seat_no      int    `gorm:"seat_no"`
	Booking_flag string `gorm:"booking_flag"`
}

type SeatTicketRepository interface {
	CheckAllSeatAvailable() (int, error)
	UpdateBookingFlag(SeatTicketInfo) error
}
