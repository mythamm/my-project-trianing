package service

import "haxagonal-train/common"

type UserResponse struct {
	User_name    string `json:"user_name"`
	Created_at   string `json:"created_at"`
	Lasted_login string `json:"lasted_at"`
}

type UserRequest struct {
	// user_id	username	created_at	lasted_login
	User_id   string `json:"user_id" validate:"required"`
	User_name string `json:"user_name"`
}

type UserService interface {
	GetUserById(id string) (UserResponse, error)
	CreateUser(*UserRequest) (common.CommonResponse, error)
	UpdateLastLogin(string) (common.CommonResponse, error)
	DeleteUser(string) (common.CommonResponse, error)
}
