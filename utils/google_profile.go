package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type GoogleUserInfo struct {
	ID            string `json:"id"`
	Email         string `json:"email"`
	VerifiedEmail bool   `json:"verified_email"`
	Name          string `json:"name"`
	GivenName     string `json:"given_name"`
	FamilyName    string `json:"family_name"`
	Picture       string `json:"picture"`
	Locale        string `json:"locale"`
}

func GetGoogleUserInfo(accessToken string) (GoogleUserInfo, error) {
    resp, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + accessToken)
    if err != nil {
        return GoogleUserInfo{}, fmt.Errorf("failed to get user info from Google: %w", err)
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        return GoogleUserInfo{}, errors.New("failed to get user info from Google")
    }

    var userInfo GoogleUserInfo
    err = json.NewDecoder(resp.Body).Decode(&userInfo)
    if err != nil {
        return GoogleUserInfo{}, fmt.Errorf("failed to decode user info response: %w", err)
    }

    return userInfo, nil
}

