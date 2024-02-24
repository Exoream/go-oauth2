package config

import (

	"log"
	"os"

	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

func ConfigGoogleAuth() oauth2.Config {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}

	GoogleLoginConfig := oauth2.Config{
		RedirectURL:  os.Getenv("GOOGLE_REDIRECT_URL"),
		ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
		Scopes:       []string{
			"https://www.googleapis.com/auth/userinfo.profile", 
			"https://www.googleapis.com/auth/userinfo.email",
		},
		Endpoint:     google.Endpoint,
	}

	return GoogleLoginConfig
}
