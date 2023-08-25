package product_repository

import (
	"github.com/jinzhu/gorm"
	"log"
	"web-service-test/database"
)

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository() ProductRepository {
	db, err := database.NewDatabaseConnection()
	if err != nil {
		log.Println("database connection error")
		panic(err)
	}
	return ProductRepository{db: db}
}
