package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // postgresql
)

// GetDB get db connection
func GetDB() (*gorm.DB, error) {
	db, err := gorm.Open("postgres", "host=127.0.0.1 port=5432 user=postgres dbname=dp password=pg_password sslmode=disable")
	if err != nil {
		return nil, err
	}
	return db, nil
}
