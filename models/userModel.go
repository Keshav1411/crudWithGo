package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	First_name string
	Last_name  string
	Email      string `gorm:"unique"`
	Password   string `gorm:"-"`
	Age        int
	Phone_no   int
}