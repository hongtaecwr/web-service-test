package middlewares

import (
	"github.com/golang-jwt/jwt"
	"log"
	"time"
	"web-service-test/config"
	authentication "web-service-test/entities/authentication_entity"
	"web-service-test/entities/database_entity"
)

func GenToken(data *database_entity.Users, expiresIn time.Duration, permission int) (string, int64) {
	expiresAt := int64(0)
	now := time.Now()
	if expiresIn > 0 {
		expiresAt = now.Add(expiresIn).Unix()
	}
	tokenClaim := jwt.NewWithClaims(jwt.SigningMethodRS256, authentication.TokenClaim{
		UserId:     data.ID,
		UserName:   data.FirstName + " " + data.LastName,
		Permission: permission,
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  now.Unix(),
			ExpiresAt: expiresAt,
		},
	})
	privateKey, _ := config.GetKey()
	token, err := tokenClaim.SignedString(privateKey)
	if err != nil {
		log.Println(err)
	}
	return token, expiresAt
}
