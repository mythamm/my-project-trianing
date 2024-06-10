package repository

import (
	"fmt"

	"gorm.io/gorm"
)

type bookRepositoryDb struct {
	db *gorm.DB
}

func NewBookRepositoryDb(db *gorm.DB) bookRepositoryDb {
	return bookRepositoryDb{db: db}
}

// GetAll implements BookingRepository.
func (r bookRepositoryDb) GetAll() ([]Booking, error) {
	booking := []Booking{}
	err := r.db.Find(&booking).Error
	if err != nil {
		fmt.Println("Error : ", err)
	}
	return booking, nil
}

// GetById implements BookingRepository.
func (r bookRepositoryDb) GetById(id string) ([]Booking, error) {
	booking := []Booking{}

	result := r.db.Raw("SELECT * from bookings WHERE user_id = ?", id).Scan(&booking)

	if result.Error != nil {
		fmt.Println("Error : ", result.Error)
	}
	return booking, nil
}

// CreateNew implements BookingRepository.
func (r bookRepositoryDb) CreateNew(data Booking) error {
	fmt.Println("Create New Booking : ", data)
	result := r.db.Exec("INSERT INTO bookings(booking_id, user_id, ticket_name, booking_seat ,seat_id ,amount) VALUES(? ,? ,? ,? ,?, ?);", data.Booking_id, data.User_id, data.Ticket_name, data.Booking_seat, data.Seat_id, data.Amount)

	if result.Error != nil {
		fmt.Println("Error : ", result.Error)
	} else if(result.RowsAffected <=0) {
		err := fmt.Errorf("cannot create booking")
		return err
	}
	return result.Error
}
