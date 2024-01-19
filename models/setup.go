package models

import (
	"fmt"
	"os"

	"github.com/gilangmahardhika/golang-web/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var conn = config.DatabaseConfig{
	User:         os.Getenv("MYSQL_USER"),
	Password:     os.Getenv("MYSQL_PASSWORD"),
	Host:         os.Getenv("MYSQL_HOST"),
	Port:         os.Getenv("MYSQL_PORT"),
	DatabaseName: os.Getenv("GOLANG_WEB_DB_NAME"),
}

var DB *gorm.DB

func ConnectDatabase() {
	fmt.Println(config.ConnectDB(&conn))
	db, err := gorm.Open(mysql.Open(config.ConnectDB(&conn)))
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&Book{}, &User{}, &Review{})
	DB = db
}
