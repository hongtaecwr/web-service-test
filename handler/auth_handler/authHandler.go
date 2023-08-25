package auth_handler

import (
	"github.com/labstack/echo/v4"
	"strings"
	"web-service-test/entities"
	authentication "web-service-test/entities/authentication_entity"
	"web-service-test/usecases/auth_usecase"
)

type AuthHandler struct {
	authUseCase auth_usecase.AuthUseCase
}

func NewAuthHandler(authUseCase auth_usecase.AuthUseCase) AuthHandler {
	return AuthHandler{authUseCase: authUseCase}
}

func (ah AuthHandler) Register(c echo.Context) error {
	requestRegister := new(authentication.RequestRegister)
	errBindRegister := c.Bind(requestRegister)
	if errBindRegister != nil {
		return c.JSON(400, entities.ResponseMessage{
			Status:  400,
			Message: "error bind data",
			Result:  errBindRegister,
		})
	}

	ok, err := ah.authUseCase.Register(*requestRegister)
	if !ok {
		if strings.Contains(err.Error(), "users.unique_username") {
			return c.JSON(400, entities.ResponseMessage{
				Status:  400,
				Message: "register failed username exist",
				Result:  err.Error(),
			})
		} else {
			return c.JSON(400, entities.ResponseMessage{
				Status:  400,
				Message: "register failed",
				Result:  err.Error(),
			})
		}
	}
	return c.JSON(200, entities.ResponseMessage{
		Status:  200,
		Message: "registered",
		Result:  nil,
	})
}

func (ah AuthHandler) Login(c echo.Context) error {
	requestLogin := new(authentication.RequestLogin)
	errBindLogin := c.Bind(requestLogin)
	if errBindLogin != nil {
		return c.JSON(400, entities.ResponseMessage{
			Status:  400,
			Message: "error bind data",
			Result:  errBindLogin,
		})
	}

	ok, err, res := ah.authUseCase.Login(*requestLogin)
	if !ok {
		return c.JSON(400, entities.ResponseMessage{
			Status:  400,
			Message: "login failed",
			Result:  err,
		})
	}
	return c.JSON(200, entities.ResponseMessage{
		Status:  200,
		Message: "log in success",
		Result:  res,
	})
}
