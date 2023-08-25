package auth_repository

import (
	"github.com/jinzhu/gorm"
	"log"
	"web-service-test/database"
	authentication "web-service-test/entities/authentication_entity"
	"web-service-test/entities/database_entity"
)

type AuthRepository struct {
	db *gorm.DB
}

func NewAuthRepository() AuthRepository {
	db, err := database.NewDatabaseConnection()
	if err != nil {
		log.Println("database connection error")
		panic(err)
	}
	return AuthRepository{db: db}
}

type AuthRepositories interface {
	Register(requestRegister authentication.RequestRegister) error
	FindLoginUser(requestLogin authentication.RequestLogin) (bool, error, database_entity.Users)
}
