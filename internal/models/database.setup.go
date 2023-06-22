package models

import (
	"github.com/asargin-dev/soundproof-backend-demo/pkg/logger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDataBase(dsn string, log *logger.Logger) {

	var err error

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Database connection failed")
	} else {
		log.Info("Successfully connected to the database")
	}

	err = DB.AutoMigrate(&User{})
	if err != nil {
		return
	}
}
