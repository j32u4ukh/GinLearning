package database

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

type Config struct {
	UserName     string `yaml:"username"`
	Password     string `yaml:"password"`
	Server       string `yaml:"server"`
	Port         int    `yaml:"port"`
	DatabaseName string `yaml:"databasename"`
}

func ConnectMySQL(username string, password string, server string, port int, database string) {
	// dsn := "[username]:[password]@tcp(127.0.0.1:3306)/[databasename]?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, server, port, database)
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
}
