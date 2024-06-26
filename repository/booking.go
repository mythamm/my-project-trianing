package repository

// type Booking struct {
// 	Booking_id     string `db:"booking_id"`
// 	User_id        string `db:"user_id"`
// 	Ticket_name    string `db:"ticket_name"`
// 	Booking_seat   string `db:"booking_seat"`
// 	Payment_method string `db:"payment_method"`
// 	Payment_status string `db:"payment_status"`
// 	Seat_id        int    `db:"seat_id"`
// 	Amount         int    `db:"amount"`
// }

type Booking struct {
	Booking_id     string `gorm:"booking_id"`
	User_id        string `gorm:"user_id"`
	Ticket_name    string `gorm:"ticket_name"`
	Booking_seat   string `gorm:"booking_seat"`
	Payment_method string `gorm:"payment_method"`
	Payment_status string `gorm:"payment_status"`
	Seat_id        int    `gorm:"seat_id"`
	Amount         int    `gorm:"amount"`
}

type BookingRepository interface {
	GetAll() ([]Booking, error)
	GetById(string) ([]Booking, error)
	CreateNew(Booking) error
}
