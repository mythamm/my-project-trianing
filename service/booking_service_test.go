package service

import (
	"fmt"
	"haxagonal-train/common"
	"haxagonal-train/repository"
	"reflect"
	"testing"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestNewBookingService(t *testing.T) {
	db, err := gorm.Open(sqlite.Open(bookingDb), &gorm.Config{})
	if err != nil {
		fmt.Println("Error connect database : ", err)
	}
	type args struct {
		bookingRepository    repository.BookingRepository
		seatTicketRepository repository.SeatTicketRepository
	}
	tests := []struct {
		name string
		args args
		want bookingService
	}{
		// TODO: Add test cases.
		{
			name: "case1_new_booking_service",
			args: args{repository.NewBookRepositoryDb(db), repository.NewSeatTicketRepositoryDb(db)},
			want: NewBookingService(repository.NewBookRepositoryDb(db), repository.NewSeatTicketRepositoryDb(db)),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewBookingService(tt.args.bookingRepository, tt.args.seatTicketRepository); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewBookingService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_bookingService_GetAllBooking(t *testing.T) {
	db, err := gorm.Open(sqlite.Open(bookingDb), &gorm.Config{})
	if err != nil {
		fmt.Println("Error connect database : ", err)
	}
	type fields struct {
		bookingRepository    repository.BookingRepository
		seatTicketRepository repository.SeatTicketRepository
	}
	tests := []struct {
		name    string
		fields  fields
		want    []BookingResponse
		wantErr bool
	}{
		// TODO: Add test cases.
		// [
		// 	{
		// 		"ticket_name": "MARATHON THUNDER DAY",
		// 		"booking_seat": "B02",
		// 		"seat_id": 2
		// 	},
		// 	{
		// 		"ticket_name": "MARATHON THUNDER DAY",
		// 		"booking_seat": "B03",
		// 		"seat_id": 3
		// 	}
		// ]

		{
			name: "case1_get_all_booking_success",
			fields: fields{repository.NewBookRepositoryDb(db), repository.NewSeatTicketRepositoryDb(db)},
			want: []BookingResponse{{Ticket_name: "MARATHON THUNDER DAY", Booking_seat: "B02" ,Seat_id: 2}, {Ticket_name: "MARATHON THUNDER DAY", Booking_seat: "B03" ,Seat_id: 3}},
			wantErr: false,
		},
		// {
		// 	name: "case2_get_all_booking_fail",
		// 	fields: fields{repository.NewBookRepositoryDb(db), repository.NewSeatTicketRepositoryDb(db)},
		// 	want: []BookingResponse{},
		// 	wantErr: false,
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := bookingService{
				bookingRepository:    tt.fields.bookingRepository,
				seatTicketRepository: tt.fields.seatTicketRepository,
			}
			got, err := s.GetAllBooking()
			if (err != nil) != tt.wantErr {
				t.Errorf("bookingService.GetAllBooking() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("bookingService.GetAllBooking() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_bookingService_GetAllBookingByUser(t *testing.T) {
	db, err := gorm.Open(sqlite.Open(bookingDb), &gorm.Config{})
	if err != nil {
		fmt.Println("Error connect database : ", err)
	}
	type fields struct {
		bookingRepository    repository.BookingRepository
		seatTicketRepository repository.SeatTicketRepository
	}
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []BookingResponse
		wantErr bool
	}{
		{
			name : "case1_get_all_booking_by_user_success",
			fields: fields{repository.NewBookRepositoryDb(db), repository.NewSeatTicketRepositoryDb(db)},
			args: args{"USID00001"},
			want: []BookingResponse{{Ticket_name: "MARATHON THUNDER DAY", Booking_seat: "B02", Seat_id: 2}},
			wantErr: false,
		},
		// {
		// 	name : "case2_get_all_booking_by_user_fail",
		// 	fields: fields{repository.NewBookRepositoryDb(db), repository.NewSeatTicketRepositoryDb(db)},
		// 	args: args{""},
		// 	want: []BookingResponse{BookingResponse{}},
		// 	wantErr: true,
		// },
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := bookingService{
				bookingRepository:    tt.fields.bookingRepository,
				seatTicketRepository: tt.fields.seatTicketRepository,
			}
			got, err := s.GetAllBookingByUser(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("bookingService.GetAllBookingByUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("bookingService.GetAllBookingByUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_bookingService_CreateNewBooking(t *testing.T) {
	db, err := gorm.Open(sqlite.Open(bookingDb), &gorm.Config{})
	if err != nil {
		fmt.Println("Error connect database : ", err)
	}
	type fields struct {
		bookingRepository    repository.BookingRepository
		seatTicketRepository repository.SeatTicketRepository
	}
	type args struct {
		req *BookingRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    common.CommonResponse
		wantErr bool
	}{
		// TODO: Add test cases.
		// {
		// 	"booking_id": "AB100620240001",
		// 	"user_id": "USID00001",
		// 	"ticket_name": "MARATHON THUNDER DAY",
		// 	"booking_seat": "B02",
		// 	"seat_id": 2,
		// 	"amount": 2
		// }

		{
			name: "case1_create_new_booking_success",
			fields: fields{repository.NewBookRepositoryDb(db), repository.NewSeatTicketRepositoryDb(db)},
			args: args{&BookingRequest{Booking_id: "AB100620240009", User_id: "USID00002" , Ticket_name: "MARATHON THUNDER DAY", Booking_seat: "B03", Seat_id: 3, Amount: 1}},
			want: common.CommonResponse{Status_code: 200, Status_desc: "Create New Booking Success"},
			wantErr: false,
		},
		{
			name: "case2_create_new_booking_fail",
			fields: fields{repository.NewBookRepositoryDb(db), repository.NewSeatTicketRepositoryDb(db)},
			args: args{&BookingRequest{Booking_id: "AB100620240009", User_id: "USID00002" , Ticket_name: "MARATHON THUNDER DAY", Booking_seat: "B03", Seat_id: 3, Amount: 1}},
			want: common.CommonResponse{Status_code: 500, Status_desc: "Crete Booking Error : UNIQUE constraint failed: bookings.booking_id"},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := bookingService{
				bookingRepository:    tt.fields.bookingRepository,
				seatTicketRepository: tt.fields.seatTicketRepository,
			}
			got, err := s.CreateNewBooking(tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("bookingService.CreateNewBooking() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("bookingService.CreateNewBooking() = %v, want %v", got, tt.want)
			}
		})
	}
}


