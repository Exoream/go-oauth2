package service

import (
	"context"
	"fmt"
	"goauth/helper"
	"goauth/model"
	"goauth/repository"
	"goauth/utils"
	"log"
	"net/http"

	"golang.org/x/oauth2"
)

type UserService struct {
	userRepository    repository.RepositoryInterface
	googleOAuthConfig oauth2.Config
}

type UserServiceInterface interface {
	CreateUserFromGoogleCode(code string) (model.Users, error)
	HandleGoogleLogin(w http.ResponseWriter, r *http.Request) error
}

func NewUserService(userRepository repository.RepositoryInterface, googleOAuthConfig oauth2.Config) UserServiceInterface {
	return &UserService{
		userRepository:    userRepository,
		googleOAuthConfig: googleOAuthConfig,
	}
}

func (us *UserService) CreateUserFromGoogleCode(code string) (model.Users, error) {
	log.Println("Received code:", code)

	// Exchange OAuth code dengan token akses
	token, err := us.googleOAuthConfig.Exchange(context.Background(), code)
	if err != nil {
		return model.Users{}, fmt.Errorf("failed to exchange OAuth code: %w", err)
	}

	// Ambil informasi pengguna dari Google
	userInfo, err := utils.GetGoogleUserInfo(token.AccessToken)
	if err != nil {
		return model.Users{}, fmt.Errorf("failed to get user info from Google: %w", err)
	}

	newPassword, err := utils.GenerateRandomPassword(12)
	if err != nil {
		return model.Users{}, err
	}

	user, err := us.userRepository.FindByEmail(userInfo.Email)
	if err != nil {
		newUser := model.Users{
			Name:     userInfo.Name,
			Email:    userInfo.Email,
			Password: newPassword,
			Image:    userInfo.Picture,
		}

		err := us.userRepository.CreateUser(newUser)
		if err != nil {
			return model.Users{}, fmt.Errorf("failed to register user: %w", err)
		}

		user = newUser
	}

	return user, nil
}

func (gs *UserService) HandleGoogleLogin(w http.ResponseWriter, r *http.Request) error {
	state, _ := helper.GenerateRandomState(6)
	url := gs.googleOAuthConfig.AuthCodeURL(state)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)

	return nil
}
