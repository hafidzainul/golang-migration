package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// ConfigDB variable
var ConfigDB DBConfig

// DBConfig struct
type DBConfig struct {
	Username string
	Password string
	Host     string
	Port     string
	Name     string
}

// InitMigration function
func InitMigration(conf DBConfig) {
	ConfigDB = conf
}

// Connection function
func Connection() *gorm.DB {
	var conn string

	conn = ConfigDB.Username + ":" + ConfigDB.Password + "@tcp(" + ConfigDB.Host + ":" + ConfigDB.Port + ")/" + ConfigDB.Name + "?charset=utf8mb4&parseTime=True&loc=Local"

	db, errConn := gorm.Open(mysql.Open(conn), &gorm.Config{})

	if errConn != nil {
		panic("failed to connect to database!")
	} else {
		fmt.Println("database connection success!")
	}

	return db
}
