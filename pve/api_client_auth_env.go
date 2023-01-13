package pve

import (
	"errors"
	"os"
)

type AuthOptions struct {
	APIEndpoint string `json:"apiEndpoint"`

	// Ticket Cookie Authentication
	// https://pve.proxmox.com/wiki/Proxmox_VE_API#Ticket_Cookie
	Username string `json:"username"`
	Password string `json:"password"`
	Otp      string `json:"otp"`
	// if use ticket cookie authentication, support reauth when cookie invalidation,
	// for example, cookie is expired.
	// AllowReauth bool `json:"allowReauth"`

	// API Tokens Authentication
	// https://pve.proxmox.com/wiki/Proxmox_VE_API#API_Tokens
	ApiToken string `json:"apiToken"`
}

func AuthOptionsFromEnv() (*AuthOptions, error) {
	apiEndpoint := os.Getenv("PVE_API_ENDPOINT")
	username := os.Getenv("PVE_USERNAME")
	password := os.Getenv("PVE_PASSWORD")
	otp := os.Getenv("PVE_OTP")
	apiToken := os.Getenv("PVE_API_TOKEN")

	if apiEndpoint == "" {
		return nil, errors.New("PVE_API_ENDPOINT is missing")
	}

	authOptions := AuthOptions{
		APIEndpoint: apiEndpoint,
		Username:    username,
		Password:    password,
		Otp:         otp,
		ApiToken:    apiToken,
	}

	return &authOptions, nil
}
