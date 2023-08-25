package order_handler

import (
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"web-service-test/entities"
	authentication "web-service-test/entities/authentication_entity"
	"web-service-test/entities/orders_entity"
	"web-service-test/usecases/order_usecase"
)

type OrderHandler struct {
	orderUseCase order_usecase.OrderUseCase
}

func NewOrderHandler(orderUseCase order_usecase.OrderUseCase) OrderHandler {
	return OrderHandler{orderUseCase: orderUseCase}
}

func (oh OrderHandler) CreateOrder(c echo.Context) error {
	tokenLogin := c.Get("user").(*jwt.Token)
	userId := tokenLogin.Claims.(*authentication.TokenClaim).UserId

	bindCreateOrder := new(orders_entity.BindCreateOrder)
	errBindLogin := c.Bind(bindCreateOrder)
	if errBindLogin != nil {
		return c.JSON(400, entities.ResponseMessage{
			Status:  400,
			Message: "error bind data",
			Result:  errBindLogin,
		})
	}
	bindCreateOrder.UserId = userId
	ok, err := oh.orderUseCase.InsertOrder(bindCreateOrder)
	if !ok {
		return c.JSON(400, entities.ResponseMessage{
			Status:  400,
			Message: "create order failed",
			Result:  err,
		})
	}
	return c.JSON(200, entities.ResponseMessage{
		Status:  200,
		Message: "order created",
	})
}

func (oh OrderHandler) CancelOrder(c echo.Context) error {
	tokenLogin := c.Get("user").(*jwt.Token)
	userId := tokenLogin.Claims.(*authentication.TokenClaim).UserId
	paramOrderId := c.QueryParam("id")
	ok, err := oh.orderUseCase.CancelOrder(paramOrderId, userId)
	if !ok {
		return c.JSON(400, entities.ResponseMessage{
			Status:  400,
			Message: "cancel order failed",
			Result:  err,
		})
	}
	return c.JSON(200, entities.ResponseMessage{
		Status:  200,
		Message: "order cancel",
	})
}

func (oh OrderHandler) OrderDetail(c echo.Context) error {
	tokenLogin := c.Get("user").(*jwt.Token)
	userId := tokenLogin.Claims.(*authentication.TokenClaim).UserId
	paramOrderId := c.QueryParam("id")
	ok, err, data := oh.orderUseCase.OrderDetail(paramOrderId, userId)
	if !ok {
		return c.JSON(400, entities.ResponseMessage{
			Status:  400,
			Message: "cancel order failed",
			Result:  err,
		})
	}
	return c.JSON(200, entities.ResponseMessage{
		Status:  200,
		Message: "order canceled",
		Result:  data,
	})
}
