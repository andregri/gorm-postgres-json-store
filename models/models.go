package models

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Orders []Order
	Data   string `sql:"type:JSONB NOT NULL DEFAULT '{}'::JSONB" json:"-"`
}

type Order struct {
	gorm.Model
	UserID uint
	Data   string `sql:"type:JSONB NOT NULL DEFAULT '{}'::JSONB"`
}

func InitDB() (*gorm.DB, error) {
	dsn := "host=localhost user=postgres password=password dbname=mydb port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	} else {
		// Create tables
		db.AutoMigrate(&User{}, &Order{})
		return db, nil
	}
}
