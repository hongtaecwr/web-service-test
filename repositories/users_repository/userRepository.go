package users_repository

import (
	"github.com/jinzhu/gorm"
	"log"
	"web-service-test/database"
	"web-service-test/entities/user_entity"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository() UserRepository {
	db, err := database.NewDatabaseConnection()
	if err != nil {
		log.Println("database connection error")
		panic(err)
	}
	return UserRepository{db: db}
}

type UserRepositories interface {
	FindOneUser(id int) (bool, error, user_entity.ResponseUserDetail)
}
