package main

import (
	"goauth/config"
	"goauth/database"
	"goauth/route"

	"github.com/gin-gonic/gin"
)

func main() {
	g := gin.New()
	db := database.InitPostgresDB()

	// Initialize Google OAuth configuration
	googleOAuthConfig := config.ConfigGoogleAuth()
	route.Route(g, db, googleOAuthConfig)

	g.Run(":8080")
}
