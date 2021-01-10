package models

import (
	"gorm.io/gorm"
)


type Driver struct {
	gorm.Model
	ID uint `gorm:"primaryKey;autoIncrement"`
	Name string
	CarrierId int
}