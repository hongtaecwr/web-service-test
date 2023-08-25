package database

import (
	"github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
	"time"
	"web-service-test/config"
)

func NewDatabaseConnection() (*gorm.DB, error) {
	DBMS := "mysql"
	timezone, _ := time.LoadLocation("Asia/Bangkok")

	mySqlConfig := &mysql.Config{
		User:      config.C.Database.User,
		Passwd:    config.C.Database.Password,
		Net:       config.C.Database.Net,
		Addr:      config.C.Database.Addr,
		DBName:    config.C.Database.DBName,
		Loc:       timezone,
		ParseTime: true,
	}

	db, err := gorm.Open(DBMS, mySqlConfig.FormatDSN())
	if err != nil {
		log.Println(err)
		return db, err
	} else {
		log.Println("database connected")
	}

	return db, nil
}
