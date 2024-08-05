package main

import (
	"errors"
	"log"
	"time"

	"github.com/google/uuid"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Base struct {
	ID        string `gorm:"type:char(36);primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type User struct {
	Base
	SsoId string
	Name  string
}

type UserInput struct {
	SsoId string
	Name  string
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	u.ID = uuid.New().String()
	return nil
}

func connectDB() *gorm.DB {
	dsn := "admin:admin1234@tcp(127.0.0.1:3306)/demo?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	return db
}

func CreateUser(db *gorm.DB, input *UserInput) error {
	err := db.Where("sso_id = ?", input.SsoId).First(&User{}).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return db.Create(&User{SsoId: input.SsoId, Name: input.Name}).Error
	}

	return err
}
