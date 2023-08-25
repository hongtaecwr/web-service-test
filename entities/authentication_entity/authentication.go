package authentication

import "github.com/golang-jwt/jwt"

type TokenClaim struct {
	UserId     int    `json:"user_id"`
	UserName   string `json:"user_name"`
	Permission int    `json:"permission"`
	jwt.StandardClaims
}

type LoginResponse struct {
	AccessToken string `json:"access_token,omitempty"`
	ExpiresIn   string `json:"expires_in,omitempty"` // uint: seconds
	Permission  int    `json:"permission"`
	Name        string `json:"name,omitempty"`
}
