package auth_usecase

import (
	"errors"
	"time"
	authentication "web-service-test/entities/authentication_entity"
	"web-service-test/entities/database_entity"
	"web-service-test/functions"
	"web-service-test/middlewares"
	"web-service-test/repositories/auth_repository"
)

type AuthUseCase struct {
	authRepo auth_repository.AuthRepository
}

func NewAuthUseCase(authRepo auth_repository.AuthRepository) AuthUseCase {
	return AuthUseCase{authRepo: authRepo}
}

func (au AuthUseCase) Register(requestRegister authentication.RequestRegister) (bool, error) {
	requestRegister.UserStatus = database_entity.ActiveStatus
	requestRegister.Permission = database_entity.PermissionUser
	requestRegister.CreatedBy = requestRegister.FirstName + " " + requestRegister.LastName
	requestRegister.UpdatedBy = requestRegister.FirstName + " " + requestRegister.LastName
	requestRegister.Password = functions.HashData(requestRegister.Password)

	ok, err := au.authRepo.InsertUser(requestRegister)
	if !ok {
		return false, err
	}

	return true, nil
}

func (au AuthUseCase) Login(requestLogin authentication.RequestLogin) (bool, error, *authentication.LoginResponse) {
	loginResponse := new(authentication.LoginResponse)
	ok, err, user := au.authRepo.FindLoginUser(requestLogin)
	if !ok {
		return false, err, nil
	}

	isMatch := functions.CompareData(user.Password, requestLogin.Password)
	if !isMatch {
		return false, errors.New("password incorrect"), nil
	}

	const tokenExpired = time.Hour * 12
	token, expired := middlewares.GenToken(user, tokenExpired, user.Permission)
	expiredTime := time.Unix(expired, 0).Format("2 Jan 2006 15:01:05")

	loginResponse.AccessToken = token
	loginResponse.Name = user.FirstName + " " + user.LastName
	loginResponse.ExpiresIn = expiredTime
	loginResponse.Permission = user.Permission

	return true, nil, loginResponse
}
