package route

import (
	"goauth/controller"
	"goauth/repository"
	"goauth/service"

	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"gorm.io/gorm"
)

func Route(g *gin.Engine, db *gorm.DB, oauth oauth2.Config) {
	userRepository := repository.NewRepository(db)
	userService := service.NewUserService(userRepository, oauth)
	userHandler := controller.NewUserController(userService)

	g.GET("/auth", userHandler.GoogleCallback)
	g.GET("/google/callback", userHandler.LoginGoogle)
}
