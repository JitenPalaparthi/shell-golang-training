package database

import (
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	RETRY_TIMES    = 10
	RETRY_DURATION = time.Second * 5
)

func GetConnection(dsn string) (db *gorm.DB, err error) {
	for i := 1; i <= RETRY_TIMES; i++ {
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			time.Sleep(RETRY_DURATION)
		} else {
			return db, nil
		}
	}
	return db, err
}

//dsn := "host=localhost user=postgres password=postgres dbname=demodb port=9920 sslmode=disable TimeZone=Asia/Shanghai"
