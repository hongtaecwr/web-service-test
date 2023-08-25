package user_handler

import (
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"web-service-test/entities"
	authentication "web-service-test/entities/authentication_entity"
	"web-service-test/usecases/user_usecase"
)

type UserHandler struct {
	userUseCase user_usecase.UserUseCase
}

func NewUserHandler(userUseCase user_usecase.UserUseCase) *UserHandler {
	return &UserHandler{userUseCase: userUseCase}
}

func (uh UserHandler) GetUserDetail(c echo.Context) error {
	tokenLogin := c.Get("user").(*jwt.Token)
	userId := tokenLogin.Claims.(*authentication.TokenClaim).UserId
	ok, err, data := uh.userUseCase.GetUserDetail(userId)
	if !ok {
		return c.JSON(400, entities.ResponseMessage{
			Status:  400,
			Message: "user not found",
			Result:  err,
		})
	}

	return c.JSON(200, entities.ResponseMessage{
		Status:  200,
		Message: "user found",
		Result:  data,
	})
}

func (uh UserHandler) GetUserOrderHistories(c echo.Context) error {
	tokenLogin := c.Get("user").(*jwt.Token)
	userId := tokenLogin.Claims.(*authentication.TokenClaim).UserId
	ok, err, data := uh.userUseCase.GetUserOrderHistories(userId)
	if !ok {
		return c.JSON(400, entities.ResponseMessage{
			Status:  400,
			Message: "orders not found",
			Result:  err,
		})
	}

	return c.JSON(200, entities.ResponseMessage{
		Status:  200,
		Message: "orders found",
		Result:  data,
	})
}
