package config

import (
	"gorm.io/gorm"
)

var (
	db     *gorm.DB
	logger *Logger
)

func Init() error {
	var err error
	// Initialize SQLite
	db, err = InitializeSQLite()
	if err != nil {
		panic(err)
	}
	return nil
}

func GetSQLite() *gorm.DB {
	return db
}

func GetLogger() *Logger {
	logger := newLogger()
	return logger
}
