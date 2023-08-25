package order_usecase

import (
	"strconv"
	"web-service-test/entities/orders_entity"
	"web-service-test/repositories/order_repository"
)

type OrderUseCase struct {
	orderRepo order_repository.OrderRepository
}

func NewOrderUseCase(orderRepo order_repository.OrderRepository) OrderUseCase {
	return OrderUseCase{orderRepo: orderRepo}
}

func (ou OrderUseCase) InsertOrder(bindOrder *orders_entity.BindCreateOrder) (bool, error) {
	ok, err := ou.orderRepo.InsertOrder(bindOrder.AddressId, bindOrder.SummaryPrice, bindOrder.UserId, bindOrder.ProductId)
	if !ok {
		return false, err
	}
	return true, nil
}

func (ou OrderUseCase) CancelOrder(orderId string, userId int) (bool, error) {
	intOrderId, _ := strconv.Atoi(orderId)
	ok, err := ou.orderRepo.CancelOrder(intOrderId, userId)
	if !ok {
		return false, err
	}
	return true, nil
}

func (ou OrderUseCase) OrderDetail(orderId string, userId int) (bool, error, *orders_entity.ResponseOrderDetail) {
	intOrderId, _ := strconv.Atoi(orderId)
	ok, err, data := ou.orderRepo.OrderDetail(intOrderId, userId)
	if !ok {
		return false, err, nil
	}
	return true, nil, data
}
