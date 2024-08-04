package db

import (
	"gorm.io/gorm"
	"database/sql"
	"time"

	"github.com/go-sql/driver-mysql"

)

func NewDB() *gorm.DB {
	dsn := mysql.Config(
		DBName: 
	)
}