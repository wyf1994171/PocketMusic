package dal

import (
	"github.com/jinzhu/gorm"
	"time"
)

var (
	db *gorm.DB
)

func InitDB(dsn string) error{
	var err error
	db, err = gorm.Open("mysql", dsn)
	if err != nil {
		return err
	}
	db.DB().SetConnMaxLifetime(time.Duration(1) * time.Minute)
	db.DB().SetMaxIdleConns(5)
	db.DB().SetMaxOpenConns(20)
	return nil
}