package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
)

func GormMysql() *gorm.DB {
	dsn := os.Getenv("CONNECT_DB") + os.Getenv("DB_HOST") + os.Getenv("CONNECT_DB_LAST")
	var db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})
	db = db.Debug()
	if err != nil {
		log.Println("gorm.open", err)
	}
	return db

}
