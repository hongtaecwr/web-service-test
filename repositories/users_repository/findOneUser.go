package users_repository

import "web-service-test/entities/user_entity"

func (userRepo UserRepository) FindOneUser(id int) (bool, error, *user_entity.ResponseUserDetail) {
	user := new(user_entity.ResponseUserDetail)
	findOneUser := userRepo.db.Where("users.id = ?", id).Find(&user)
	if findOneUser.RowsAffected == 0 {
		return false, findOneUser.Error, nil
	}
	return true, nil, user
}
