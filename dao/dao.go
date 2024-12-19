package dao

import (
	"fmt"
	"gokweb/config"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	Db *gorm.DB
	// err error
)

func init() {

	db, err := gorm.Open(mysql.Open(config.MysqlDB), &gorm.Config{})
	Db = db
	if err != nil {
		fmt.Println("failed to connect database", err.Error())
		panic("failed to connect database")
	}

	if Db.Error != nil {
		fmt.Println("this error is database error", Db.Error)
		panic("this error is database error")
	}

	sqlDB, err := Db.DB()

	if err != nil {
		fmt.Println("hahaha")
	}

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(time.Hour)

}
