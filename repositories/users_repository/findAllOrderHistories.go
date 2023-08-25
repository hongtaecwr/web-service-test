package users_repository

import (
	"errors"
	"web-service-test/entities/user_entity"
)

func (userRepo UserRepository) FindAllOrderHistories(userId int) (bool, error, *[]user_entity.ResponseOrderHistories) {
	orderHistories := new([]user_entity.ResponseOrderHistories)
	findAllOrder := userRepo.db.Where("orders.user_id = ?", userId).Find(&orderHistories)
	if findAllOrder.RowsAffected == 0 {
		return false, errors.New("data not found"), nil
	}
	return true, nil, orderHistories
}
