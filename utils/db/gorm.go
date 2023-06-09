package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

func GormMysql() *gorm.DB {
	var db *gorm.DB
	var err error

	for {
		// Attempt to establish a database connection
		db, err = connectToMySQL()
		if err == nil {
			break // Connection successful, break out of the loop
		}

		log.Println("Failed to connect to the database:", err)
		time.Sleep(3 * time.Second) // Wait for 3 seconds before the next attempt
	}

	return db
}

func connectToMySQL() (*gorm.DB, error) {
	dsn := os.Getenv("CONNECT_DB") + os.Getenv("DB_HOST") + os.Getenv("CONNECT_DB_LAST")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})
	if err != nil {
		return nil, err
	}

	db = db.Debug()
	return db, nil
}
