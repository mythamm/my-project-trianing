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
	// var booking []Booking
	booking := []Booking{}
	// Get all records
	err := r.db.Find(&booking).Error
	// result := r.db.First(&booking)
	if err != nil {
		fmt.Println("Error : ", err)
	}

	return booking, nil
}

// GetById implements BookingRepository.
func (r bookRepositoryDb) GetById(id string) ([]Booking, error) {
	// var booking []Booking
	booking := []Booking{}

	// err := r.db.First(&booking, id).Error
	result := r.db.Raw("SELECT * from bookings WHERE user_id = ?", id).Scan(&booking)

	if result.Error != nil {
		fmt.Println("Error : ", result.Error)
	}
	return booking, nil
}

// CreateNew implements BookingRepository.
func (r bookRepositoryDb) CreateNew(data Booking) error {
	fmt.Println("Data for create : ", data)
	// result := u.db.Create(&data)
	// result := u.db.Raw("INSERT INTO user_info (user_id, username, created_at, lasted_login) VALUES(?, ?, ?, ?);" ,data.User_id, data.Username ,data.Created_at, "last").Commit().Error
	result := r.db.Exec("INSERT INTO bookings(booking_id, user_id, ticket_name, booking_seat ,seat_id ,amount) VALUES(? ,? ,? ,? ,?, ?);", data.Booking_id, data.User_id, data.Ticket_name, data.Booking_seat, data.Seat_id, data.Amount)

	fmt.Println("result :", result)
	fmt.Println("Error : ", result.Error)
	if result.Error != nil {
		fmt.Println("Error : ", result.Error)
	} else if(result.RowsAffected <=0) {
		err := fmt.Errorf("don't have this record")
		return err
	}
	return result.Error
}
