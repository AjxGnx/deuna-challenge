package db

import (
	"sync"

	"github.com/AjxGnx/deuna-challenge/internal/domain/model"
	"github.com/labstack/gommon/log"

	_ "github.com/lib/pq"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	instance *gorm.DB
	once     sync.Once
)

func ConnInstance() *gorm.DB {
	once.Do(func() {
		instance = getConnection()
	})

	return instance
}

func getConnection() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("deuna.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	if err = db.AutoMigrate(model.Customer{}, model.Merchant{}, model.Transaction{}); err != nil {
		log.Fatal(err)
	}

	var customerCount, merchantCount int64
	db.Model(&model.Customer{}).Count(&customerCount)
	db.Model(&model.Merchant{}).Count(&merchantCount)

	if customerCount == 0 {
		db.Create(&model.Customer{})
	}
	if merchantCount == 0 {
		db.Create(&model.Merchant{})
	}

	return db

	return db
}
