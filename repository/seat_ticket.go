package repository

type SeatTicket struct {
	Ticket_id   int    `db:"ticket_id"`
	Ticket_name string `db:"ticket_name"`
}

type SeatTicketInfo struct {
	Seat_id      int    `db:"seat_id"`
	Ticket_id    int    `db:"ticket_id"`
	Zone         string `db:"zone"`
	Price        int    `db:"price"`
	Seat_no      int    `db:"seat_no"`
	Booking_flag string `db:"booking_flag"`
}

type SeatTicketRepository interface {
	CheckAllSeatAvailable() (int, error)
	UpdateBookingFlag(SeatTicketInfo) error
}
