package config

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitializeDatabase() error {
	Logger.Info("F03: Initializing database connection..")
	db, err := gorm.Open(postgres.Open(envVars.dsn), &gorm.Config{})

	if err != nil {
		return fmt.Errorf("F04: Errror while connecting database %s", err.Error())
	}
	DB = db
	Logger.Info("F09: Successfully connected to database..")
	return nil
}
