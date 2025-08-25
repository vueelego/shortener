package dbrepo

import (
	"fmt"
	"shortener/internal/models"

	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"gorm.io/driver/mysql"
)

var Db *gorm.DB

func OpenDb() error {
	dbUserName := "root"
	dbPassword := "root.cc"
	dbHost := "127.0.0.1"
	dbPort := 3306
	dbName := "db_shortener"
	source := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbUserName,
		dbPassword,
		dbHost,
		dbPort,
		dbName,
	)

	conn, err := gorm.Open(mysql.Open(source), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		return err
	}

	if conn.Error != nil {
		return err
	}

	mysqlDB, err := conn.DB()
	if err != nil {
		return err
	}

	mysqlDB.SetMaxIdleConns(10)
	mysqlDB.SetMaxOpenConns(25)

	Db = conn

	return nil
}

func DbMigrate() error {
	return Db.AutoMigrate(
		&models.User{},
		&models.Entry{},
		&models.Session{},
		&models.Quota{},
		&models.Click{},
		&models.Tag{},
	)
}
