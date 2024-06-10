package handler

import (
	"fmt"
	"haxagonal-train/common"
	"haxagonal-train/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

type userHandler struct {
	userService service.UserService
}

func NewUserHandler(userService service.UserService) userHandler {
	return userHandler{
		userService: userService,
	}
}

func (h *userHandler) GetUserById(c *gin.Context) {
	fmt.Println("----- Get User By Id ----")

	var user service.UserResponse

	req := service.UserRequest{}
	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": err.Error(), 
		})
		return
	}

	user, _ = h.userService.GetUserById(req.User_id)
	fmt.Println("Get User response : ", user)

	c.JSON(http.StatusOK, user)
}

func (h *userHandler) CreateNewUser(c *gin.Context) {
	fmt.Println("----- Create user ----")
	var user common.CommonResponse

	req := service.UserRequest{}

	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": err.Error(),
		})
		return
	}

	// Validate the request
	var validate = validator.New()
	if err := validate.Struct(req); err != nil {
		user.Status_code = http.StatusBadRequest
		user.Status_desc =fmt.Sprintf("Validation error: %s", err)
		c.JSON(http.StatusBadRequest, user)
		return
	}

	user, err = h.userService.CreateUser(&req)
	if err != nil {
		c.JSON(http.StatusBadGateway, user)
	} else {
		c.JSON(http.StatusOK, user)
	}
	fmt.Println("Create User response : ", user)
}

func (h *userHandler) UpdateLastLogin(c *gin.Context) {
	fmt.Println("----- Update last login ----")
	var user common.CommonResponse

	req := service.UserRequest{}
	err := c.BindJSON(&req)
	fmt.Println("req " , req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": err.Error(),
		})
		return
	}

	// Validate the request
	var validate = validator.New()
	if err := validate.Struct(req); err != nil {
		user.Status_code = http.StatusBadRequest
		user.Status_desc =fmt.Sprintf("Validation error: %s", err)
		c.JSON(http.StatusBadRequest, user)
		return
	}

	user, err = h.userService.UpdateLastLogin(req.User_id)
	if err != nil {
		c.JSON(http.StatusBadGateway, user)
	} else {
		c.JSON(http.StatusOK, user)
	}
}

func (h *userHandler) DeleteUser(c *gin.Context) {
	fmt.Println("----- Delete user ----")
	var user common.CommonResponse

	req := service.UserRequest{}
	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": err.Error(), 
		})
		return
	}

	user, err = h.userService.DeleteUser(req.User_id)
	if err != nil {
		c.JSON(http.StatusBadGateway, user)
	} else {
		c.JSON(http.StatusOK, user)
	}
}
