package controller

import (
	"goauth/dto"
	"goauth/service"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService service.UserServiceInterface
}

func NewUserController(userService service.UserServiceInterface) *UserController {
	return &UserController{
		userService: userService,
	}
}

func (uc *UserController) LoginGoogle(c *gin.Context) {
	code := c.Query("code")
	if code == "" {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "Authorization code not provided!"})
		return
	}

	user, err := uc.userService.CreateUserFromGoogleCode(code)
	if err != nil {
		log.Printf("Failed to create user from Google OAuth: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"status": "fail", "message": "Internal Server Error"})
		return
	}

	responseUser := dto.ResponseUser{
        Name:  user.Name,
        Email: user.Email,
        Image: user.Image,
    }

	c.JSON(http.StatusOK, gin.H{"status": "success", "data": responseUser})
}

func (uc *UserController) GoogleCallback(c *gin.Context) {
	err := uc.userService.HandleGoogleLogin(c.Writer, c.Request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
}