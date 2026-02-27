package config

import (
	"gopportunities/schemas"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitializeSQLite() (*gorm.DB, error) {
	logger := GetLogger()

	dbPath := "./db/main.db"

	// Check if database file exists
	_, err := os.Stat(dbPath)

	if os.IsNotExist(err) {
		logger.Info("Creating database file...")

		err = os.MkdirAll("./db", os.ModePerm)
		if err != nil {
			return nil, err
		}

		file, err := os.Create(dbPath)

		if err != nil {
			return nil, err
		}

		file.Close()
	}

	// Create DB and Connect
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		logger.Errorf("Failed to connect database %v", err)
		return nil, err
	}

	// Migrate the schema
	err = db.AutoMigrate(&schemas.Opening{})
	if err != nil {
		logger.Errorf("Failed to auto migrate database %v", err)
		return nil, err
	}

	// Return DB
	return db, nil
}
