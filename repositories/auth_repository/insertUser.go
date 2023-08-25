package auth_repository

import (
	authentication "web-service-test/entities/authentication_entity"
)

func (authRepo AuthRepository) InsertUser(requestRegister authentication.RequestRegister) (bool, error) {
	transaction := authRepo.db.Begin()
	insertUser := transaction.Create(&requestRegister)
	if insertUser.RowsAffected == 0 {
		transaction.Rollback()
		return false, insertUser.Error
	}
	transaction.Commit()
	return true, nil
}
