package models

import (
	"gorm.io/gorm"
)


type Delivery struct {
	gorm.Model
	ID uint `gorm:"primaryKey;autoIncrement"`
	SupplierId int
	DriverId  int
}