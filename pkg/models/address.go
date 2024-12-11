package models

import "gorm.io/gorm"

type Address struct {
	gorm.Model
	UserID uint `validate:"required"`
	//User       User   `gorm:"foriegnkey:UserID;references:ID"`
	Country    string `validate:"required"`
	State      string `validate:"required"`
	District   string `validate:"required"`
	StreetName string `validate:"required"`
	PinCode    string `validate:"required,numeric"`
	Phone      string `validate:"required,numeric,len=10"`
	Default    bool   `gorm:"default:false" validate:"required"`
}
