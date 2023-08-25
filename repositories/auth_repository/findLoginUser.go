package auth_repository

import (
	"errors"
	authentication "web-service-test/entities/authentication_entity"
	"web-service-test/entities/database_entity"
)

func (authRepo AuthRepository) FindLoginUser(requestLogin authentication.RequestLogin) (bool, error, *database_entity.Users) {
	user := new(database_entity.Users)
	findUser := authRepo.db.
		Select([]string{
			"users.id",
			"users.first_name",
			"users.last_name",
			"users.username",
			"users.password",
			"users.permission",
		}).
		Where("users.username = ? AND users.user_status = ?", requestLogin.Username, database_entity.ActiveStatus).
		Find(&user)
	if findUser.RowsAffected == 0 {
		return false, errors.New("user not found"), nil
	}
	return true, nil, user
}
