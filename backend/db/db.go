package db

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/joho/godotenv"
)

func NewDB() *gorm.DB {
	if os.Getenv("GO_ENV") == "dev" {
		err := godotenv.Load()
		if err != nil {
			log.Fatalln(err)
		}
	}
	// dsn := mysql.Config{
	// 	DBName: os.Getenv("MYSQL_DATABASE"),
	// 	User:   os.Getenv("MYSQL_USER"),
	// 	Passwd: os.Getenv("MYSQL_PASSWORD"),
	// 	Addr: "http://localhost:3306/",
	// }
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("MYSQL_USER"),
		os.Getenv("MYSQL_PASSWORD"),
		os.Getenv("MYSQL_HOST"),
		os.Getenv("MYSQL_PORT"),
		os.Getenv("MYSQL_DATABASE"))
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Connected")
	return db
}

func CloseDB(db *gorm.DB) {
	sqlDB, _ := db.DB()
	if err := sqlDB.Close(); err != nil {
		log.Fatalln(err)
	}
}