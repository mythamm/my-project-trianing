package repository

import (
	"fmt"

	"gorm.io/gorm"
)

type seatTicketRepositoryDb struct {
	db *gorm.DB
}

func NewSeatTicketRepositoryDb(db *gorm.DB) seatTicketRepositoryDb {
	return seatTicketRepositoryDb{db: db}
}

// CheckAllSeatAvailable implements SeatTicketRepository.
func (s seatTicketRepositoryDb) CheckAllSeatAvailable() (int, error) {
	// allAvailbleCount := 0
	// result := s.db.Find(&seat)
	seat := []SeatTicketInfo{}
	result := s.db.Raw("SELECT * FROM seat_ticket_info WHERE booking_flag != 'Y';").Scan(&seat)

	if result.Error != nil {
		fmt.Println("Error : ", result.Error)
	}
	return len(seat), nil
}

// UpdateBookingFlag implements SeatTicketRepository.
func (s seatTicketRepositoryDb) UpdateBookingFlag(data SeatTicketInfo) error {
	// data := User{}
	fmt.Println("----- UpdateBookingFlag ------")

	fmt.Println("data : " ,data)

	// err := r.db.First(&booking, id).Error
	// UPDATE user_info SET lasted_login = datetime('now', 'localtime') WHERE user_id='USID00001';
	result := s.db.Exec("UPDATE seat_ticket_info SET booking_flag=? WHERE seat_id=?;", data.Booking_flag ,data.Seat_id)

	if result.Error != nil {
		fmt.Println("Error : ", result.Error)
		return result.Error
	} else if (result.RowsAffected <= 0) {
		err := fmt.Errorf("don't have this record")
		return err
	}
	return nil
}
