package user_usecase

import (
	"web-service-test/entities/user_entity"
	"web-service-test/repositories/users_repository"
)

type UserUseCase struct {
	userRepo users_repository.UserRepository
}

func NewUserUseCase(userRepo users_repository.UserRepository) UserUseCase {
	return UserUseCase{userRepo: userRepo}
}

func (us UserUseCase) GetUserDetail(id int) (bool, error, *user_entity.ResponseUserDetail) {
	ok, err, user := us.userRepo.FindOneUser(id)
	if !ok {
		return false, err, nil
	}
	return true, nil, user
}

func (us UserUseCase) GetUserOrderHistories(id int) (bool, error, *[]user_entity.ResponseOrderHistories) {
	ok, err, user := us.userRepo.FindAllOrderHistories(id)
	if !ok {
		return false, err, nil
	}
	return true, nil, user
}
