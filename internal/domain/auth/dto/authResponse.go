package dto

import (
	"github.com/hudayberdipolat/golang-auth-jwtToken/internal/models"
)

type AuthResponse struct {
	ID          uint   `json:"id"`
	Username    string `json:"username"`
	FullName    string `json:"full_name"`
	CreatedAt   string `json:"created_at"`
	AccessToken string `json:"assess_token"`
}

func NewAuthResponse(user models.User, token string) AuthResponse {
	return AuthResponse{
		ID:          user.ID,
		Username:    user.Username,
		FullName:    user.FullName,
		CreatedAt:   user.CreatedAt.Format("01-02-2006"),
		AccessToken: token,
	}
}
