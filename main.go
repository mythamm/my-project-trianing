package main

import (
	"haxagonal-train/handler"
	"haxagonal-train/repository"
	"haxagonal-train/service"

	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"gorm.io/driver/sqlite"
)

const bookingDb = "/Users/a677161/Code/haxagonal/booking.db"

func main() {
	r := gin.New()

	db, err := gorm.Open(sqlite.Open(bookingDb), &gorm.Config{})
	if err != nil {
		fmt.Println("Error connect database : ", err)
	}
	fmt.Println("Connect database : ", db)
	// Migrate the schema
	db.AutoMigrate(&repository.Booking{})
	db.AutoMigrate(&repository.SeatTicket{})
	db.AutoMigrate(&repository.SeatTicketInfo{})
	db.AutoMigrate(&repository.User_info{})

	// ----- Bookings -----
	bookingRepository := repository.NewBookRepositoryDb(db)
	seatTicketRepository := repository.NewSeatTicketRepositoryDb(db)

	bookingService := service.NewBookingService(bookingRepository, seatTicketRepository)
	bookingHandler := handler.NewBookingHandler(bookingService)

	r.GET("/all-book-seat", bookingHandler.GetBookingAll)
	r.GET("/get-book-seat-by-user", bookingHandler.GetBookingByUser)
	r.POST("/create-book", bookingHandler.CreateNewBooking)
	// ----- Bookings -----

	// ----- User -----
	userRepository := repository.NewUserRepositoryDb(db)
	userService := service.NewUserService(userRepository)
	userHandler := handler.NewUserHandler(userService)

	r.GET("/get-user", userHandler.GetUserById)
	r.POST("/create-user", userHandler.CreateNewUser)
	r.POST("/update-last-login", userHandler.UpdateLastLogin)
	r.POST("/delete-user", userHandler.DeleteUser)
	// ----- User -----

	// ----- Seat Ticket -----
	// seatTicketRepository := repository.NewSeatTicketRepositoryDb(db)
	seatTicketService := service.NewSeatTicketService(seatTicketRepository)
	seatTicketHandler := handler.NewSeatTicketHandler(seatTicketService)

	r.GET("/get-all-available-seat", seatTicketHandler.GetAllSeatAvailable)
	// ----- Seat Ticket -----

	r.Run()

}
