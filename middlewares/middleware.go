package middlewares

import (
	"github.com/golang-jwt/jwt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
	"web-service-test/config"
	"web-service-test/entities"
	authentication "web-service-test/entities/authentication_entity"
)

// SetCORS set core origin api
func SetCORS(e *echo.Echo) {
	corsConfig := middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{"*"},
		AllowMethods: []string{"*"},
	}
	e.Use(middleware.CORSWithConfig(corsConfig))
}

// SetLogin ตัวกลางเช็ค Login
func SetLogin() echo.MiddlewareFunc {
	_, publicKey := config.GetKey()
	return middleware.JWTWithConfig(middleware.JWTConfig{
		TokenLookup:   "header:" + echo.HeaderAuthorization,
		SigningMethod: "RS256",
		AuthScheme:    "Bearer",
		Claims:        &authentication.TokenClaim{},
		SigningKey:    publicKey,
	})
}

func IsUser(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(*authentication.TokenClaim)
		isMember := claims.Permission
		if isMember != 1 {
			return c.JSON(http.StatusUnauthorized, entities.ResponseMessage{
				Status:  http.StatusUnauthorized,
				Message: "permission denied",
				Result:  "please login again",
			})
		}
		return next(c)
	}
}

func IsAdmin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(*authentication.TokenClaim)
		isAdmin := claims.Permission
		if isAdmin != 1 {
			return c.JSON(http.StatusUnauthorized, entities.ResponseMessage{
				Status:  http.StatusUnauthorized,
				Message: "permission denied",
				Result:  "please login again",
			})
		}
		return next(c)
	}
}
