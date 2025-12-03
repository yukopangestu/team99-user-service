package database

import (
	"log"
	"team99_user_service/config"
	"team99_user_service/module/model"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB(cfg *config.Config) (*gorm.DB, error) {
	dsn := cfg.GetDSN()

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	sqlDB, err := DB.DB()
	if err != nil {
		return nil, err
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	log.Println("Database connection established successfully")
	return DB, nil
}

func AutoMigrate() error {
	log.Println("Running auto-migration...")

	err := DB.AutoMigrate(
		model.User{})

	if err != nil {
		return err
	}

	log.Println("Auto-migration completed successfully")
	return nil
}

func GetDB() *gorm.DB {
	return DB
}
