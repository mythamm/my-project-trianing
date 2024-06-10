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

const bookingDb = "/Users/a677161/Code/haxagonal/booking.db"

func Test_userService_CreateUser(t *testing.T) {

	db, err := gorm.Open(sqlite.Open(bookingDb), &gorm.Config{})
	if err != nil {
		fmt.Println("Error connect database : ", err)
	}

	type fields struct {
		userRepository repository.UserRepository
	}
	type args struct {
		req *UserRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    common.CommonResponse
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "case1_create_user_success",
			fields:  fields{userRepository: repository.NewUserRepositoryDb(db)},
			args:    args{&UserRequest{User_id: "UID00000001", User_name: "Test Create"}},
			want:    common.CommonResponse{Status_code: 200, Status_desc: "Create New User Success"},
			wantErr: false,
		},
		{
			name:    "case2_create_user_fail_duplicate_user_id",
			fields:  fields{userRepository: repository.NewUserRepositoryDb(db)},
			args:    args{&UserRequest{User_id: "UID00000001", User_name: "Test Create"}},
			want:    common.CommonResponse{Status_code: 500, Status_desc: "Cannot create new user : UNIQUE constraint failed: user_info.user_id"},
			wantErr: true,
		},
		// {
		// 	name:    "case3_create_user_fail_req_null",
		// 	fields:  fields{userRepository: repository.NewUserRepositoryDb(db)},
		// 	args:    args{&UserRequest{}},
		// 	want:    common.CommonResponse{Status_code: 400, Status_desc:"Validation error: Key: 'UserRequest.User_id' Error:Field validation for 'User_id' failed on the 'required' tag"},
		// 	wantErr: true,
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := userService{
				userRepository: tt.fields.userRepository,
			}
			got, err := u.CreateUser(tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("userService.CreateUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("userService.CreateUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_userService_GetUserById(t *testing.T) {
	db, err := gorm.Open(sqlite.Open(bookingDb), &gorm.Config{})
	if err != nil {
		fmt.Println("Error connect database : ", err)
	}

	type fields struct {
		userRepository repository.UserRepository
	}
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    UserResponse
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "case1_get_user_by_id_success",
			fields:  fields{repository.NewUserRepositoryDb(db)},
			args:    args{"USID00001"},
			want:    UserResponse{User_name: "tue4_jun", Created_at: "2024-06-04 14:44:06", Lasted_login: "2024-06-09 21:15:32"},
			wantErr: false,
		},
		{
			name:    "case2_get_user_by_id_success_no_data",
			fields:  fields{repository.NewUserRepositoryDb(db)},
			args:    args{""},
			want:    UserResponse{User_name: "", Created_at: "", Lasted_login: ""},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := userService{
				userRepository: tt.fields.userRepository,
			}
			got, err := u.GetUserById(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("userService.GetUserById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("userService.GetUserById() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_userService_UpdateLastLogin(t *testing.T) {
	db, err := gorm.Open(sqlite.Open(bookingDb), &gorm.Config{})
	if err != nil {
		fmt.Println("Error connect database : ", err)
	}

	type fields struct {
		userRepository repository.UserRepository
	}
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    common.CommonResponse
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "case1_update_last_login_success",
			fields:  fields{repository.NewUserRepositoryDb(db)},
			args:    args{"USID00002"},
			want:    common.CommonResponse{Status_code: 200, Status_desc: "Updated last login Success"},
			wantErr: false,
		},
		{
			name:    "case2_update_last_login_fail",
			fields:  fields{repository.NewUserRepositoryDb(db)},
			args:    args{"USID0000XXX"},
			want:    common.CommonResponse{Status_code: 500, Status_desc: "Cannot create new user : don't have this record"},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := userService{
				userRepository: tt.fields.userRepository,
			}
			got, err := u.UpdateLastLogin(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("userService.UpdateLastLogin() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("userService.UpdateLastLogin() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_userService_DeleteUser(t *testing.T) {
	db, err := gorm.Open(sqlite.Open(bookingDb), &gorm.Config{})
	if err != nil {
		fmt.Println("Error connect database : ", err)
	}

	type fields struct {
		userRepository repository.UserRepository
	}
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    common.CommonResponse
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "case1_delete_user_success",
			fields:  fields{repository.NewUserRepositoryDb(db)},
			args:    args{"USID00004"},
			want:    common.CommonResponse{Status_code: 200, Status_desc: "Delete user Success"},
			wantErr: false,
		},
		{
			name:    "case2_delete_user_fail",
			fields:  fields{repository.NewUserRepositoryDb(db)},
			args:    args{"UID00000009"},
			want:    common.CommonResponse{Status_code: 500, Status_desc: "Cannot delete user : don't have this record"},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := userService{
				userRepository: tt.fields.userRepository,
			}
			got, err := u.DeleteUser(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("userService.DeleteUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("userService.DeleteUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewUserService(t *testing.T) {
	db, err := gorm.Open(sqlite.Open(bookingDb), &gorm.Config{})
	if err != nil {
		fmt.Println("Error connect database : ", err)
	}

	type args struct {
		userRepository repository.UserRepository
	}
	tests := []struct {
		name string
		args args
		want userService
	}{
		// TODO: Add test cases.
		{
			name: "new_user_service",
			args: args{userRepository: repository.NewUserRepositoryDb(db)},
			want: NewUserService(repository.NewUserRepositoryDb(db)),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewUserService(tt.args.userRepository); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUserService() = %v, want %v", got, tt.want)
			}
		})
	}
}
