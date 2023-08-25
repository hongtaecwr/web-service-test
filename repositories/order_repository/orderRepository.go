package order_repository

import (
	"github.com/jinzhu/gorm"
	"log"
	"web-service-test/database"
)

type OrderRepository struct {
	db *gorm.DB
}

func NewOrderRepository() OrderRepository {
	db, err := database.NewDatabaseConnection()
	if err != nil {
		log.Println("database connection error")
		panic(err)
	}
	return OrderRepository{db: db}
}

type OrderRepositories interface {
	InsertOrder(addressId, userId int, productIds []int) (bool, error)
}
