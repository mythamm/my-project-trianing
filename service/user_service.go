package service

import (
	"fmt"
	"haxagonal-train/common"
	"haxagonal-train/repository"
	"time"
)

type userService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) userService {
	return userService{userRepository: userRepository}
}

// CreateUser implements UserService.
func (u userService) CreateUser(req *UserRequest) (common.CommonResponse, error) {
	var response common.CommonResponse

	fmt.Println("req : ", req.User_id)

	data := repository.User_info{
		User_id:  req.User_id,
		Username: req.User_name,
		Created_at: time.Now().Format("2006-01-02 15:04:05"),
		Lasted_login: "",
	}

	_, err := u.userRepository.CreateUser(data)

	if err != nil {
		fmt.Println("Error : ", err)
		response.Status_code = 500
		response.Status_desc = "Cannot create new user : " + err.Error()
		return response , err
	}

	response.Status_code = 200
	response.Status_desc = "Create New User Success"

	return response, nil

}

// GetUserById implements UserService.
func (u userService) GetUserById(id string) (UserResponse, error) {
	var response UserResponse
	user , err := u.userRepository.GetById(id)
	if err != nil {
		fmt.Println("Error : ", err)
		return response, err
	}
	response = UserResponse{
		User_name: user.Username,
		Created_at: user.Created_at,
		Lasted_login: user.Lasted_login,
	}
	
	return response, err
}

// UpdateLastLogin implements UserService.
func (u userService) UpdateLastLogin(id string) (common.CommonResponse, error) {
	var response common.CommonResponse

	err := u.userRepository.UpdateLastLogin(id)

	if err != nil {
		fmt.Println("Error : ", err)
		response.Status_code = 500
		response.Status_desc = "Cannot create new user : " + err.Error()
		return response ,err
	}

	response.Status_code = 200
	response.Status_desc = "Updated last login Success"
	return response, nil
	
}

func (u userService) DeleteUser(id string) (common.CommonResponse, error) {
	var response common.CommonResponse

	fmt.Println("req : ", id)

	err := u.userRepository.DeleteUser(id)

	if err != nil {
		fmt.Println("Error : ", err)
		response.Status_code = 500
		response.Status_desc = "Cannot delete user : " + err.Error()
		return response ,err
	}

	response.Status_code = 200
	response.Status_desc = "Delete user Success"
	return response, nil
	
}