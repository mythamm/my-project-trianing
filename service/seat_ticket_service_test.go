package service

import (
	"fmt"
	"haxagonal-train/repository"
	"reflect"
	"testing"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Test_seatTicketService_CheckAllSeatAvailable(t *testing.T) {
	db, err := gorm.Open(sqlite.Open(bookingDb), &gorm.Config{})
	if err != nil {
		fmt.Println("Error connect database : ", err)
	}

	type fields struct {
		seatTicketRepository repository.SeatTicketRepository
	}
	tests := []struct {
		name    string
		fields  fields
		want    SeatTicketResponse
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "case1_check_all_seat_available_success",
			fields:  fields{repository.NewSeatTicketRepositoryDb(db)},
			want:    SeatTicketResponse{All_available_seat: 6, Last_time_updated: time.Now().Format("2006-01-02 15:04:05"), Availble_seat: "", Zone: ""},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := seatTicketService{
				seatTicketRepository: tt.fields.seatTicketRepository,
			}
			got, err := s.CheckAllSeatAvailable()
			if (err != nil) != tt.wantErr {
				t.Errorf("seatTicketService.CheckAllSeatAvailable() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("seatTicketService.CheckAllSeatAvailable() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewSeatTicketService(t *testing.T) {
	db, err := gorm.Open(sqlite.Open(bookingDb), &gorm.Config{})
	if err != nil {
		fmt.Println("Error connect database : ", err)
	}

	type args struct {
		seatTicketRepository repository.SeatTicketRepository
	}
	tests := []struct {
		name string
		args args
		want seatTicketService
	}{
		// TODO: Add test cases.
		{
			name : "case1_new_seat_ticket_service",
			args: args{repository.NewSeatTicketRepositoryDb(db)},
			want: NewSeatTicketService(repository.NewSeatTicketRepositoryDb(db)),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewSeatTicketService(tt.args.seatTicketRepository); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSeatTicketService() = %v, want %v", got, tt.want)
			}
		})
	}
}
